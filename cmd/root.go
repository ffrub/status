/*
Copyright © 2022 Finn Rueb ff.rueb@gmail.com

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
	"os"

	"github.com/spf13/cobra"
)

type Response interface {
    Print()
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "status",
    Version: "v0.1.0",
	Short: "print the reported status of services",
    Long: `status lets you print the reported status of:
    Heroku, GitHub, Targetprocess and Courier`,
	Run: func(cmd *cobra.Command, args []string) {
        var herokuRes Response = &HerokuRes{}
        var githubRes Response = &GithubRes{}
        var tpRes Response = &TpRes{}
        var courierRes Response = &CourierRes{}

        for err := range Merge(
            GetStatus("https://status.heroku.com/api/v4/current-status", herokuRes),
            GetStatus("https://www.githubstatus.com/api/v2/summary.json", githubRes),
            GetStatus("https://status.targetprocess.com/api/v2/status.json", tpRes),
            GetStatus("https://status.courier.com/api/v2/status.json", courierRes),
        ) {
            if err != nil {
                cobra.CheckErr(err)
            }
        }
        
        herokuRes.Print()
        githubRes.Print()
        tpRes.Print()
        courierRes.Print()
   },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

