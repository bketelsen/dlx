package dlx

import (
	"os"
	"strings"

	"github.com/bketelsen/dlx/state"
	"github.com/olekukonko/tablewriter"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

type CmdList struct {
	Global *state.Global
}

func (c *CmdList) Command() *cobra.Command {

	// listCmd represents the list command
	var listCmd = &cobra.Command{
		Use:     "list",
		Short:   "list containers",
		Aliases: []string{"ls"},
		Long:    `List containers and their status.`,
		RunE:    c.Run,
	}
	return listCmd
}

func (c *CmdList) Run(cmd *cobra.Command, args []string) error {

	var err error
	cfg, err = getDlxConfig()
	if err != nil {
		log.Error(err.Error())
		return err
	}
	conf := c.Global.Conf

	d, err := conf.GetInstanceServer(conf.DefaultRemote)
	if err != nil {
		return err
	}

	names, err := d.GetContainerNames()
	if err != nil {
		errors.Wrap(err, "get container names")
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Status", "Profile(s)"})

	for _, name := range names {
		container, _, err := d.GetInstance(name)
		if err != nil {
			log.Error("Get Container: " + err.Error())
			os.Exit(1)
		}
		table.Append([]string{container.Name, container.Status, strings.Join(container.Profiles, ",")})

	}
	if len(names) < 1 {

		table.Append([]string{"{None Found}", "", ""})
	}
	table.Render()
	return nil
}
