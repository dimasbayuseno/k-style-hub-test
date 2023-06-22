package models

import "time"

type Member struct {
	ID        uint       `gorm:"column:id;primaryKey"`
	Username  string     `gorm:"column:username;not null"`
	Gender    string     `gorm:"column:gender;not null"`
	SkinType  string     `gorm:"column:skin_type;not null"`
	SkinColor string     `gorm:"column:skin_color;not null"`
	CreatedAt time.Time  `gorm:"column:created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at"`
}
