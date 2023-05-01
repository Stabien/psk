/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"log"
	"os/exec"
	"psk/utils"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var stacks = map[string]string{
	"React.js": "reactjs",
	"Node.js": "nodejs",
}

func chooseStack() (string, error) {
	prompt := promptui.Select{
		Label: "Select your stack",
		Items: []string{"React.js", "Node.js"},
	}

	_, result, err := prompt.Run()

	return result, err
}

func getStack(args []string) (string, error) {
	if (len(args) == 0) {
		return chooseStack()
	} else {
		userInput := strings.ToLower(args[0])
		if (!utils.IsInMap(stacks, userInput)) {
			return chooseStack()
		} else {
			return userInput, nil
		}
	}
}

func getScriptPath(stackName string) string {
	for key := range stacks {
		if key == stackName {
			stackName = stacks[key]
		}
	}
	return "/mnt/d/Projets/psk-cli/scripts/" + stackName + ".sh"
}

func getProjectName() (string, error) {
	validate := func(input string) error {
		if len(input) == 0 {
			return errors.New("Enter a project name")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Project name",
		Validate: validate,
	}

	return prompt.Run()
}

func callScript(stackName string, projectName string) {
	
	path := getScriptPath(stackName)
	command := exec.Command("bash", path, projectName)
	command.Dir = "."
	output, err := command.Output()

	if err != nil {
		log.Println(err)
	}
	fmt.Printf("%s", output)
}

func createCommand(cmd *cobra.Command, args []string) {
	stack, err := getStack(args)

	if err != nil {
		log.Fatal(err)
	}

	projectName, err := getProjectName()

	if err != nil {
		log.Fatal(err)
	}

	callScript(stack, projectName)
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates create starter kit project",
	Long:  ``,
	Run: createCommand,
}

func init() {
	rootCmd.AddCommand(createCmd)
}
