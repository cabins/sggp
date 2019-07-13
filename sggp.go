package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"

	"fmt"

	"github.com/olekukonko/tablewriter"
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

func main() {
	// Define the cmd args
	var sort, order, lan string
	flag.StringVar(&sort, "sort", "stars", "排序规则，取值stars,forks")
	flag.StringVar(&order, "order", "desc", "正向或反向排序，取值desc,asc")
	flag.StringVar(&lan, "lan", "go", "检索编程语言范围，默认go")
	flag.Parse()

	q := flag.Args()[0]

	api := "https://api.github.com/search/repositories?q=" + q + "+language:" + lan + "&sort=" + sort + "&order=" + order
	resp, err := http.Get(api)
	if err != nil {
		log.Fatalln("查询Github出错：" + err.Error())
		return
	}

	defer resp.Body.Close()

	var repos Repos

	json.NewDecoder(resp.Body).Decode(&repos)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Stars", "Forks", "Name", "Url", "Language", "License", "Description"})

	length := len(repos.Items)
	if length > 10 {
		length = 10
	}

	for _, item := range repos.Items[0:length] {
		var desc string
		if len(item.Description) > 30 {
			desc = item.Description[0:30] + "..."
		} else {
			desc = item.Description
		}

		table.Append([]string{fmt.Sprintf("%d", item.StarGazersCount), fmt.Sprintf("%d", item.ForksCount), item.Name, item.HtmlUrl, item.Language, item.License.Name, desc})
	}

	table.Render()
}
