package task

type Task struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
}
