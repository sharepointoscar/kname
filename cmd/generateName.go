/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"github.com/sharepointoscar/kname/themes"
	"github.com/spf13/cobra"
	"strings"
)

// generateNameCmd represents the generateName command
var generateNameCmd = &cobra.Command{
	Use:   "generate-name",
	Short: "Generates a random kubernetes cluster name based on theme flag.",
	Long:  `This kubectl plugin generates your EKS or GKE cluster name.`,
	Run: func(cmd *cobra.Command, args []string) {

		_theme, _ := rootCmd.Flags().GetString("theme")
		_usedTheme := themes.NewTheme(_theme)

		fmt.Println("what theme???", _usedTheme.Name)

		if strings.Compare(_theme, "yoga") == 0 {

			var _clusterName = _usedTheme.ClusterName
			fmt.Println(_clusterName)

		} else if strings.Compare(_usedTheme.Name, "cocktails") == 0 {

			var _clusterName = _usedTheme.ClusterName
			fmt.Println(_clusterName)

		} else if strings.Compare(_theme, "national-parks") == 0 {

			var _clusterName = _usedTheme.ClusterName
			fmt.Println(_clusterName)
		} else {
			fmt.Println("theme not supported. Want to contribute and create it? Fork the repo!")
		}
	},
}

func init() {
	rootCmd.AddCommand(generateNameCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateNameCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateNameCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
