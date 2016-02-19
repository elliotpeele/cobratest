// Copyright © 2016 Elliot Peele <elliot@bentlogic.net>
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
)

// sub1Cmd represents the sub1 command
var sub1Cmd = &cobra.Command{
	Use:   "sub1",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("sub1 called")
		toggle, err := cmd.Parent().PersistentFlags().GetBool("toggle")
		if err != nil {
			return err
		}
		fmt.Printf("sub1\ttoggle\t%v\n", toggle)
		return nil
	},
}

func init() {
	RootCmd.AddCommand(sub1Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sub1Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sub1Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
