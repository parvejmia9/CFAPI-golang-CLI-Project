package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"net/http"
)

type Problem struct {
	ContestID int    `json:"contestId"`
	Index     string `json:"index"`
	Name      string `json:"name"`
}

type Result struct {
	Problems []Problem `json:"problems"`
}

type Response struct {
	Status string `json:"status"`
	Result Result `json:"result"`
}

func init() {
	rootCmd.AddCommand(problemCmd)
}

var problemCmd = &cobra.Command{
	Use:   "problem",
	Short: "Show the Problems",
	Long:  `Show Problems with specified tags`,
	Run: func(cmd *cobra.Command, args []string) {
		url := "https://codeforces.com/api/problemset.problems"
		if len(args) > 0 {
			url += "?tags="
			for i, val := range args {
				if i != 0 {
					url += ";"
				}
				url += val
			}

		}
		res, err := http.Get(url)
		if err != nil {
			fmt.Println("Cannot Load Api")
			return
		}
		body, err := io.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}
		var ProblemList Response
		err = json.Unmarshal(body, &ProblemList)
		if err != nil {
			panic("Cannot unmarshal result")
		}
		if ProblemList.Status != "OK" {
			fmt.Println("No problem found")
		} else {
			for i, val := range ProblemList.Result.Problems {
				if i < 10 {
					fmt.Printf("%d:  %d%s    %s\n", i+1, val.ContestID, val.Index, val.Name)
				}
			}
		}

	},
}
