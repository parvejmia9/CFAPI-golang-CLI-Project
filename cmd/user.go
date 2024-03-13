package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(userCmd)
}

var userCmd = &cobra.Command{
	Use:   "usr",
	Short: "Show the user statistics",
	Long:  `Show user info contests and submissions`,
}
