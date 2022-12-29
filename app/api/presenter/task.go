package presenter

type Task struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status" gorm:"index"`
}
