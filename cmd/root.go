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
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/codeschooldropout/gophertask/data"
	"github.com/codeschooldropout/gophertask/ui"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gophertask",
	Short: "Use gophertask to manage your task and time allotment",
	Long:  `Modeled after pomodoro, gophertask is a tool for managing your time and tasks.`,
	Run: func(cmd *cobra.Command, args []string) {

		// this is the default command if no other is specified
		p := tea.NewProgram(ui.InitialModel(data.TaskList()))
		if err := p.Start(); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

	},
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
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
