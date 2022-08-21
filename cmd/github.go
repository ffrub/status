package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)


func init() {
    rootCmd.AddCommand(githubCmd)
}

var (
        githubCmd = &cobra.Command{
            Use: "github",
            Short: "Print the reported status for github",
            Run: func(cmd *cobra.Command, args []string) {
                var githubRes Response = &GithubRes{}
                errChan := GetStatus("https://www.githubstatus.com/api/v2/summary.json", githubRes)
                for err := range errChan {
                    if err != nil { cobra.CheckErr(err) }
                }

                githubRes.Print()
            },
        }
    )

type GithubRes struct {
	Page struct {
		ID        string    `json:"id"`
		Name      string    `json:"name"`
		URL       string    `json:"url"`
		TimeZone  string    `json:"time_zone"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"page"`
	Components []struct {
		ID                 string      `json:"id"`
		Name               string      `json:"name"`
		Status             string      `json:"status"`
		CreatedAt          time.Time   `json:"created_at"`
		UpdatedAt          time.Time   `json:"updated_at"`
		Position           int         `json:"position"`
		Description        string      `json:"description"`
		Showcase           bool        `json:"showcase"`
		StartDate          interface{} `json:"start_date"`
		GroupID            interface{} `json:"group_id"`
		PageID             string      `json:"page_id"`
		Group              bool        `json:"group"`
		OnlyShowIfDegraded bool        `json:"only_show_if_degraded"`
	} `json:"components"`
	Incidents             []interface{} `json:"incidents"`
	ScheduledMaintenances []interface{} `json:"scheduled_maintenances"`
	Status                struct {
		Indicator   string `json:"indicator"`
		Description string `json:"description"`
	} `json:"status"`
}

func (res *GithubRes) Print() {
    if res.Status.Description == "All Systems Operational" {
        fmt.Printf("\x1b[%dm%s\x1b[0m", 32, "✔ ")
        fmt.Printf("\x1b[%dm%s\x1b[0m\n", 1, "Github")
    } else {
        fmt.Printf("\x1b[%dm%s\x1b[0m", 31, "✘ ")
        fmt.Printf("\x1b[%dm%s\x1b[0m\n", 1, "Github")


        for _, comp := range res.Components {
            if comp.Status != "operational" {
                fmt.Printf("\t\x1b[%dm%s\x1b[0m", 31, "✘ ")
                fmt.Printf("\x1b[%dm%s\x1b[0m\n", 240, comp.Name)
                fmt.Printf("\t\t\x1b[%dm%s\x1b[0m\n\n", 3, comp.Description)
            } else {
                fmt.Printf("\t\x1b[%dm%s\x1b[0m", 32, "✔ ")
                fmt.Printf("%s\n", comp.Name)
            }
        }
    }
}
