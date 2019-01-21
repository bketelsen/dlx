// Copyright © 2019 Brian Ketelsen mail@bjk.fyi
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// profileCmd represents the profile command
var profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "create or replace the provisioning profile for lxdev",
	Long: `Profile creates or replaces the 'gui', 'cli', and 'util' profiles in lxc that allows you
to connect to running containers and possibly display X11 applications on the host.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("TODO: profile called", viper.GetString("ethernet"))
	},
}

func init() {
	rootCmd.AddCommand(profileCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	profileCmd.PersistentFlags().String("ethernet", "", "the name of your ethernet device e.g. 'enp5s0'")
	viper.BindPFlag("ethernet", profileCmd.PersistentFlags().Lookup("ethernet"))

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// profileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
