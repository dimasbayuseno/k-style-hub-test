package models

type ProductReview struct {
	Product        Product         `gorm:"foreignKey:ID_PRODUCT"`
	Reviews        []ReviewProduct `gorm:"foreignKey:ID_PRODUCT"`
	TotalLikeCount int             `gorm:"-"`
}
