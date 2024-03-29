/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"
)

// healthCmd represents the health command
var healthCmd = &cobra.Command{
	Use:   "health",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return health()
	},
}

// Leaving this for now, in case I change my mind and go back to using json.Unmarshal.
type healthResponse struct {
	ClusterName         string `json:"cluster_name"`
	Status              string `json:"string"`
	Nodes               int    `json:"number_of_nodes"`
	DataNodes           int    `json:"number_of_data_nodes"`
	ActivePrimaryShards int    `json:"active_primary_shards"`
}

func health() error {
	cluster := esAPI().Cluster
	resp, err := cluster.Health(esAPI().Cluster.Health.WithPretty())
	if err != nil {
		return err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	return nil
}

func init() {
	showCmd.AddCommand(healthCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// healthCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// healthCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
