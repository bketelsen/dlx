package dlx

import (
	"github.com/bketelsen/dlx/dlx/config"
	"github.com/bketelsen/dlx/state"
	"github.com/spf13/cobra"
)

type CmdConfig struct {
	Global *state.Global
}

// configCmd represents the config command
func (c *CmdConfig) Command() *cobra.Command {

	var configCmd = &cobra.Command{
		Use:   "config",
		Short: "manage global configurations",
		Long:  ``,
		RunE:  c.Run,
	}
	configCmd.Flags().BoolP("create", "c", false, "Create global config file")
	configCmd.Flags().StringP("user", "u", "ubuntu", "Container username")
	configCmd.Flags().StringP("baseimage", "b", "dlxbase", "Default base image for new containers")
	configCmd.Flags().StringArrayP("profiles", "p", []string{}, "Profiles to use")
	configCmd.Flags().StringP("sshkey", "s", "", "Path to ssh private key authorized for HOST")
	return configCmd
}
func (c *CmdConfig) Run(cmd *cobra.Command, args []string) error {

	log.Running("Configuration")
	createConfig, err := cmd.Flags().GetBool("create")
	if err != nil {
		log.Error("Error getting flags: " + err.Error())
		return err
	}

	if createConfig {
		if verbose {
			log.Info("Creating dlx config file")
		}
		err := config.Create(cmd, c.Global.Conf)
		if err != nil {
			log.Error("Error creating config file: " + err.Error())
			return err
		}
		log.Success("Default configuration file created")
		return nil
	}
	return nil

}
