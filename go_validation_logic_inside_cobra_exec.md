package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "mycli",
		Short: "My CLI application",
		Run: func(cmd *cobra.Command, args []string) {
			if err := cobra.MaximumNArgs(1)(cmd, args); err != nil {
				fmt.Println("Please enter valid data")
				os.Exit(1)
			}
			// Your command logic here
			fmt.Println("Valid data received:", args)
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
