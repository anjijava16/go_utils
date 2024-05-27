package cmd

import (
    "github.com/spf13/cobra"
    "fmt"
)

var rootCmd = &cobra.Command{
    Use:   "serving",
    Short: "Main command for serving",
    Long:  "A CLI application to manage serving tasks with subcommands for listing, creating, updating, and deleting.",
}

func Execute() error {
    return rootCmd.Execute()
}

func init() {
    // Add subcommands to the root command
    rootCmd.AddCommand(listCmd)
    var res string
    res=createCmd.Name()
    fmt.Println(res)
    rootCmd.AddCommand(createCmd)
    rootCmd.AddCommand(updateCmd)
    rootCmd.AddCommand(deleteCmd)
}

