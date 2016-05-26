package snow

import "time"

// Username ...
type Username string

// CI ...
type CI string

// DeploymentMethod ...
type DeploymentMethod struct {
	ID   string
	Type string
}

// Change ...
type Change struct {
	Template  string `json:"template"`
	Type      string `json:"type"`
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

// RoutineChange ...
type RoutineChange struct {
	chg Change
}

// GetChange ...
func (rc *RoutineChange) GetChange() Change {
	return rc.chg
}

// SetChange ...
func (rc *RoutineChange) SetChange(chg Change) {
	rc.chg = chg
}

// NewRoutineChange ...
func NewRoutineChange(chg Change) RoutineChange {
	rc := RoutineChange{}
	rc.SetChange(chg)
	return rc
}
