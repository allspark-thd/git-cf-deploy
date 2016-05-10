package snow

import "time"

type Username string
type CI string

type DeploymentMethod struct {
	ID   string
	Type string
}

type Change struct {
	CIs       []CI
	Requestor Username
	Assignee  Username
	Validator string
	StartDate time.Time
	EndDate   time.Time
	DeploymentMethod
	ProjectID string
	Comments  string
	WorkNotes string
}
type routineChange struct {
}

func NewRoutineChange(template string, chg Change) routineChange {
	return routineChange{}
}
