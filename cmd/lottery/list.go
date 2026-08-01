package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/onyx-and-iris/lottery-cli"
)

// listCmd represents the list command.
var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List all available lottery kinds.",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println(renderKindList(lottery.AllKinds()))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
