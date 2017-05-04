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
	"github.com/spf13/cobra"
	"github.com/tony24681379/bookstore/bundle"
)

const (
	defaultElasticSearchIP = "0.0.0.0:9200"
	flagElasticSearchIP    = "elastic-search-ip"
	flagElasticSearchPort  = "elastic-search-port"
	flagAkkaIP             = "akka-ip"
)

type serverOptions struct {
	elasticSearchIP   string
	elasticSearchPort string
	akkaIP            string
}

//NewBundleCommand initial BundleCommand
func NewBundleCommand() *cobra.Command {
	opts := serverOptions{}
	cmd := &cobra.Command{
		Use:   "bundle",
		Short: "bundle some books in a bundle",
		Run: func(cmd *cobra.Command, args []string) {
			runBundle(opts)
		},
	}

	cmd.AddCommand(NewBundleAddCmd())
	return cmd
}

func runBundle(opts serverOptions) {
	bundle.Bundle()
}
