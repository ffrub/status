package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

func init() {
    rootCmd.AddCommand(courierCmd)
}

var (
        courierCmd = &cobra.Command{
            Use: "courier",
            Short: "Print the reported status for courier",
            Run: func(cmd *cobra.Command, args []string) {
                var courierRes Response = &CourierRes{}
                errChan := GetStatus("https://status.courier.com/api/v2/status.json", courierRes)
                for err := range errChan {
                    if err != nil { cobra.CheckErr(err) }
                }

                courierRes.Print()
            },
        }
    )

type CourierRes struct {
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

func (res *CourierRes) Print() {
    if res.Status.Description == "All Systems Operational" {
        fmt.Printf("\x1b[%dm%s\x1b[0m", 32, "✔ ")
        fmt.Printf("\x1b[%dm%s\x1b[0m\n", 1, "Courier")
    } else {
        fmt.Printf("\x1b[%dm%s\x1b[0m", 31, "✘ ")
        fmt.Printf("\x1b[%dm%s\x1b[0m\n", 1, "Courier")
        fmt.Printf("\t\x1b[%dm%s\x1b[0m\n", 1, res.Status.Indicator)
        fmt.Printf("\t\x1b[%dm%s\x1b[0m\n", 1, res.Status.Description)
    }
}
