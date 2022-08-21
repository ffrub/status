package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

func init() {
    rootCmd.AddCommand(tpCmd)
}

var (
        tpCmd = &cobra.Command{
            Use: "tp",
            Short: "Print the reported status for tp",
            Run: func(cmd *cobra.Command, args []string) {
                var tpRes Response = &TpRes{}
                errCh := GetStatus("https://status.targetprocess.com/api/v2/status.json", tpRes)
                for err := range errCh {
                    if err != nil { cobra.CheckErr(err) }
                }

                tpRes.Print()
            },
        }
    )

type TpRes struct {
	Page struct {
		ID        string    `json:"id"`
		Name      string    `json:"name"`
		URL       string    `json:"url"`
		TimeZone  string    `json:"time_zone"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"page"`
	Status struct {
		Indicator   string `json:"indicator"`
		Description string `json:"description"`
	} `json:"status"`
}

func (res *TpRes) Print() {
    if res.Status.Description == "All Systems Operational" {
        fmt.Printf("\x1b[%dm%s\x1b[0m", 32, "✔ ")
        fmt.Printf("\x1b[%dm%s\x1b[0m\n", 1, "Targetprocess")
    } else {
        fmt.Printf("\x1b[%dm%s\x1b[0m", 31, "✘ ")
        fmt.Printf("\x1b[%dm%s\x1b[0m\n", 1, "Targetprocess")
        fmt.Printf("\t\x1b[%dm%s\x1b[0m\n", 1, res.Status.Indicator)
        fmt.Printf("\t\x1b[%dm%s\x1b[0m\n", 1, res.Status.Description)
    }
}
