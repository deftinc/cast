/*
Copyright © 2019 Patrick Wiseman <patrick.wiseman@deft.services>

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
	"github.com/octokit/go-octokit/octokit"
	"github.com/spf13/cobra"
)

var labelsCreateCmd = &cobra.Command{
	Use:   "create [org-name/repo-name ...]",
	Short: "Creates the Deft labels on the repository",
	Long:  `Creates the Deft labels on the repository`,
	Args:  validateRepoNames(),
	Run: func(cmd *cobra.Command, args []string) {
		url := octokit.RepoLabelsURL
		for _, repoName := range args {
			owner, repo := splitRepoName(repoName)
			for _, label := range DefaultLabels {
				_, result := client.Labels().Create(
					&url,
					octokit.M{"owner": owner, "repo": repo},
					label,
				)
				if result.HasError() {
					fmt.Println(result.Error())
					return
				}
				fmt.Print(".")
			}
		}
		fmt.Println("done")
	},
}

func init() {
	labelsCmd.AddCommand(labelsCreateCmd)
}
