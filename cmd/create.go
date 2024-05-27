package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
    Use:   "create",
    Short: "Create a new item",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Creating a new item")
    },
}

