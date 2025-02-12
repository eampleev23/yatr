package api_requests

import (
	"bytes"
	"encoding/json"
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
		return fmt.Errorf("c.HttpClient.Do failed %w", err)
	}

	if response.StatusCode != http.StatusCreated {
		return fmt.Errorf("response status code not 201, but: %d", response.StatusCode)
	}
	fmt.Println(newIssueModel.Type, newIssueModel.Summary, "успешно создана..")
	// Получаем в ответ ID созданной таски
	var newIssueResponse models.NewIssueResponse
	var buf bytes.Buffer
	// читаем тело запроса
	_, err = buf.ReadFrom(response.Body)
	if err != nil {
		return fmt.Errorf("buf.ReadFrom failed %w", err)
	}
	// десериализуем JSON в newIssueResponse
	if err = json.Unmarshal(buf.Bytes(), &newIssueResponse); err != nil {
		return fmt.Errorf("json.Unmarshal failed %w", err)
	}
	fmt.Println("createdKey= ", newIssueResponse.Key)
	return nil
}
