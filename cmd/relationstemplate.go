// Copyright (c) 2019 bketelsen
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package cmd

import (
	"fmt"
	"github.com/bketelsen/lxdev/lxd"
	"github.com/lxc/lxd/shared/i18n"
	"github.com/mitchellh/go-homedir"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
	"strings"
)

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

// store the container template (image) relation in yaml file
// takes as arguments the container name, relating template name
// for store set variable store true
func setContainerTemplateRealation(c *lxd.Client, container string, tmpl string, store bool) error {
	var templates Templates

	// store the realation
	if store {

		err := templates.parse()
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
				err = templates.store()
				if err != nil {
					return errors.Wrap(err, "Failed storing")
				}
				return nil
			}
		}

		fingerprint, err := c.GetImageFingerprint(tmpl)
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

		err := templates.parse()
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
					err = templates.store()
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

	err := t.store()
	if err != nil {
		return err
	}

	return nil
}

// Removing the templates which is an LXC image
func removeTemplate(c *lxd.Client, tmpl string) error {
	var templates Templates

	fingerprint, err := c.GetImageFingerprint(tmpl)
	if err != nil {
		return errors.Wrap(err, "Failed getting fingerprint")
	}

	err = templates.parse()
	if err != nil {
		return errors.Wrap(err, "Failed parse data")
	}

	for i, template := range templates.Templates {
		if template.UsedBy == "" && template.Name == tmpl {
			err = c.ContainerRemove(tmpl)
			if err != nil {
				return errors.Wrap(err, "Failed remove template container")
			}
			err = lxd.RemoveTemplateImage(c, fingerprint)
			if err != nil {
				return errors.Wrap(err, "Failed remove template image "+tmpl)
			}
			copy(templates.Templates[i:], templates.Templates[i+1:])
			templates.Templates[len(templates.Templates)-1] = Template{}
			templates.Templates = templates.Templates[:len(templates.Templates)-1]
			err = templates.store()
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

// Unmarshal helper function
// path is hardcoded?
func (t *Templates) parse() error {
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
func (t *Templates) store() error {
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
