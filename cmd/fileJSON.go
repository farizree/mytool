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
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// fileJSONCmd represents the fileJSON command
var fileJSONCmd = &cobra.Command{
	Use:   "/var/log/nginx/error.log -t",
	Short: "A brief description of your command",
	Long:  `A longer json / text`,
	Run: func(cmd *cobra.Command, args []string) {
		out, err := exec.Command("").Output()
		//Directory log
		err = os.Chdir("/var/log/nginx")
		check(err)

		output := string(out[:])
		fmt.Println(output)

		// Read File
		content, err := ioutil.ReadFile("error.log")
		check(err)
		myString := string(content)

		// Back to working directory
		err = os.Chdir("/Users/farizree/mytools")
		// End Back to working directory
		if err != nil {
			fmt.Printf("%s", err)
		}

		argsString := strings.Join(args, " ")
		fmt.Println(argsString)

		if argsString == "json" {
			file, err := os.Create("error.json")
			if err != nil {
				log.Fatalf("failed creating file: %s", err)
			}
			defer file.Close()

			len, err := file.WriteString(myString)

			if err != nil {
				log.Fatalf("failed writing to file: %s", err)
			}
			fmt.Printf("\nFile Name: %s", file.Name())
			fmt.Printf("\nLength: %d bytes", len)
		} else if argsString == "text" {
			file, err := os.Create("error.text")
			if err != nil {
				log.Fatalf("failed creating file: %s", err)
			}
			defer file.Close()

			len, err := file.WriteString(myString)

			if err != nil {
				log.Fatalf("failed writing to file: %s", err)
			}
			fmt.Printf("\nFile Name: %s", file.Name())
			fmt.Printf("\nLength: %d bytes", len)
		} else {
			fmt.Println(myString)
		}
	},
}

var json string

func init() {
	rootCmd.AddCommand(fileJSONCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fileJSONCmd.PersistentFlags().String("json", "-o", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	fileJSONCmd.Flags().BoolP("convert", "t", false, "convert file")

	// fileJSONCmd.Flags().StringVarP(&json, "json", "t", "", "set bar (*)")
}
