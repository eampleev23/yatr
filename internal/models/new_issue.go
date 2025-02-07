package models

type NewIssue struct {
	Queue       string `json:"queue"`
	Summary     string `json:"summary"`
	Type        string `json:"type"`
	Project     string `json:"project"`
	Start       string `json:"start"`
	DueDate     string `json:"due_date"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Assignee    string `json:"assignee"`
	Parent      string `json:"parent"`
	Priority    string `json:"priority"`
}
