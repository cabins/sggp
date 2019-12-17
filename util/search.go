package util

import (
	"encoding/json"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"log"
	"net/http"
	"os"
)

type Repos struct {
	Items []Item `json:"items,omitempty"`
}

type Item struct {
	Name            string  `json:"name"`
	HtmlUrl         string  `json:"full_name"`
	Description     string  `json:"description,omitempty"`
	Language        string  `json:"language,omitempty"`
	StarGazersCount int     `json:"stargazers_count,omitempty"`
	ForksCount      int     `json:"forks_Count,omitempty"`
	License         License `json:"license,omitempty"`
}

type License struct {
	Name string `json:"name,omitempty"`
}


func (repos Repos) PrintAsTable() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Stars", "Forks", "Name", "Url", "Language", "License", "Description"})

	for _, repo := range repos.Items {
		end := 30
		if len(repo.Description) < end {
			end = len(repo.Description)
		}

		table.Append([]string{fmt.Sprintf("%d", repo.StarGazersCount),
			fmt.Sprintf("%d", repo.ForksCount),
			repo.Name,
			repo.HtmlUrl,
			repo.Language,
			repo.License.Name,
			repo.Description[0:end],
		})
	}

	table.Render()
}


func SearchRepos(name, sort, order, language string) Repos {
	api := fmt.Sprintf("https://api.github.com/search/repositories?q=%s+language:%s&sort=%s&order=%s", name, language, sort, order)
	resp, err := http.Get(api)
	defer resp.Body.Close()

	if err != nil {
		log.Fatalln("查询Github出错：" + err.Error())
		return Repos{}
	}


	var repos Repos

	err = json.NewDecoder(resp.Body).Decode(&repos)
	if err != nil {
		log.Fatalln(err)
		return Repos{}
	}

	return repos
}