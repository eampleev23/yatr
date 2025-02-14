package models

type NewIssue struct {
	Key         string `json:"key,omitempty"`
	Queue       string `json:"queue,omitempty"`
	Summary     string `json:"summary,omitempty"`
	Type        string `json:"type,omitempty"`
	Project     int    `json:"project,omitempty"`
	Start       string `json:"start,omitempty"`
	DueDate     string `json:"dueDate,omitempty"`
	Author      string `json:"author,omitempty"`
	Description string `json:"description,omitempty"`
	Assignee    string `json:"assignee,omitempty"`
	Parent      string `json:"parent,omitempty"`
	Priority    string `json:"priority,omitempty"`
}
