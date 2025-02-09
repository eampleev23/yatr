package main

import (
	"fmt"
	"github.com/eampleev23/yatr/internal/api_requests"
	"github.com/eampleev23/yatr/internal/client_config"
	"github.com/eampleev23/yatr/internal/models"
	"github.com/eampleev23/yatr/internal/my_csv"
	"log"
	"unicode/utf8"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	c := client_config.NewConfig()

	if c.IsCreating {
		result := my_csv.CsvParse(c.FilePath)

		var newIssues []models.NewIssue

		for i := 1; i < len(result); i++ {
			newIssues = append(newIssues, models.NewIssue{})
			newIssues[i-1].Queue = result[i][2]
			newIssues[i-1].Summary = result[i][3]
			newIssues[i-1].Type = result[i][4]
			newIssues[i-1].Project = result[i][5]
			newIssues[i-1].Start = trimFirstRune(result[i][6])
			newIssues[i-1].DueDate = trimFirstRune(result[i][7])
			newIssues[i-1].Description = result[i][8]
			newIssues[i-1].Assignee = result[i][9]
			newIssues[i-1].Author = result[i][10]
			newIssues[i-1].Parent = result[i][11]
			newIssues[i-1].Priority = result[i][12]

			if err := api_requests.Create(c, newIssues[i-1]); err != nil {
				return fmt.Errorf("api_requests.Create failed: %w", err)
			}
		}

	}
	return nil
}

func trimFirstRune(s string) string {
	_, i := utf8.DecodeRuneInString(s)
	return s[i:]
}
