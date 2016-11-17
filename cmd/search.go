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
	"fmt"
	"log"
	"strings"

	"github.com/codegoalie/goutrack/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Query for YouTrack issues",
	Long: `Provide the same search as you would in the YouTrack search bar to retrieve
stories. For example:

goutrack search "Assignee: me"
goutrack search "Type: Bug #Unresolved"
`,
	Run: func(cmd *cobra.Command, args []string) {
		terms := strings.Join(args, " ")
		client := client.NewYouTrackClient(viper.GetString("host"), viper.GetString("username"), viper.GetString("password"))
		issue, err := client.SearchIssues(terms)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(issue)
	},
}

func init() {
	RootCmd.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
