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
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/k1LoW/tbls-cacoo/version"
	"github.com/spf13/cobra"
)

var out string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tbls-cacoo",
	Short: "tbls-cacoo is an external subcommand of tbls for Cacoo",
	Long:  `tbls-cacoo is an external subcommand of tbls for Cacoo.`,
}

func Execute() {
	rootCmd.SetOut(os.Stdout)
	rootCmd.SetErr(os.Stderr)

	log.SetOutput(ioutil.Discard)
	if env := os.Getenv("DEBUG"); env != "" {
		debug, err := os.Create(fmt.Sprintf("%s.debug", version.Name))
		if err != nil {
			printFatalln(rootCmd, err)
		}
		log.SetOutput(debug)
	}

	if err := rootCmd.Execute(); err != nil {
		printFatalln(rootCmd, err)
	}
}

// https://github.com/spf13/cobra/pull/894
func printErrln(c *cobra.Command, i ...interface{}) {
	c.PrintErr(fmt.Sprintln(i...))
}

func printErrf(c *cobra.Command, format string, i ...interface{}) {
	c.PrintErr(fmt.Sprintf(format, i...))
}

func printFatalln(c *cobra.Command, i ...interface{}) {
	printErrln(c, i...)
	os.Exit(1)
}

func printFatalf(c *cobra.Command, format string, i ...interface{}) {
	printErrf(c, format, i...)
	os.Exit(1)
}

func init() {}
