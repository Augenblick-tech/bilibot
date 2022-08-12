package model

type Dynamic struct {
	DynamicID string `gorm:"type:text;not null;primary_key"`
	PubTS     uint64 `gorm:"type:integer;not null"`
	Content   string `gorm:"type:text;not null"`
	AuthorID  uint   `gorm:"type:integer;not null"`
	Author    Author `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
