package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
    Use:   "delete",
    Short: "Delete an existing item",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Deleting an item")
    },
}

