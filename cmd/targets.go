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
	"os"
	"text/tabwriter"

	"bender/bender"
	"github.com/spf13/cobra"
)

// targetsCmd represents the targets command
var targetsCmd = &cobra.Command{
	Use:   "targets",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		targets, err := bender.GetTargets(url)
		if err != nil {
			fmt.Println(err)
			return
		}
		w := new(tabwriter.Writer)
		w.Init(os.Stdout, 8, 8, 8, '\t', 0)

		fmt.Fprintln(w, "ID\tName\tOS")
		for _, target := range targets {
			fmt.Fprintln(w, fmt.Sprintf("%d\t%s\t%s\n", target.ID, target.Name, target.OS))
		}
		return
	},
}

func init() {
	getCmd.AddCommand(targetsCmd)
}
