package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"io"
	"net/http"
	"time"
)

type Res struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	Phase            string `json:"phase"`
	Duration         int    `json:"duration"`
	StartTimeSeconds int    `json:"startTimeSeconds"`
}
type Contest struct {
	Status string `json:"status"`
	Result []Res  `json:"result"`
}

func init() {
	rootCmd.AddCommand(contestCmd)
}

var contestCmd = &cobra.Command{
	Use:   "contest",
	Short: "Show the upcoming Contests",
	Long:  `Show all the upcoming contests on codeforces`,
	Run: func(cmd *cobra.Command, args []string) {
		url := "https://codeforces.com/api/contest.list?gym=false"
		res, err := http.Get(url)
		if err != nil {
			fmt.Println("Cannot Load Api")
			return
		}
		body, err := io.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}
		var ContestList Contest
		err = json.Unmarshal(body, &ContestList)
		if err != nil {
			panic("Cannot unmarshal result")
		}
		cnt := 1
		for _, val := range ContestList.Result {
			if val.Phase == "BEFORE" {
				seconds := int64(val.StartTimeSeconds)
				timeObj := time.Unix(seconds, 0)
				dateStr := timeObj.Format("03:04 PM Mon, Jan, 2006 ")
				fmt.Printf("%d:  %s	", cnt, val.Name)
				color.Magenta(dateStr)
				cnt++
			}
		}
	},
}
