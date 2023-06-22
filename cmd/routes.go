package main

import (
	"github.com/dimasbayuseno/k-style-test/handlers"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func registerRoutes(e *echo.Echo, db *gorm.DB) {
	e.GET("/members", handlers.GetMembers(db))
	e.GET("/members/:id", handlers.GetMember(db))
	e.POST("/members", handlers.CreateMember(db))
	e.PUT("/members/:id", handlers.UpdateMember(db))
	e.DELETE("/members/:id", handlers.DeleteMember(db))

	e.GET("/products/:id", handlers.GetProduct(db))
	e.GET("/products", handlers.DeleteProduct(db))
	e.POST("/products", handlers.CreateProduct(db))
	e.PUT("/products/:id", handlers.UpdateProduct(db))
	e.DELETE("/products/:id", handlers.DeleteProduct(db))

	e.GET("/product-reviews/:id", handlers.GetProductWithReview(db))
	e.POST("/product-reviews", handlers.CreateReviewProduct(db))

	e.POST("/reviews/:id_review/members/:id_member/like", handlers.LikeReview(db))
	e.DELETE("/reviews/:id_review/members/:id_member/like", handlers.CancelLikeReview(db))

}
