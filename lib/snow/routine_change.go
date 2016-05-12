package snow

import "time"

type Username string
type CI string

type DeploymentMethod struct {
	ID   string
	Type string
}

type Change struct {
	Template  string `json:"template"`
	Type      string
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
	chg Change
}

func (rc *routineChange) GetChange() string {
	return " "
}

func (rc *routineChange) setChange(chg Change) {
	rc.chg = chg
}

func NewRoutineChange(chg Change) routineChange {
	rc := routineChange{}
	rc.setChange(chg)
	return rc
}
