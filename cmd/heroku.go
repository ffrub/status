package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
    rootCmd.AddCommand(herokuCmd)
}

var (
        herokuCmd = &cobra.Command{
            Use: "heroku",
            Short: "Print the reported status for heroku",
            Run: func(cmd *cobra.Command, args []string) {
                var herokuRes Response = &HerokuRes{}
                errCh := GetStatus("https://status.heroku.com/api/v4/current-status", herokuRes)
                for err := range errCh {
                    if err != nil { cobra.CheckErr(err) }
                }
                    
                herokuRes.Print()
            },
        }
    )

type HerokuRes struct {
    Status []struct {
        System string `json:"system"`
        Status string `json:"status"`
    } `json:"status"`
	Incidents []interface{} `json:"incidents"`
	Scheduled []interface{} `json:"scheduled"`
}

func isAllGreen(s []struct {
    System string `json:"system"`
    Status string `json:"status"`
}) bool {
    for _, rec := range s {
        if rec.Status != "green" {
            return false
        }
    }

    return true
}

func (res *HerokuRes) Print() {
    var systemDict = map[string]string{
        "Apps": "\t\tDynos, Routing, Scheduler",
        "Data": "\t\tHeroku Postgres, Heroku Data for Redis, Heroku Connect",
        "Tools": "\t\tGit-push deployments and Deployment API\n" +
        "\t\tGitHub Sync, API / CLI / Heroku Dashboard (scaling up/down, changing configuration, etc.\n" +
        "\t\tLogging, Unidling free dynos",
    }

    if isAllGreen(res.Status) {
        fmt.Printf("\x1b[%dm%s\x1b[0m", 32, "✔ ")
        fmt.Printf("\x1b[%dm%s\x1b[0m\n", 1, "Heroku")
    } else {
        fmt.Printf("\x1b[%dm%s\x1b[0m", 31, "✘ ")
        fmt.Printf("\x1b[%dm%s\x1b[0m\n", 1, "Heroku")
        for _, rec := range res.Status {
            if rec.Status != "green" {
                fmt.Printf("\t\x1b[%dm%s\x1b[0m", 31, "✘ ")
                fmt.Printf("\x1b[%dm%s\x1b[0m\n", 240, rec.System)
                fmt.Printf("\x1b[%dm%s\x1b[0m\n", 3, systemDict[rec.System])
            } else {
                fmt.Printf("\t\x1b[%dm%s\x1b[0m", 32, "✔ ")
                fmt.Printf("%s\n", rec.System)
            }
        }
    }
}
