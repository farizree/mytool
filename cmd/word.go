/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"strings"

	"github.com/spf13/cobra"
)

// wordCmd represents the word command
var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "A brief description of your command",
	Long:  `A longer description.`,
	Run: func(cmd *cobra.Command, args []string) {
		var result string
		switch mode {
		case "lower":
			result = strings.ToLower(str)
		case "upper":
			result = strings.ToUpper(str)
		default:
			result = "please input mode"
		}
		fmt.Printf(result)
	},
}

var str, mode string

func init() {
	rootCmd.AddCommand(wordCmd)
	wordCmd.Flags().StringVar(&str, "string", "default string", "input your wording")
	wordCmd.Flags().StringVar(&mode, "mode", "lower", "choose mode like upper, lower")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// wordCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// wordCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
