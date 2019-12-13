/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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

var labelsListCmd = &cobra.Command{
	Use:   "list [org-name/repo-name]",
	Short: "lists the labels on the repository",
	Long:  `lists the labels on the repository`,
	Args:  validateRepoNames(),
	Run: func(cmd *cobra.Command, args []string) {
		owner, repo := splitRepoName(args[0])
		url := octokit.RepoLabelsURL
		labels, result := client.Labels().All(&url, octokit.M{"owner": owner, "repo": repo})
		if result.HasError() {
			return
		}
		for _, label := range labels {
			fmt.Println(label.Name)
		}
	},
}

func init() {
	labelsCmd.AddCommand(labelsListCmd)
}
