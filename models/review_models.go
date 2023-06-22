package models

import "time"

type ReviewProduct struct {
	ID         uint       `gorm:"primary_key"`
	ProductID  uint       `gorm:"column:product_id; not null"`
	MemberID   uint       `gorm:"column:member_id; not null"`
	DescReview string     `gorm:"column:desc_review; not null"`
	CreatedAt  time.Time  `gorm:"column:created_at"`
	UpdatedAt  time.Time  `gorm:"column:updated_at"`
	DeletedAt  *time.Time `gorm:"column:deleted_at"`
}
