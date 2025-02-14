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

type Type struct {
	Id  string `json:"id,omitempty"`
	Key string `json:"key,omitempty"`
}
type Parent struct {
	Id  string `json:"id,omitempty"`
	Key string `json:"key,omitempty"`
}

type Priority struct {
	Id  string `json:"id,omitempty"`
	Key string `json:"key,omitempty"`
}

type UpdateRequest struct {
	Summary     string   `json:"summary,omitempty"`
	Description string   `json:"description,omitempty"`
	Type        Type     `json:"type,omitempty"`
	Priority    Priority `json:"priority,omitempty"`
	Parent      Parent   `json:"parent,omitempty"`
	Assignee    string   `json:"assignee,omitempty"`
}

func Update(c *client_config.Config, updateIssueModel models.NewIssue) error {
	url, err := url2.JoinPath("https://api.tracker.yandex.net/", "v2/issues/", updateIssueModel.Key)
	if err != nil {
		return fmt.Errorf("url2.JoinPath failed %w", err)
	}

	req := UpdateRequest{
		Summary:     updateIssueModel.Summary,
		Description: updateIssueModel.Description,
		Type: Type{
			Key: updateIssueModel.Type,
		},
		Priority: Priority{
			Key: updateIssueModel.Priority,
		},
		Parent: Parent{
			Key: updateIssueModel.Parent,
		},
		Assignee: updateIssueModel.Assignee,
	}

	jsonData, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("json.Marshal failed %w", err)
	}
	request, err := http.NewRequest("PATCH", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("http.NewRequest failed %w", err)
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	request.Header.Set("Authorization", "OAuth "+c.YTrToken)
	request.Header.Set("X-Cloud-Org-Id", c.CloudOrgId)

	response, err := c.HttpClient.Do(request)
	if err != nil {
		return fmt.Errorf("c.HttpClient.Do failed %w", err)
	}
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("response status code not 200, but: %d", response.StatusCode)
	}
	fmt.Println(updateIssueModel.Key, "успешно обновлена")
	return nil
}
