// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
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

package bundleCmd

import "github.com/spf13/cobra"
import "github.com/jinzhu/gorm"

//NewBundleCommand initial BundleCommand
func NewBundleCommand(db *gorm.DB) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bundle",
		Short: "manage the bundles",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	cmd.AddCommand(
		NewBundleAddCmd(db),
		NewBundleDeleteCmd(db),
		NewBundleCreateCmd(db),
		NewBundleRemoveCmd(db),
		NewBundleUpdateCmd(db),
		NewBundleListCmd(db),
	)
	return cmd
}
