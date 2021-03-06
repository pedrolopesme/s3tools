//MIT License
//
//Copyright (c) 2018 Pedro Mendes
//
//Permission is hereby granted, free of charge, to any person obtaining a copy
//of this software and associated documentation files (the "Software"), to deal
//in the Software without restriction, including without limitation the rights
//to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
//copies of the Software, and to permit persons to whom the Software is
//furnished to do so, subject to the following conditions:
//
//The above copyright notice and this permission notice shall be included in all
//copies or substantial portions of the Software.
//
//THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
//IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
//AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
//LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
//SOFTWARE.

package cmd

import (
	"fmt"
	"github.com/pedrolopesme/s3tools/utils"
	"github.com/spf13/cobra"
)

// grepCmd represents the grep command
var grepCmd = &cobra.Command{
	Use:   "grep",
	Short: "The grep utility searches any given input files, selecting lines that match one or more patterns.",

	Long: `The grep utility searches any given input files, selecting lines that match one or more patterns.  By default, a pattern matches an input line if
     the regular expression (RE) in the pattern matches the input line without its trailing newline. Each input
     line that matches at least one of the patterns is written to the standard output. For example:

s3tools grep my-bucket "search string"`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("You must provide bucket and a search pattern")
			return
		}

		bucket := args[0]
		pattern := args[1]

		path, err := cmd.Flags().GetString("path")
		if err != nil {
			fmt.Println("It was impossible to parse path parameter")
		}
		utils.GrepFiles(bucket, pattern, path)
	},
}

func init() {
	grepCmd.Flags().StringP("path", "p", "", "path where grep will look for")
	rootCmd.AddCommand(grepCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// grepCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// grepCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
