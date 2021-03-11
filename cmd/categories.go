/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"os"
	"text/tabwriter"

	"bender/bender"
	"github.com/spf13/cobra"
)

// categoriesCmd represents the categories command
var categoriesCmd = &cobra.Command{
	Use:   "categories",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		categories, err := bender.GetCategories(url)
		if err != nil {
			fmt.Println(err)
			return
		}
		w := new(tabwriter.Writer)
		w.Init(os.Stdout, 8, 8, 8, '\t', 0)

		fmt.Fprintln(w, "ID\tName\tDescription")
		for _, category := range categories {
			fmt.Fprintln(w, fmt.Sprintf("%d\t%s\t%s\n", category.ID, category.Name, category.Description))
		}
		return
	},
}

func init() {
	getCmd.AddCommand(categoriesCmd)
}
