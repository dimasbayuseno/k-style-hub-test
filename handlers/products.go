package handlers

import (
	"github.com/dimasbayuseno/k-style-test/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

func GetProducts(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var products []models.Product
		result := db.Find(&products)
		if result.Error != nil {
			return c.JSON(http.StatusInternalServerError, result.Error)
		}
		return c.JSON(http.StatusOK, products)
	}
}

func GetProduct(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		productID := c.Param("id")

		product := new(models.Product)
		if err := db.First(product, productID).Error; err != nil {
			return c.JSON(http.StatusNotFound, err.Error())
		}

		return c.JSON(http.StatusOK, product)
	}
}

func CreateProduct(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		product := new(models.Product)
		if err := c.Bind(product); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		now := time.Now()
		product.CreatedAt = now
		product.UpdatedAt = now

		if err := db.Create(product).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, product)
	}
}

func UpdateProduct(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		product := new(models.Product)
		if err := c.Bind(product); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		product.ID = uint(id)

		result := db.Save(product)
		if result.Error != nil {
			return c.JSON(http.StatusInternalServerError, result.Error)
		}

		return c.JSON(http.StatusOK, product)
	}
}

func DeleteProduct(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		productID := c.Param("id")

		product := new(models.Product)
		if err := db.First(product, productID).Error; err != nil {
			return c.JSON(http.StatusNotFound, err.Error())
		}

		if err := db.Delete(product).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.NoContent(http.StatusNoContent)
	}
}
