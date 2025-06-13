package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var HelloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Hello from shared CLI",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello from shared CLI!")
	},
}

func init() {
	Register(HelloCmd)
}
