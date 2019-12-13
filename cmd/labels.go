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
	"github.com/octokit/go-octokit/octokit"
	"github.com/spf13/cobra"
)

var DefaultLabels = []octokit.Label{
	{Name: "In Development", Color: "030d7c"},
	{Name: "Ready for Review", Color: "eeb916"},
	{Name: "Review Remediation", Color: "fef2c0"},
	{Name: "Ready for Merge", Color: "1b7c03"},
	{Name: "Blocker", Color: "ee0701"},
	{Name: "No Release Email", Color: "bfdadc"},
}

var labelsCmd = &cobra.Command{
	Use:   "labels",
	Short: "Manipulate labels on a repository",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

func init() {
	rootCmd.AddCommand(labelsCmd)
}
