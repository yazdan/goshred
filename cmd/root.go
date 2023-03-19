/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/yazdan/goshred/internal"
)

// rootCmd represents the base command when called without any subcommands
var (
	rootCmd = &cobra.Command{
		Use:   "goshred filename",
		Short: "A simple file shreder in go",
		Long: `This is a simple file shreder written in go:

This tool will overwrite your file 3 times with random content then deletes it`,
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Inside rootCmd Run with args: %v\n", args)
			// Call the function here
			// fmt.Printf("Inside toggle: %t\n", toggle)
			fmt.Printf("file size %d\n", internal.GetFileSize(args[0]))
			fmt.Printf("wiping %s %t\n", args[0], internal.Shred(args[0]))
		},
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
}
