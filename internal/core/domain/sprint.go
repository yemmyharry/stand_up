package domain

type Sprint struct {
	Model
	Name  string `json:"name"`
	Tasks []Task `json:"tasks"`
}

type UserTask struct {
	TaskID string `json:"task_id"`
	UserID string `json:"user_id"`
}

type Task struct {
	Model
	SprintID    string `json:"sprint_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Assigned    []User `json:"user" gorm:"many2many:user_tasks;"`
}
