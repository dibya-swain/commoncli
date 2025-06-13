package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var TimeCmd = &cobra.Command{
	Use:   "show-time",
	Short: "Show Current Time (from common cli)",
	Run: func(cmd *cobra.Command, args []string) {
		currentTime := time.Now()
		fmt.Println("Current Time", currentTime)
	},
}

func init() {
	Register(TimeCmd)
}
