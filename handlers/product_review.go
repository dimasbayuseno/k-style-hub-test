package handlers

import (
	"github.com/dimasbayuseno/k-style-test/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func GetProductWithReview(db *gorm.DB) echo.HandlerFunc {
	type ReviewData struct {
		Username         string `json:"username"`
		Gender           string `json:"gender"`
		SkinType         string `json:"skinType"`
		SkinColor        string `json:"skinColor"`
		DescReview       string `json:"descReview"`
		JumlahLikeReview int    `json:"jumlahLikeReview"`
	}

	type ProductWithReview struct {
		models.Product
		Reviews []ReviewData `json:"reviews"`
	}

	return func(c echo.Context) error {
		productID := c.Param("id")

		product := new(models.Product)
		if err := db.First(product, productID).Error; err != nil {
			return c.JSON(http.StatusNotFound, err.Error())
		}

		var reviews []ReviewData
		if err := db.Table("review_products").
			Select("members.username, members.gender, members.skin_type, members.skin_color, review_products.desc_review, COUNT(like_reviews.review_id) as jumlah_like_review").
			Joins("JOIN members ON members.id = review_products.member_id").
			Joins("LEFT JOIN like_reviews ON like_reviews.review_id = review_products.id").
			Where("review_products.product_id = ?", productID).
			Group("members.username, members.gender, members.skin_type,  members.skin_color,  review_products.desc_review").
			Find(&reviews).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		productWithReview := ProductWithReview{
			Product: *product,
			Reviews: reviews,
		}

		return c.JSON(http.StatusOK, productWithReview)
	}
}

func CreateReviewProduct(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Parse request body
		review := new(models.ReviewProduct)
		if err := c.Bind(review); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		now := time.Now()
		review.CreatedAt = now
		review.UpdatedAt = now

		if err := db.Create(review).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, review)
	}
}
