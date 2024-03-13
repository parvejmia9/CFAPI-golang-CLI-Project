package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"io"
	"net/http"
)

type InfoResult struct {
	LastName     string `json:"lastName"`
	Country      string `json:"country"`
	City         string `json:"city"`
	Rating       int    `json:"rating"`
	Organization string `json:"organization"`
	Rank         string `json:"rank"`
	MaxRating    int    `json:"maxRating"`
	FirstName    string `json:"firstName"`
}

type InfoResponse struct {
	Status string       `json:"status"`
	Result []InfoResult `json:"result"`
}

func init() {
	userCmd.AddCommand(infoCmd)
}

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Show the user statistics",
	Long:  `Show user info contests and submissions`,
	Run: func(cmd *cobra.Command, args []string) {
		url := "https://codeforces.com/api/user.info?handles="
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
			var info InfoResponse
			err = json.Unmarshal(body, &info)
			if err != nil {
				panic("Cannot unmarshal result")
			}
			if info.Status != "OK" {
				panic("Handle Not Found")
			}
			fmt.Println("Name:", info.Result[0].FirstName, info.Result[0].LastName)
			fmt.Println(info.Result[0].City, info.Result[0].Country)
			fmt.Println(info.Result[0].Organization)
			rtg := info.Result[0].Rating
			if rtg < 1200 {
				fmt.Println(info.Result[0].Rank)
			} else if rtg < 1400 {
				color.Green("%s", info.Result[0].Rank)
			} else if rtg < 1900 {
				color.Blue("%s", info.Result[0].Rank)
			} else if rtg < 2100 {
				color.Magenta("%s", info.Result[0].Rank)
			} else if rtg < 2400 {
				color.Yellow("%s", info.Result[0].Rank)
			} else {
				color.Red("%s", info.Result[0].Rank)
			}
			fmt.Println("Max Rating:", info.Result[0].MaxRating)
			fmt.Println("Current Rating:", info.Result[0].Rating)
		} else {
			fmt.Println("Please provide one handle")
		}
	},
}
