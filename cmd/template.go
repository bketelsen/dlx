// Copyright © 2019 bketelsen
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package cmd

import (
	"github.com/spf13/cobra"
)

var templateCmd = &cobra.Command{
	Use:   "template",
	Short: "manage container templates",
	Long:  `Manage the templates used to create containers`,
}

func init() {
	rootCmd.AddCommand(templateCmd)
}
