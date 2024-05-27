package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
    Use:   "update",
    Short: "Update an existing item",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Updating an item")
    },
}

