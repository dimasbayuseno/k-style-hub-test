package handlers

import (
	"github.com/dimasbayuseno/k-style-test/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func GetMembers(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var members []models.Member
		result := db.Find(&members)
		if result.Error != nil {
			return c.JSON(http.StatusInternalServerError, result.Error)
		}
		return c.JSON(http.StatusOK, members)
	}
}

func GetMember(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		productID := c.Param("id")

		product := new(models.Member)
		if err := db.First(product, productID).Error; err != nil {
			return c.JSON(http.StatusNotFound, err.Error())
		}

		return c.JSON(http.StatusOK, product)
	}
}

func CreateMember(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		member := new(models.Member)
		if err := c.Bind(member); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		result := db.Create(member)
		if result.Error != nil {
			return c.JSON(http.StatusInternalServerError, result.Error)
		}

		return c.JSON(http.StatusCreated, member)
	}
}

func UpdateMember(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		member := new(models.Member)
		if err := c.Bind(member); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		member.ID = uint(id)

		result := db.Save(member)
		if result.Error != nil {
			return c.JSON(http.StatusInternalServerError, result.Error)
		}

		return c.JSON(http.StatusOK, member)
	}
}

func DeleteMember(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		result := db.Delete(&models.Member{}, id)
		if result.Error != nil {
			return c.JSON(http.StatusInternalServerError, result.Error)
		}

		return c.JSON(http.StatusOK, "Member deleted")
	}
}
