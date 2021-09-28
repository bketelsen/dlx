package lxc

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"

	"github.com/bketelsen/dlx/state"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"

	"github.com/lxc/lxd/lxc/utils"
	"github.com/lxc/lxd/shared"
	"github.com/lxc/lxd/shared/api"
	cli "github.com/lxc/lxd/shared/cmd"
	"github.com/lxc/lxd/shared/i18n"
	"github.com/lxc/lxd/shared/termios"
	"github.com/lxc/lxd/shared/units"
)

type cmdProject struct {
	global *state.Global
}

func (c *cmdProject) Command() *cobra.Command {
	cmd := &cobra.Command{}
	cmd.Use = usage("project")
	cmd.Short = i18n.G("Manage projects")
	cmd.Long = cli.FormatSection(i18n.G("Description"), i18n.G(
		`Manage projects`))

	// Create
	projectCreateCmd := cmdProjectCreate{global: c.global, project: c}
	cmd.AddCommand(projectCreateCmd.Command())

	// Delete
	projectDeleteCmd := cmdProjectDelete{global: c.global, project: c}
	cmd.AddCommand(projectDeleteCmd.Command())

	// Edit
	projectEditCmd := cmdProjectEdit{global: c.global, project: c}
	cmd.AddCommand(projectEditCmd.Command())

	// Get
	projectGetCmd := cmdProjectGet{global: c.global, project: c}
	cmd.AddCommand(projectGetCmd.Command())

	// List
	projectListCmd := cmdProjectList{global: c.global, project: c}
	cmd.AddCommand(projectListCmd.Command())

	// Rename
	projectRenameCmd := cmdProjectRename{global: c.global, project: c}
	cmd.AddCommand(projectRenameCmd.Command())

	// Set
	projectSetCmd := cmdProjectSet{global: c.global, project: c}
	cmd.AddCommand(projectSetCmd.Command())

	// Unset
	projectUnsetCmd := cmdProjectUnset{global: c.global, project: c, projectSet: &projectSetCmd}
	cmd.AddCommand(projectUnsetCmd.Command())

	// Show
	projectShowCmd := cmdProjectShow{global: c.global, project: c}
	cmd.AddCommand(projectShowCmd.Command())

	// Info
	projectGetInfo := cmdProjectInfo{global: c.global, project: c}
	cmd.AddCommand(projectGetInfo.Command())

	// Set default
	projectSwitchCmd := cmdProjectSwitch{global: c.global, project: c}
	cmd.AddCommand(projectSwitchCmd.Command())

	// Workaround for subcommand usage errors. See: https://github.com/spf13/cobra/issues/706
	cmd.Args = cobra.NoArgs
	cmd.Run = func(cmd *cobra.Command, args []string) { cmd.Usage() }
	return cmd
}

// Create
type cmdProjectCreate struct {
	global     *state.Global
	project    *cmdProject
	flagConfig []string
}

func (c *cmdProjectCreate) Command() *cobra.Command {
	cmd := &cobra.Command{}
	cmd.Use = usage("create", i18n.G("[<remote>:]<project>"))
	cmd.Short = i18n.G("Create projects")
	cmd.Long = cli.FormatSection(i18n.G("Description"), i18n.G(
		`Create projects`))
	cmd.Flags().StringArrayVarP(&c.flagConfig, "config", "c", nil, i18n.G("Config key/value to apply to the new project")+"``")

	cmd.RunE = c.Run

	return cmd
}

func (c *cmdProjectCreate) Run(cmd *cobra.Command, args []string) error {
	// Quick checks.
	exit, err := c.global.CheckArgs(cmd, args, 1, 1)
	if exit {
		return err
	}

	// Parse remote
	resources, err := c.global.ParseServers(args[0])
	if err != nil {
		return err
	}

	resource := resources[0]

	if resource.Name == "" {
		return fmt.Errorf(i18n.G("Missing project name"))
	}

	// Create the project
	project := api.ProjectsPost{}
	project.Name = resource.Name

	project.Config = map[string]string{}
	for _, entry := range c.flagConfig {
		if !strings.Contains(entry, "=") {
			return fmt.Errorf(i18n.G("Bad key=value pair: %s"), entry)
		}

		fields := strings.SplitN(entry, "=", 2)
		project.Config[fields[0]] = fields[1]
	}

	err = resource.Server.CreateProject(project)
	if err != nil {
		return err
	}

	if !c.global.FlagQuiet {
		fmt.Printf(i18n.G("Project %s created")+"\n", resource.Name)
	}

	return nil
}

// Delete
type cmdProjectDelete struct {
	global  *state.Global
	project *cmdProject
}

func (c *cmdProjectDelete) Command() *cobra.Command {
	cmd := &cobra.Command{}
	cmd.Use = usage("delete", i18n.G("[<remote>:]<project>"))
	cmd.Aliases = []string{"rm"}
	cmd.Short = i18n.G("Delete projects")
	cmd.Long = cli.FormatSection(i18n.G("Description"), i18n.G(
		`Delete projects`))

	cmd.RunE = c.Run

	return cmd
}

func (c *cmdProjectDelete) Run(cmd *cobra.Command, args []string) error {
	// Quick checks.
	exit, err := c.global.CheckArgs(cmd, args, 1, 1)
	if exit {
		return err
	}

	// Parse remote
	remoteName, _, err := c.global.Conf.ParseRemote(args[0])
	if err != nil {
		return err
	}

	resources, err := c.global.ParseServers(args[0])
	if err != nil {
		return err
	}

	resource := resources[0]

	if resource.Name == "" {
		return fmt.Errorf(i18n.G("Missing project name"))
	}

	// Delete the project
	err = resource.Server.DeleteProject(resource.Name)
	if err != nil {
		return err
	}

	if !c.global.FlagQuiet {
		fmt.Printf(i18n.G("Project %s deleted")+"\n", resource.Name)
	}

	// Switch back to default project
	if resource.Name == c.global.Conf.Remotes[remoteName].Project {
		rc := c.global.Conf.Remotes[remoteName]
		rc.Project = ""
		c.global.Conf.Remotes[remoteName] = rc
		return c.global.Conf.SaveConfig(c.global.ConfPath)
	}

	return nil
}

// Edit
type cmdProjectEdit struct {
	global  *state.Global
	project *cmdProject
}

func (c *cmdProjectEdit) Command() *cobra.Command {
	cmd := &cobra.Command{}
	cmd.Use = usage("edit", i18n.G("[<remote>:]<project>"))
	cmd.Short = i18n.G("Edit project configurations as YAML")
	cmd.Long = cli.FormatSection(i18n.G("Description"), i18n.G(
		`Edit project configurations as YAML`))
	cmd.Example = cli.FormatSection("", i18n.G(
		`lxc project edit <project> < project.yaml
    Update a project using the content of project.yaml`))

	cmd.RunE = c.Run

	return cmd
}

func (c *cmdProjectEdit) helpTemplate() string {
	return i18n.G(
		`### This is a YAML representation of the project.
### Any line starting with a '# will be ignored.
###
### A project consists of a set of features and a description.
###
### An example would look like:
### name: my-project
### features:
###   images: True
###   profiles: True
### description: My own project
###
### Note that the name is shown but cannot be changed`)
}

func (c *cmdProjectEdit) Run(cmd *cobra.Command, args []string) error {
	// Quick checks.
	exit, err := c.global.CheckArgs(cmd, args, 1, 1)
	if exit {
		return err
	}

	// Parse remote
	resources, err := c.global.ParseServers(args[0])
	if err != nil {
		return err
	}

	resource := resources[0]

	if resource.Name == "" {
		return fmt.Errorf(i18n.G("Missing project name"))
	}

	// If stdin isn't a terminal, read text from it
	if !termios.IsTerminal(getStdinFd()) {
		contents, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			return err
		}

		newdata := api.ProjectPut{}
		err = yaml.Unmarshal(contents, &newdata)
		if err != nil {
			return err
		}

		return resource.Server.UpdateProject(resource.Name, newdata, "")
	}

	// Extract the current value
	project, etag, err := resource.Server.GetProject(resource.Name)
	if err != nil {
		return err
	}

	data, err := yaml.Marshal(&project)
	if err != nil {
		return err
	}

	// Spawn the editor
	content, err := shared.TextEditor("", []byte(c.helpTemplate()+"\n\n"+string(data)))
	if err != nil {
		return err
	}

	for {
		// Parse the text received from the editor
		newdata := api.ProjectPut{}
		err = yaml.Unmarshal(content, &newdata)
		if err == nil {
			err = resource.Server.UpdateProject(resource.Name, newdata, etag)
		}

		// Respawn the editor
		if err != nil {
			fmt.Fprintf(os.Stderr, i18n.G("Config parsing error: %s")+"\n", err)
			fmt.Println(i18n.G("Press enter to open the editor again or ctrl+c to abort change"))

			_, err := os.Stdin.Read(make([]byte, 1))
			if err != nil {
				return err
			}

			content, err = shared.TextEditor("", content)
			if err != nil {
				return err
			}
			continue
		}
		break
	}

	return nil
}

// Get
type cmdProjectGet struct {
	global  *state.Global
	project *cmdProject
}

func (c *cmdProjectGet) Command() *cobra.Command {
	cmd := &cobra.Command{}
	cmd.Use = usage("get", i18n.G("[<remote>:]<project> <key>"))
	cmd.Short = i18n.G("Get values for project configuration keys")
	cmd.Long = cli.FormatSection(i18n.G("Description"), i18n.G(
		`Get values for project configuration keys`))

	cmd.RunE = c.Run

	return cmd
}

func (c *cmdProjectGet) Run(cmd *cobra.Command, args []string) error {
	// Quick checks.
	exit, err := c.global.CheckArgs(cmd, args, 2, 2)
	if exit {
		return err
	}

	// Parse remote
	resources, err := c.global.ParseServers(args[0])
	if err != nil {
		return err
	}

	resource := resources[0]

	if resource.Name == "" {
		return fmt.Errorf(i18n.G("Missing project name"))
	}

	// Get the configuration key
	project, _, err := resource.Server.GetProject(resource.Name)
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", project.Config[args[1]])
	return nil
}

// List
type cmdProjectList struct {
	global  *state.Global
	project *cmdProject

	flagFormat string
}

func (c *cmdProjectList) Command() *cobra.Command {
	cmd := &cobra.Command{}
	cmd.Use = usage("list", i18n.G("[<remote>:]"))
	cmd.Aliases = []string{"ls"}
	cmd.Short = i18n.G("List projects")
	cmd.Long = cli.FormatSection(i18n.G("Description"), i18n.G(
		`List projects`))
	cmd.Flags().StringVarP(&c.flagFormat, "format", "f", "table", i18n.G("Format (csv|json|table|yaml)")+"``")

	cmd.RunE = c.Run

	return cmd
}

func (c *cmdProjectList) Run(cmd *cobra.Command, args []string) error {
	conf := c.global.Conf

	// Quick checks.
	exit, err := c.global.CheckArgs(cmd, args, 0, 1)
	if exit {
		return err
	}

	// Parse remote
	remote := conf.DefaultRemote
	if len(args) > 0 {
		remote = args[0]
	}
	remoteName := strings.TrimSuffix(remote, ":")

	resources, err := c.global.ParseServers(remote)
	if err != nil {
		return err
	}

	resource := resources[0]

	// List projects
	projects, err := resource.Server.GetProjects()
	if err != nil {
		return err
	}

	currentProject := conf.Remotes[remoteName].Project
	if currentProject == "" {
		currentProject = "default"
	}

	data := [][]string{}
	for _, project := range projects {
		images := i18n.G("NO")
		if shared.IsTrue(project.Config["features.images"]) {
			images = i18n.G("YES")
		}

		profiles := i18n.G("NO")
		if shared.IsTrue(project.Config["features.profiles"]) {
			profiles = i18n.G("YES")
		}

		storageVolumes := i18n.G("NO")
		if shared.IsTrue(project.Config["features.storage.volumes"]) {
			storageVolumes = i18n.G("YES")
		}

		networks := i18n.G("NO")
		if shared.IsTrue(project.Config["features.networks"]) {
			networks = i18n.G("YES")
		}

		name := project.Name
		if name == currentProject {
			name = fmt.Sprintf("%s (%s)", name, i18n.G("current"))
		}

		strUsedBy := fmt.Sprintf("%d", len(project.UsedBy))
		data = append(data, []string{name, images, profiles, storageVolumes, networks, project.Description, strUsedBy})
	}
	sort.Sort(byName(data))

	header := []string{
		i18n.G("NAME"),
		i18n.G("IMAGES"),
		i18n.G("PROFILES"),
		i18n.G("STORAGE VOLUMES"),
		i18n.G("NETWORKS"),
		i18n.G("DESCRIPTION"),
		i18n.G("USED BY"),
	}

	return utils.RenderTable(c.flagFormat, header, data, projects)
}

// Rename
type cmdProjectRename struct {
	global  *state.Global
	project *cmdProject
}

func (c *cmdProjectRename) Command() *cobra.Command {
	cmd := &cobra.Command{}
	cmd.Use = usage("rename", i18n.G("[<remote>:]<project> <new-name>"))
	cmd.Aliases = []string{"mv"}
	cmd.Short = i18n.G("Rename projects")
	cmd.Long = cli.FormatSection(i18n.G("Description"), i18n.G(
		`Rename projects`))

	cmd.RunE = c.Run

	return cmd
}

func (c *cmdProjectRename) Run(cmd *cobra.Command, args []string) error {
	// Quick checks.
	exit, err := c.global.CheckArgs(cmd, args, 2, 2)
	if exit {
		return err
	}

	// Parse remote
	resources, err := c.global.ParseServers(args[0])
	if err != nil {
		return err
	}

	resource := resources[0]

	if resource.Name == "" {
		return fmt.Errorf(i18n.G("Missing project name"))
	}

	// Rename the project
	op, err := resource.Server.RenameProject(resource.Name, api.ProjectPost{Name: args[1]})
	if err != nil {
		return err
	}

	err = op.Wait()
	if err != nil {
		return err
	}

	if !c.global.FlagQuiet {
		fmt.Printf(i18n.G("Project %s renamed to %s")+"\n", resource.Name, args[1])
	}

	return nil
}

// Set
type cmdProjectSet struct {
	global  *state.Global
	project *cmdProject
}

func (c *cmdProjectSet) Command() *cobra.Command {
	cmd := &cobra.Command{}
	cmd.Use = usage("set", i18n.G("[<remote>:]<project> <key>=<value>..."))
	cmd.Short = i18n.G("Set project configuration keys")
	cmd.Long = cli.FormatSection(i18n.G("Description"), i18n.G(
		`Set project configuration keys

For backward compatibility, a single configuration key may still be set with:
    lxc project set [<remote>:]<project> <key> <value>`))

	cmd.RunE = c.Run

	return cmd
}

func (c *cmdProjectSet) Run(cmd *cobra.Command, args []string) error {
	// Quick checks.
	exit, err := c.global.CheckArgs(cmd, args, 2, -1)
	if exit {
		return err
	}

	// Parse remote
	resources, err := c.global.ParseServers(args[0])
	if err != nil {
		return err
	}

	resource := resources[0]

	if resource.Name == "" {
		return fmt.Errorf(i18n.G("Missing project name"))
	}

	// Get the project
	project, etag, err := resource.Server.GetProject(resource.Name)
	if err != nil {
		return err
	}

	// Set the configuration key
	keys, err := getConfig(args[1:]...)
	if err != nil {
		return err
	}

	for k, v := range keys {
		project.Config[k] = v
	}

	return resource.Server.UpdateProject(resource.Name, project.Writable(), etag)
}

// Unset
type cmdProjectUnset struct {
	global     *state.Global
	project    *cmdProject
	projectSet *cmdProjectSet
}

func (c *cmdProjectUnset) Command() *cobra.Command {
	cmd := &cobra.Command{}
	cmd.Use = usage("unset", i18n.G("[<remote>:]<project> <key>"))
	cmd.Short = i18n.G("Unset project configuration keys")
	cmd.Long = cli.FormatSection(i18n.G("Description"), i18n.G(
		`Unset project configuration keys`))

	cmd.RunE = c.Run

	return cmd
}

func (c *cmdProjectUnset) Run(cmd *cobra.Command, args []string) error {
	// Quick checks.
	exit, err := c.global.CheckArgs(cmd, args, 2, 2)
	if exit {
		return err
	}

	args = append(args, "")
	return c.projectSet.Run(cmd, args)
}

// Show
type cmdProjectShow struct {
	global  *state.Global
	project *cmdProject
}

func (c *cmdProjectShow) Command() *cobra.Command {
	cmd := &cobra.Command{}
	cmd.Use = usage("show", i18n.G("[<remote>:]<project>"))
	cmd.Short = i18n.G("Show project options")
	cmd.Long = cli.FormatSection(i18n.G("Description"), i18n.G(
		`Show project options`))

	cmd.RunE = c.Run

	return cmd
}

func (c *cmdProjectShow) Run(cmd *cobra.Command, args []string) error {
	// Quick checks.
	exit, err := c.global.CheckArgs(cmd, args, 1, 1)
	if exit {
		return err
	}

	// Parse remote
	resources, err := c.global.ParseServers(args[0])
	if err != nil {
		return err
	}

	resource := resources[0]

	if resource.Name == "" {
		return fmt.Errorf(i18n.G("Missing project name"))
	}

	// Show the project
	project, _, err := resource.Server.GetProject(resource.Name)
	if err != nil {
		return err
	}

	data, err := yaml.Marshal(&project)
	if err != nil {
		return err
	}

	fmt.Printf("%s", data)

	return nil
}

// Switch project
type cmdProjectSwitch struct {
	global  *state.Global
	project *cmdProject
}

func (c *cmdProjectSwitch) Command() *cobra.Command {
	cmd := &cobra.Command{}
	cmd.Use = usage("switch", i18n.G("[<remote>:]<project>"))
	cmd.Short = i18n.G("Switch the current project")
	cmd.Long = cli.FormatSection(i18n.G("Description"), i18n.G(
		`Switch the current project`))

	cmd.RunE = c.Run

	return cmd
}

func (c *cmdProjectSwitch) Run(cmd *cobra.Command, args []string) error {
	conf := c.global.Conf

	// Quick checks.
	exit, err := c.global.CheckArgs(cmd, args, 1, 1)
	if exit {
		return err
	}

	// Parse the remote
	remote, project, err := conf.ParseRemote(args[0])
	if err != nil {
		return err
	}

	// Make sure the remote exists
	rc, ok := conf.Remotes[remote]
	if !ok {
		return fmt.Errorf(i18n.G("Remote %s doesn't exist"), remote)
	}

	// Make sure the project exists
	d, err := conf.GetInstanceServer(remote)
	if err != nil {
		return err
	}

	_, _, err = d.GetProject(project)
	if err != nil {
		return err
	}

	rc.Project = project

	conf.Remotes[remote] = rc

	return conf.SaveConfig(c.global.ConfPath)
}

// Info
type cmdProjectInfo struct {
	global  *state.Global
	project *cmdProject

	flagFormat string
}

func (c *cmdProjectInfo) Command() *cobra.Command {
	cmd := &cobra.Command{}
	cmd.Use = usage("info", i18n.G("[<remote>:]<project> <key>"))
	cmd.Short = i18n.G("Get a summary of resource allocations")
	cmd.Long = cli.FormatSection(i18n.G("Description"), i18n.G(
		`Get a summary of resource allocations`))
	cmd.Flags().StringVarP(&c.flagFormat, "format", "f", "table", i18n.G("Format (csv|json|table|yaml)")+"``")

	cmd.RunE = c.Run

	return cmd
}

func (c *cmdProjectInfo) Run(cmd *cobra.Command, args []string) error {
	// Quick checks.
	exit, err := c.global.CheckArgs(cmd, args, 1, 1)
	if exit {
		return err
	}

	// Parse remote
	resources, err := c.global.ParseServers(args[0])
	if err != nil {
		return err
	}

	resource := resources[0]

	if resource.Name == "" {
		return fmt.Errorf(i18n.G("Missing project name"))
	}

	// Get the current allocations
	projectState, err := resource.Server.GetProjectState(resource.Name)
	if err != nil {
		return err
	}

	// Render the output
	byteLimits := []string{"disk", "memory"}
	data := [][]string{}
	for k, v := range projectState.Resources {
		limit := i18n.G("UNLIMITED")
		if v.Limit >= 0 {
			if shared.StringInSlice(k, byteLimits) {
				limit = units.GetByteSizeStringIEC(v.Limit, 2)
			} else {
				limit = fmt.Sprintf("%d", v.Limit)
			}
		}

		usage := ""
		if shared.StringInSlice(k, byteLimits) {
			usage = units.GetByteSizeStringIEC(v.Usage, 2)
		} else {
			usage = fmt.Sprintf("%d", v.Usage)
		}

		data = append(data, []string{strings.ToUpper(k), limit, usage})
	}
	sort.Sort(byName(data))

	header := []string{
		i18n.G("RESOURCE"),
		i18n.G("LIMIT"),
		i18n.G("USAGE"),
	}

	return utils.RenderTable(c.flagFormat, header, data, projectState)
}
