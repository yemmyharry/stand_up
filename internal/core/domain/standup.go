package domain

import "time"

const (
	Before string = "before"
	After  string = "after"
	Within string = "Within"
)

type StandUpdate struct {
	Previous  string `json:"previous"`
	Current   string `json:"current"`
	Blocked   bool   `json:"blocked"`
	StandupID string `json:"standup_id"`
	SprintID  string `json:"sprint_id"`
	Standup   string `json:"standup"`
	Status    string `json:"status"`
	UserID    string `json:"user_id"`
	BreakAway string `json:"break_away"`
	Duration  int    `json:"duration"`
	Sprint    Sprint `json:"sprint"`
	Model
}

type Standup struct {
	Model
	Title     string    `json:"title"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}
