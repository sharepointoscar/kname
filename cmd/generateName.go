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
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strings"
)

// generateNameCmd represents the generateName command
var generateNameCmd = &cobra.Command{
	Use:   "generate-name",
	Short: "Generates the cluster name.",
	Long:  `This addOn generates your EKS or GKE cluster name from a JSON file.`,
	Run: func(cmd *cobra.Command, args []string) {

		// prob should go elsewhere but including it here for now.
		// get an entry from the cocktails.json
		// Open our jsonFile
		jsonFile, err := os.Open("cocktails.json")

		// if we os.Open returns an error then handle it
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Successfully Opened cocktails.json")

		// read our opened jsonFile as a byte array.
		// byteValue, _ := ioutil.ReadAll(jsonFile)

		// type Cocktail struct {
		// 	Name  string `json:"Name"`
		// 	Glass string `json:"Glass"`
		// }

		// type Cocktails struct {
		// 	Cocktails []Cocktail `json:"cocktails"`
		// }
		// defer the closing of our jsonFile so that we can parse it later on
		defer jsonFile.Close()
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