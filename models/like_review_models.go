package models

import "time"

type LikeReview struct {
	ReviewID  uint       `gorm:"column:review_id;primaryKey"`
	MemberID  uint       `gorm:"column:member_id;primaryKey"`
	CreatedAt time.Time  `gorm:"column:created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at"`
}
