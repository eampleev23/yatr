package services

import (
	"fmt"
	"github.com/eampleev23/yatr/internal/api_requests"
	"github.com/eampleev23/yatr/internal/client_config"
	"github.com/eampleev23/yatr/internal/models"
	"github.com/eampleev23/yatr/internal/my_csv"
	"unicode/utf8"
)

func GenerateIssues(c *client_config.Config) error {
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

		createdKey, err := api_requests.Create(c, newIssues[i-1])
		if err != nil {
			return fmt.Errorf("create issues: %w", err)
		}
		newIssues[i-1].Key = createdKey
		result[i][1] = createdKey
	}
	err := my_csv.CsvSave(c.FilePath, result)
	if err != nil {
		return fmt.Errorf("save issues: %w", err)
	}
	return nil
}

func trimFirstRune(s string) string {
	_, i := utf8.DecodeRuneInString(s)
	return s[i:]
}
