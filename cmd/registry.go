package cmd

import "github.com/spf13/cobra"

var Commands = map[string]*cobra.Command{
	"hello":     HelloCmd,
	"show-time": TimeCmd,
}
