package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"io"
	"net/http"
)

type UserContestResult struct {
	ContestName string `json:"contestName"`
	Rank        int    `json:"rank"`
	OldRating   int    `json:"oldRating"`
	NewRating   int    `json:"newRating"`
}

type UserContestResponse struct {
	Status string              `json:"status"`
	Result []UserContestResult `json:"result"`
}

func init() {
	userCmd.AddCommand(userContestCmd)
}

var userContestCmd = &cobra.Command{
	Use:   "contests",
	Short: "Show users last 10 contests",
	Long:  `Show user contests`,
	Run: func(cmd *cobra.Command, args []string) {
		url := "https://codeforces.com/api/user.rating?handle="
		if len(args) == 1 {
			handle := args[0]
			url += handle
			res, err := http.Get(url)
			if err != nil {
				fmt.Println("Cannot Load Api")
				return
			}
			body, err := io.ReadAll(res.Body)
			if err != nil {
				panic(err)
			}
			var userContest UserContestResponse
			err = json.Unmarshal(body, &userContest)
			if err != nil {
				panic("Cannot unmarshal result")
			}
			if userContest.Status != "OK" {
				panic("Handle Not Found")
			}
			cnt := 1
			for id := len(userContest.Result) - 1; id >= 0; id-- {
				if cnt > 10 {
					break
				}
				val := userContest.Result[id]
				fmt.Printf("%d:\n", cnt)
				cnt++
				fmt.Println("Contest Name:", val.ContestName)
				fmt.Println("Rank:", val.Rank)
				fmt.Printf("Rating: %d -> %d\n", val.OldRating, val.NewRating)
				change := val.NewRating - val.OldRating
				fmt.Printf("Rating change: ")
				if change < 0 {
					color.Red("%d", change)
				} else if change == 0 {
					fmt.Println(change)
				} else {
					color.Green("+%d", change)
				}
				fmt.Println()
			}
		} else {
			fmt.Println("Please provide one handle")
		}
	},
}
