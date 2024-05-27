package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
    Use:   "list",
    Short: "List all items",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Listing all items")
    },
}

