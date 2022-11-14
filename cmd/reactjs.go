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

// reactjsCmd represents the reactjs command
var reactjsCmd = &cobra.Command{
	Use:   "reactjs",
	Short: "Creates ReactJS starter kit project",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Creating ReactJS project...")

		folderName, _ := cmd.Flags().GetString("name")

		command := exec.Command("bash", "/mnt/d/Projets/psk-cli/scripts/reactjs.sh", folderName)
		command.Dir = "."
		output, err := command.Output()

		if err != nil {
			log.Println(err)
		}
		fmt.Printf("%s", output)
	},
}

func init() {
	rootCmd.AddCommand(reactjsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// reactjsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// reactjsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
