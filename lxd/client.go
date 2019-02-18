// Copyright (c) 2019 bketelsen
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package lxd

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
	"sort"
	"strings"

	"github.com/bketelsen/libgo/events"
	client "github.com/lxc/lxd/client"
	lxd "github.com/lxc/lxd/client"
	"github.com/lxc/lxd/shared/api"
	"github.com/lxc/lxd/shared/i18n"
	"github.com/pkg/errors"
)

type Client struct {
	URL string

	conn client.ContainerServer
}

// represents an LXD Image
type Image struct {
	Fingerprint string `yaml:"fingerprint,omitempty"`
}

// represents a template
type Template struct {
	Name   string `yaml:"name"`
	UsedBy string `yaml:"usedBy,omitempty"`
	Image  *Image `yaml:"images,omitempty"`
}

// represents the template "collection"
type Templates struct {
	Templates []Template `yaml:"templates"`
}

// NewClient creates a new connection to an LXD Daemon,
// returning a Client
func NewClient(url string) (*Client, error) {
	c := &Client{
		URL: url,
	}
	err := c.Connect()
	return c, err
}

// Connect establishes a connection to an LXD Daemon
func (c *Client) Connect() error {
	var err error
	c.conn, err = client.ConnectLXDUnix(c.URL, nil)
	if err != nil {
		return errors.Wrap(err, "Error connecting to LXD daemon")
	}

	events.Publish(NewConnectionCreated(c))
	return nil
}

func (c *Client) ContainerProvision(name string, kind Type, provisioners []string) error {
	cont, err := GetContainer(c.conn, name)
	if err != nil {
		return errors.Wrap(err, "getting container")
	}
	return cont.Provision(kind, provisioners)
}

func (c *Client) ContainerShell(name string) error {
	cont, err := GetContainer(c.conn, name)
	if err != nil {
		return errors.Wrap(err, "getting container")
	}
	return cont.Exec("", true)
}

func (c *Client) ContainerCreate(name string, isAlias bool, image string, profiles []string) error {
	// Container creation request
	var source api.ContainerSource
	if isAlias {
		source = api.ContainerSource{
			Type: "image",
			//Server:   "https://cloud-images.ubuntu.com/daily",
			//Alias:    getImage(),
			Alias: image,
			//Protocol: "simplestreams",
		}
	} else {
		source = api.ContainerSource{
			Type:     "image",
			Server:   "https://cloud-images.ubuntu.com/daily",
			Alias:    image,
			Protocol: "simplestreams",
		}
	}
	req := api.ContainersPost{
		Name: name,
		ContainerPut: api.ContainerPut{
			Profiles: profiles,
		},
		Source: source,
	}

	events.Publish(NewContainerState(name, Creating))

	// Get LXD to create the container (background operation)
	op, err := c.conn.CreateContainer(req)
	if err != nil {
		return errors.Wrap(err, "creating container")
	}

	// Wait for the operation to complete
	err = op.Wait()
	if err != nil {
		return errors.Wrap(err, "wait for create container")
	}

	events.Publish(NewContainerState(name, Created))
	// Get LXD to start the container (background operation)
	reqState := api.ContainerStatePut{
		Action:  "start",
		Timeout: -1,
	}

	events.Publish(NewContainerState(name, Starting))
	op, err = c.conn.UpdateContainerState(name, reqState, "")
	if err != nil {
		return errors.Wrap(err, "starting container")
	}

	// Wait for the operation to complete
	err = op.Wait()
	if err != nil {
		return errors.Wrap(err, "waiting for container start")
	}

	events.Publish(NewContainerState(name, Started))
	return nil
}
func (c *Client) ContainerExec(name string, command string) error {
	cont, err := GetContainer(c.conn, name)
	if err != nil {
		return errors.Wrap(err, "getting container")
	}
	return cont.Exec(command, false)
}

func (c *Client) ContainerList() ([]string, error) {
	names, err := c.conn.GetContainerNames()
	if err != nil {
		errors.Wrap(err, "get container names")
	}
	return names, err
}

func (c *Client) ContainerInfo(name string) (*api.Container, error) {
	container, _, err := c.conn.GetContainer(name)
	if err != nil {
		errors.Wrap(err, "get container names")
	}
	return container, err
}

func (c *Client) ContainerRemove(name string) error {
	cont, err := GetContainer(c.conn, name)
	if err != nil {
		return errors.Wrap(err, "getting container")
	}
	return cont.Remove()
}

func (c *Client) ContainerStart(name string) error {
	cont, err := GetContainer(c.conn, name)
	if err != nil {
		return errors.Wrap(err, "getting container")
	}
	return cont.Start()
}

func (c *Client) ContainerStop(name string) error {
	cont, err := GetContainer(c.conn, name)
	if err != nil {
		return errors.Wrap(err, "getting container")
	}
	return cont.Stop()
}

func (c *Client) ImageList() ([]api.Image, error) {
	return c.conn.GetImages()

}
func (c *Client) ContainerSnapshot(name string, snapshotName string) error {
	cont, err := GetContainer(c.conn, name)
	if err != nil {
		return errors.Wrap(err, "getting container")
	}
	return cont.Snapshot(snapshotName)
}

func (c *Client) ContainerPublish(name string) error {

	fmt.Println("Publishing: ", name)
	// Create the image
	req := api.ImagesPost{
		Source: &api.ImagesPostSource{
			Type: "container",
			Name: name + "/template", // UGH?

		},
	}
	// skipping properties, that may be a mistake?
	req.Source.Type = "snapshot"
	req.Public = true

	alias := api.ImageAlias{}
	alias.Name = name
	alias.Description = "lxdev template: " + name
	op, err := c.conn.CreateImage(req, nil)
	if err != nil {
		return errors.Wrap(err, "create container image")
	}
	// Wait for it to complete
	err = op.Wait()
	if err != nil {
		return errors.Wrap(err, "wait for operation")
	}
	opAPI := op.Get()

	// Grab the fingerprint
	fingerprint := opAPI.Metadata["fingerprint"].(string)
	return ensureImageAliases(c.conn, []api.ImageAlias{alias}, fingerprint)

}

// Create the specified image alises, updating those that already exist
// copied from lxd source :)
func ensureImageAliases(client lxd.ContainerServer, aliases []api.ImageAlias, fingerprint string) error {
	if len(aliases) == 0 {
		return nil
	}

	names := make([]string, len(aliases))
	for i, alias := range aliases {
		names[i] = alias.Name
	}
	sort.Strings(names)

	resp, err := client.GetImageAliases()
	if err != nil {
		return err
	}

	// Delete existing aliases that match provided ones
	for _, alias := range GetExistingAliases(names, resp) {
		err := client.DeleteImageAlias(alias.Name)
		if err != nil {
			fmt.Println(fmt.Sprintf(i18n.G("Failed to remove alias %s"), alias.Name))
		}
	}
	// Create new aliases
	for _, alias := range aliases {
		aliasPost := api.ImageAliasesPost{}
		aliasPost.Name = alias.Name
		aliasPost.Target = fingerprint
		err := client.CreateImageAlias(aliasPost)
		if err != nil {
			fmt.Println(fmt.Sprintf(i18n.G("Failed to create alias %s"), alias.Name))
		}
	}
	return nil
}

// GetExistingAliases returns the intersection between a list of aliases and all the existing ones.
func GetExistingAliases(aliases []string, allAliases []api.ImageAliasesEntry) []api.ImageAliasesEntry {
	existing := []api.ImageAliasesEntry{}
	for _, alias := range allAliases {
		name := alias.Name
		pos := sort.SearchStrings(aliases, name)
		if pos < len(aliases) && aliases[pos] == name {
			existing = append(existing, alias)
		}
	}
	return existing
}

// store the container template (image) relation in yaml file
// takes as arguments the container name, relating template name
// for store set variable store true
func (c *Client) SetContainerTemplateRealation(container string, tmpl string, store bool) error {
	var templates Templates

	// store the realation
	if store {

		err := templates.Parse()
		if err != nil {
			return errors.Wrap(err, "Failed parse data")
		}

		// append template users which is a container
		// if an entry is already there an return to caller
		for i, template := range templates.Templates {
			if tmpl == template.Name {

				sep := strings.Split(template.UsedBy, ",")
				for _, usedby := range sep {
					if usedby == container {
						return fmt.Errorf(i18n.G("Entry already here nothing to do"))
					}
				}

				templates.Templates[i].UsedBy += "," + container
				err = templates.Store()
				if err != nil {
					return errors.Wrap(err, "Failed storing")
				}
				return nil
			}
		}

		fingerprint, err := getImageFingerprint(c, tmpl)
		if err != nil {
			return errors.Wrap(err, "Failed getting fingerprint")
		}

		// if not create a new entry
		err = writeEntry(&templates, container, tmpl, fingerprint)
		if err != nil {
			return errors.Wrap(err, "Failed storing")
		}

	} else {
		// else delete the relation

		err := templates.Parse()
		if err != nil {
			return errors.Wrap(err, "Failed parse data")
		}

		if templates.Templates == nil {
			return fmt.Errorf(i18n.G("Error no image entry here, maybe something went wrong?"))
		}

		for i, template := range templates.Templates {
			sep := strings.Split(template.UsedBy, ",")
			for j, str := range sep {
				if container == str {
					copy(sep[j:], sep[j+1:])
					sep[len(sep)-1] = ""
					sep = sep[:len(sep)-1]
					templates.Templates[i].UsedBy = strings.Join(sep, ",")
					err = templates.Store()
					if err != nil {
						return err
					}
					break
				}
			}
		}
	}

	return nil
}

// write container - template (image) relation
func writeEntry(t *Templates, container string, tmpl string, fingerprint string) error {
	img := &Image{Fingerprint: fingerprint}

	template := Template{Name: tmpl, UsedBy: container, Image: img}
	t.Templates = append(t.Templates, template)

	err := t.Store()
	if err != nil {
		return err
	}

	return nil
}

// retrieves the image fingerprint from template (image) name
func getImageFingerprint(c *Client, template string) (string, error) {
	var retVal string
	imagesAPI, err := c.conn.GetImages()

	if err != nil {
		return "", fmt.Errorf(i18n.G("Failed getting container metadata"))
	}

	for _, imageAPI := range imagesAPI {
		for _, alias := range imageAPI.Aliases {
			if alias.Name == template {
				retVal = imageAPI.Fingerprint
			}
		}
	}
	return retVal, nil
}

// Removing the templates which is an LXC image
func (c *Client) RemoveTemplate(tmpl string) error {
	var templates Templates

	fingerprint, err := getImageFingerprint(c, tmpl)
	if err != nil {
		return errors.Wrap(err, "Failed getting fingerprint")
	}

	err = templates.Parse()
	if err != nil {
		return errors.Wrap(err, "Failed parse data")
	}

	for i, template := range templates.Templates {
		if template.UsedBy == "" && template.Name == tmpl {
			err = c.ContainerRemove(tmpl)
			if err != nil {
				return errors.Wrap(err, "Failed remove template container")
			}
			err = removeTemplateImage(c, fingerprint)
			if err != nil {
				return errors.Wrap(err, "Failed remove template image "+tmpl)
			}
			copy(templates.Templates[i:], templates.Templates[i+1:])
			templates.Templates[len(templates.Templates)-1] = Template{}
			templates.Templates = templates.Templates[:len(templates.Templates)-1]
			err = templates.Store()
			if err != nil {
				return errors.Wrap(err, "Failed parse data")
			}
			break
		} else if template.Name == tmpl {
			return fmt.Errorf(i18n.G("Error can not remove image, it's still in use by " + template.UsedBy))
		}
	}

	return nil
}

// helper function deletes the LXC image
func removeTemplateImage(c *Client, fingerprint string) error {
	op, err := c.conn.DeleteImage(fingerprint)
	if err != nil {
		return errors.Wrap(err, "Failed to remove image")
	}

	err = op.Wait()
	if err != nil {
		return errors.Wrap(err, "waiting for removing image")
	}
	return nil
}

// Unmarshal helper function
// path is hardcoded?
func (t *Templates) Parse() error {
	home, err := homedir.Dir()
	if err != nil {
		return err
	}
	filename := filepath.Join(home, ".lxdev", "templates", "relations.yaml")

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(data, t); err != nil {
		return err
	}

	return nil
}

// Marshal helper function
func (t *Templates) Store() error {
	home, err := homedir.Dir()
	if err != nil {
		return err
	}
	filename := filepath.Join(home, ".lxdev", "templates", "relations.yaml")

	bytes, err := yaml.Marshal(t)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, bytes, 0644)
}
