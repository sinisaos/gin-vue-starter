package models

type Task struct {
	ID        uint64 `json:"id" gorm:"primary_key"`
	Name      string `json:"name"`
	Completed bool   `json:"completed" gorm:"default:false"`
	UserID    uint64 `json:"user_id"`
	User      User   `json:"user,omitempty"`
}

type TaskValidate struct {
	Name      string `json:"name" binding:"required"`
	Completed bool   `json:"completed" gorm:"default:false"`
	UserID    uint64 `json:"user_id"`
}
