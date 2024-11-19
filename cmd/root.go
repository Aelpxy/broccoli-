package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fresh",
	Short: "Fresh is a modern, efficient file server powered by Go.",
	Long:  "Fresh is a modern, efficient file server powered by Go.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
		} else {
			fmt.Println("Unknown command. Please use '--help' to see available commands.")
		}
	},
}

func Execute() {
	rootCmd.Execute()
}
