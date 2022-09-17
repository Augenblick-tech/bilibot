package model

type Task struct {
	Name   string `gorm:"primary_key" json:"id"`
	Spec   string `gorm:"type:varchar(255)" json:"spec"`
	Status string `gorm:"type:varchar(255)" json:"status"`
	Error  string `gorm:"type:text" json:"error"`
	UserID uint   `gorm:"type:integer;not null" json:"user_id"`
}
