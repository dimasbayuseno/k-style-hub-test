package handlers

import (
	"github.com/dimasbayuseno/k-style-test/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func LikeReview(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		idReview, err := strconv.Atoi(c.Param("id_review"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		idMember, err := strconv.Atoi(c.Param("id_member"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		like := models.LikeReview{
			ReviewID: uint(idReview),
			MemberID: uint(idMember),
		}

		result := db.Create(&like)
		if result.Error != nil {
			return c.JSON(http.StatusInternalServerError, result.Error)
		}

		return c.JSON(http.StatusOK, like)
	}
}

func CancelLikeReview(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		idReview, err := strconv.Atoi(c.Param("id_review"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		idMember, err := strconv.Atoi(c.Param("id_member"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		result := db.Where("id_review = ? AND id_member = ?", idReview, idMember).Delete(&models.LikeReview{})
		if result.Error != nil {
			return c.JSON(http.StatusInternalServerError, result.Error)
		}

		return c.JSON(http.StatusOK, "Like canceled")
	}
}
