package services

import (
	"fmt"
	"github.com/eampleev23/yatr/internal/api_requests"
	"github.com/eampleev23/yatr/internal/client_config"
	"github.com/eampleev23/yatr/internal/models"
	"github.com/eampleev23/yatr/internal/my_csv"
)

func UpdateIssues(c *client_config.Config) error {
	result := my_csv.CsvParse(c.FilePath)
	var updateIssues []models.NewIssue
	for i := 1; i < len(result); i++ {
		updateIssues = append(updateIssues, models.NewIssue{})
		updateIssues[i-1].Key = result[i][1]
		updateIssues[i-1].Summary = result[i][3]
		updateIssues[i-1].Type = result[i][4]
		updateIssues[i-1].Project = result[i][5]
		updateIssues[i-1].DueDate = trimFirstRune(result[i][7])
		updateIssues[i-1].Description = result[i][8]
		updateIssues[i-1].Assignee = result[i][9]
		updateIssues[i-1].Author = result[i][10]
		updateIssues[i-1].Parent = result[i][11]
		updateIssues[i-1].Priority = result[i][12]

		err := api_requests.Update(c, updateIssues[i-1])
		if err != nil {
			return fmt.Errorf("update issues: %w", err)
		}
	}
	return nil
}
