/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

type stack struct {
	label string
	value string
}

func rootCommand(cmd *cobra.Command, args []string) {
	prompt := promptui.Select{
		Label: "Select your stack",
		Items: []string{"React.js", "Node.js"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}
	
	fmt.Printf("You choose %q\n", result)
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "psk",
	Short: "A brief description of your application",
	Long:  ``,
	Run: rootCommand,
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
