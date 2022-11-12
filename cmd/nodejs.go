/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os/exec"
)

// nodejsCmd represents the nodejs command
var nodejsCmd = &cobra.Command{
	Use:   "nodejs",
	Short: "Creates NodeJS starter kit project",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Creating NodeJS project...")
		var command *exec.Cmd

		if len(args) > 0 {
			command = exec.Command("bash", "/mnt/d/Projets/psk-cli/scripts/nodejs.sh", args[0])
		} else {
			command = exec.Command("bash", "/mnt/d/Projets/psk-cli/scripts/nodejs.sh")
		}

		command.Dir = "."
		output, err := command.Output()

		if err != nil {
			log.Println(err)
		}
		fmt.Printf("%s", output)
	},
}

func init() {
	rootCmd.AddCommand(nodejsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// nodejsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// nodejsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}