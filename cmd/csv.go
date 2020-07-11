/*
Copyright © 2020 Ken'ichiro Oyama <k1lowxb@gmail.com>

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
	"io"
	"os"

	"github.com/k1LoW/tbls-cacoo/csv"
	"github.com/k1LoW/tbls/datasource"
	"github.com/spf13/cobra"
)

var csvCmd = &cobra.Command{
	Use:   "csv",
	Short: "generate CSV for Cacoo's Database Schema Importer",
	Long:  `generate CSV for Cacoo's Database Schema Importer (https://support.cacoo.com/hc/en-us/articles/360045672494).`,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			o   *os.File
			err error
		)
		if out == "" {
			o = os.Stdout
		} else {
			o, err = os.Create(out)
			if err != nil {
				printFatalln(cmd, err)
			}
		}
		if err := runCSV(cmd, o); err != nil {
			printFatalln(cmd, err)
		}
	},
}

func runCSV(cmd *cobra.Command, stdout io.Writer) error {
	c := csv.New()
	s, err := datasource.AnalyzeJSONStringOrFile(os.Getenv("TBLS_SCHEMA"))
	if err != nil {
		return err
	}
	return c.OutputSchema(stdout, s)
}

func init() {
	rootCmd.AddCommand(csvCmd)
	csvCmd.Flags().StringVarP(&out, "out", "o", "", "output file path")
}
