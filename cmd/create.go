/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"log"
	"os/exec"
	"psk/data"
	"psk/utils"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

func getStackLabels() []string {
	var stackLabels []string

	for key := range data.Stacks {
		stackLabels = append(stackLabels, key)
	}
	return stackLabels
}

func selectPrompt(label string, items []string) (string, error) {
	prompt := promptui.Select{
		Label: label,
		Items: items,
	}

	_, result, err := prompt.Run()

	return result, err
}

func getStack(args []string) (string, error) {
	stackLabels := getStackLabels()

	if (len(args) == 0) {
		stackLabel, err := selectPrompt("Select your stack", stackLabels)
		return data.Stacks[stackLabel], err
	} else {
		userInput := strings.ToLower(args[0])
		if (!utils.IsInMap(data.Stacks, userInput)) {
			stackLabel, err := selectPrompt("Select your stack", stackLabels)
			return data.Stacks[stackLabel], err
		} else {
			return userInput, nil
		}
	}
}

func getScriptPath(stackName string) string {
	for key := range data.Stacks {
		if key == stackName {
			stackName = data.Stacks[key]
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

func getOptions(stackName string) ([]string, error) {
	var options []string

	for key, option := range data.StackOptions[stackName] {
		option, err := selectPrompt("Select a " + key, option)
		
		if err != nil {
			return nil, err
		}
		options = append(options, option)
	}
	return options, nil
}

func callScript(stackName string, projectName string, optionList []string) {
	path := getScriptPath(stackName)
	options := strings.Join(optionList, " ")

	command := exec.Command("bash", path, projectName, options)
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

	options, err := getOptions(stack)

	if err != nil {
		log.Fatal(err)
	}

	callScript(stack, projectName, options)
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
