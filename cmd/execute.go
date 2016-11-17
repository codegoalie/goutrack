// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

// executeCmd represents the execute command
var executeCmd = &cobra.Command{
	Use:   "execute",
	Short: "Execute a YouTrack command on an issue",
	Long: `Apply command strings to a story. A command string will modify one
or more attributes. Additionally, any trailing text will be added as a comment
to the story. For example:

goutrack execute yt-1234 "Assignee me" "I'll take this story."
goutrack exec yt-4321 "Subsystem server-side"
goutrack e yt-3214 "Accepted Add Reviewed by me"
goutrack c yt-4312 "Add scheduled for sprint-2" "Sliding to next sprint"
`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		log.Printf("execute called as %+v", cmd)
	},
}

func init() {
	RootCmd.AddCommand(executeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// executeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// executeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
