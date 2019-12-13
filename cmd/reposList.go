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

var reposList = &cobra.Command{
	Use:   "list",
	Short: "List repositories",
	Run: func(cmd *cobra.Command, args []string) {
		url := octokit.UserRepositoriesURL
		repos, result := client.Repositories().All(&url, octokit.M{"type": "all"})
		if result.HasError() {
			fmt.Println(result.Error())
		}
		for _, repo := range repos {
			fmt.Println(repo.FullName)
		}
	},
}

func init() {
	reposCmd.AddCommand(reposList)
}
