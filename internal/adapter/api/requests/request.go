package requests

type StandUpdateRequest struct {
	Previous  string `json:"previous"`
	Current   string `json:"current"`
	Blocked   bool   `json:"blocked"`
	StandupID string `json:"standup_id"`
	SprintID  string `json:"sprint_id"`
	Standup   string `json:"standup"`
	UserID    string `json:"user_id"`
	BreakAway string `json:"break_away"`
}

type TaskRequest struct {
	SprintID    string `json:"sprint_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type SprintRequest struct {
	Name string `json:"name"`
}
