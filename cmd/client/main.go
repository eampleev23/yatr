package main

import (
	"fmt"
	"github.com/eampleev23/yatr/internal/client_config"
	"github.com/eampleev23/yatr/internal/my_csv"
	"log"
	"strconv"
)

type NewIssue struct {
	Queue       string `json:"queue"`
	Summary     string `json:"summary"`
	Type        string `json:"type"`
	Project     int    `json:"project"`
	Start       string `json:"start"`
	DueDate     string `json:"due_date"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Assignee    string `json:"assignee"`
	Parent      string `json:"parent"`
}

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	c := client_config.NewConfig()
	//if err := api_requests.Create(c); err != nil {
	//	return fmt.Errorf("api_requests.Create failed: %w", err)
	//}
	result := my_csv.CsvParse(c.FilePath)

	var newIssues []NewIssue
	for i := 1; i < len(result); i++ {
		newIssues = append(newIssues, NewIssue{})
		newIssues[i-1].Queue = result[i][2]
		newIssues[i-1].Summary = result[i][3]
		newIssues[i-1].Type = result[i][4]
		prj, err := strconv.Atoi(result[i][5])
		if err != nil {
			return fmt.Errorf("could not convert project number to int")
		}
		newIssues[i-1].Project = prj
		newIssues[i-1].Start = result[i][6]
		newIssues[i-1].DueDate = result[i][7]
		newIssues[i-1].Description = result[i][8]
		newIssues[i-1].Assignee = result[i][9]
		newIssues[i-1].Author = result[i][10]
		newIssues[i-1].Parent = result[i][11]
	}
	return nil
}
