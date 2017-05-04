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

package cmd

import (
	"fmt"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
)

type options struct {
	Debug bool
}

var opts = options{}

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "Smart manage system",
	Short: "A good cooperation tool for bookstore chain",
	Long:  `A golang backend server serve frontend which connects to ElasticSearch and Akka`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if opts.Debug == true {
			logrus.SetLevel(logrus.DebugLevel)
			fmt.Println("debug")
		}
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

//InitCmd init cobra command
func InitCmd() {
	initProgramFlag()
	RootCmd.AddCommand(NewBundleCommand())
}

func initProgramFlag() {
	RootCmd.PersistentFlags().BoolVarP(&opts.Debug, "debug", "D", false, "Enable debug mode")
}