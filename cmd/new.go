/*
Copyright © 2022 codeschooldropout <code@cay.io>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/codeschooldropout/gophertask/data"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Creates a new gophertask task",
	Long:  `Creates a new gophertask task. Needs title, description, category and status.`,
	Run: func(cmd *cobra.Command, args []string) {
		createNewTask()
	},
}

type promptContent struct {
	errorMsg string
	label    string
}

func init() {
	taskCmd.AddCommand(newCmd)

}

func promptGetInput(pc promptContent) string {
	validate := func(input string) error {
		if len(input) <= 0 {
			return errors.New(pc.errorMsg)
		}
		return nil
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }}",
		Valid:   "{{ . | green }}",
		Invalid: "{{ . | red }}",
		Success: "{{ . | bold }}",
	}

	prompt := promptui.Prompt{
		Label:     pc.label,
		Validate:  validate,
		Templates: templates,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}

	fmt.Printf("Input: %s\n", result)
	return result
}

func promptGetSelect(pc promptContent) string {
	items := []string{"home", "work", "code", "project"}
	index := -1

	var result string
	var err error
	for index < 0 {
		prompt := promptui.SelectWithAdd{
			Label:    pc.label,
			Items:    items,
			AddLabel: "Add a new category",
		}

		index, result, err = prompt.Run()

		if index == -1 {
			items = append(items, result)
		}
	}

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Input: %s\n", result)
	return result

}

func createNewTask() {
	taskPromtContent := promptContent{
		"Please enter a title for your task",
		"What are you trying to accomplish? ",
	}
	task := promptGetInput(taskPromtContent)

	descriptionPromtContent := promptContent{
		"Please enter a description for your task",
		fmt.Sprintf("What is the description of %s? ", task),
	}
	description := promptGetInput(descriptionPromtContent)

	categoryPromtContent := promptContent{
		"Please enter a category",
		fmt.Sprintf("What category is %s in?", task),
	}
	category := promptGetSelect(categoryPromtContent)

	data.InsertTask(task, description, category, "new")

}
