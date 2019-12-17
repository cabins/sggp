/*
Copyright © 2019 Cabins <kong_lingcun@163.com>

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
	"github.com/cabins/ghc/util"
	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Aliases: []string{"sr"},
	Short: "search with names and specific conditions",
	Long: `search with names and specific conditions.`,
	Run: func(cmd *cobra.Command, args []string) {
		language, _ := cmd.Flags().GetString("language")
		sort,_ := cmd.Flags().GetString("sort")
		order,_ := cmd.Flags().GetString("order")

		for _, name := range args {
			util.SearchRepos(name, sort, order, language).PrintAsTable()
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	searchCmd.Flags().StringP("language", "l", "go", "specific language, default is go")
	searchCmd.Flags().StringP("sort", "s", "stars", "sort by, default is starts, available option: stars/forks")
	searchCmd.Flags().StringP("order", "o", "desc", "order by, default is starts, available option: desc/asc")
}
