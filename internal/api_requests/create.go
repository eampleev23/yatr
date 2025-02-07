package api_requests

import (
	"bytes"
	"fmt"
	"github.com/eampleev23/yatr/internal/client_config"
	"github.com/eampleev23/yatr/internal/models"
	"net/http"
	url2 "net/url"
)

func Create(c *client_config.Config, newIssueModel models.NewIssue) error {

	url, err := url2.JoinPath("https://api.tracker.yandex.net/", "v2/issues/")
	if err != nil {
		return fmt.Errorf("url2.JoinPath failed %w", err)
	}
	fmt.Println("c", c)
	fmt.Println("url:", url)

	jsonDataStr := `{
  "queue": "` + newIssueModel.Queue + `",
  "summary": "` + newIssueModel.Summary + `",
  "type": "` + newIssueModel.Type + `",
  "project": ` + newIssueModel.Project + `,
  "start": "` + newIssueModel.Start + `",
  "dueDate": "` + newIssueModel.DueDate + `",
  "description": "` + newIssueModel.Description + `",
  "assignee": "` + newIssueModel.Assignee + `",
  "author": "` + newIssueModel.Author + `",
  "priority": "` + newIssueModel.Priority + `"
}`
	fmt.Println("jsonDataStr", jsonDataStr)
	jsonData := []byte(jsonDataStr)

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Ошибка формирования запроса, попробуйте обновить клиент")
		return fmt.Errorf("http.NewRequest failed %w", err)
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	request.Header.Set("Authorization", "OAuth "+c.YTrToken)
	request.Header.Set("X-Cloud-Org-Id", c.CloudOrgId)

	response, err := c.HttpClient.Do(request)
	if err != nil {
		fmt.Println("Ошибка получения ответа, обратитесь к администратору")
		return fmt.Errorf("c.HttpClient.Do failed %w", err)
	}
	fmt.Println("status code =", response.StatusCode)
	return nil
}
