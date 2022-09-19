package model

type Task struct {
	Name      string `gorm:"primary_key" json:"id"`
	Spec      string `gorm:"type:varchar(255)" json:"spec"`
	Type      string `gorm:"type:varchar(255)" json:"type"`
	Attribute string `gorm:"type:varchar(255)" json:"attribute"`
	UserID    uint   `gorm:"type:integer;not null" json:"user_id"`
}
