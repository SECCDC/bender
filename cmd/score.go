/*
Copyright © 2021 Bren 'fraq' Briggs

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package cmd

import (
	"fmt"

	"bender/bender"
	"github.com/spf13/cobra"
)

var (
	teams   []int
	targets []int
	cat     int
	desc    string
)

// scoreCmd represents the score command
var scoreCmd = &cobra.Command{
	Use:   "score",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := bender.Score(teams, targets, cat, desc, url)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(resp)
		return
	},
}

func init() {
	rootCmd.AddCommand(scoreCmd)

	scoreCmd.Flags().IntSliceVarP(&teams, "teams", "t", teams, "Team IDs to score agaisnt")
	scoreCmd.Flags().IntSliceVarP(&targets, "targets", "r", teams, "Systems impacted by the activity")
	scoreCmd.Flags().IntVarP(&cat, "category", "c", cat, "Category for your activity")
	scoreCmd.Flags().StringVarP(&desc, "desc", "d", desc, "string description of your activity")

	scoreCmd.MarkFlagRequired("teams")
	scoreCmd.MarkFlagRequired("targets")
	scoreCmd.MarkFlagRequired("category")
	scoreCmd.MarkFlagRequired("desc")
}
