package models

import "time"

type Product struct {
	ID        uint       `gorm:"column:id;primaryKey"`
	Name      string     `gorm:"column:name;not null"`
	Price     uint       `gorm:"column:price;not null"`
	CreatedAt time.Time  `gorm:"column:created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at"`
}
