package main

import (
	"CFAPI/cmd"
	"fmt"
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

func main() {
	err := cmd.Execute()
	if err != nil {
		fmt.Println("CMD Error")
		return
	}
}
