package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"io"
	"net/http"
)

type StatusProblem struct {
	ContestID int    `json:"contestId"`
	Index     string `json:"index"`
	Name      string `json:"name"`
}
type StatusResult struct {
	Prob     StatusProblem `json:"problem"`
	Language string        `json:"programmingLanguage"`
	Verdict  string        `json:"verdict"`
}

type StatusResponse struct {
	Status string         `json:"status"`
	Result []StatusResult `json:"result"`
}

func init() {
	userCmd.AddCommand(statusCmd)
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show users last 10 submissions",
	Long:  `Show user submissions`,
	Run: func(cmd *cobra.Command, args []string) {
		var handle string
		pref := "https://codeforces.com/api/user.status?handle="
		suf := "&from=1&count=10"
		if len(args) == 1 {
			handle = args[0]
			url := pref + handle + suf
			res, err := http.Get(url)
			if err != nil {
				fmt.Println("Cannot Load Api")
				return
			}
			body, err := io.ReadAll(res.Body)
			if err != nil {
				panic(err)
			}
			var status StatusResponse
			err = json.Unmarshal(body, &status)
			if err != nil {
				panic("Cannot unmarshal result")
			}
			if status.Status != "OK" {
				panic("Handle Not Found")
			}
			for i, val := range status.Result {
				fmt.Printf("%d:\n", i+1)
				fmt.Printf("Problem ID: %d%s\n", val.Prob.ContestID, val.Prob.Index)
				fmt.Println("Problem Name:", val.Prob.Name)
				fmt.Printf("Verdict: ")
				if val.Verdict == "OK" {
					color.Green("Accepted")
				} else {
					color.Red(val.Verdict)
				}
				fmt.Println("Language:", val.Language)
				fmt.Println()
			}
		} else {
			fmt.Println("Please provide one handle")
		}
	},
}
