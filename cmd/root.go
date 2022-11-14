/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "psk-cli",
	Short: "A brief description of your application",
	Long:  ``,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	nodejsCmd.PersistentFlags().StringP("name", "n", "template_nodejs", "Name of the folder where the project will be created")
	reactjsCmd.PersistentFlags().StringP("name", "n", "template_reactjs", "Name of the folder where the project will be created")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
