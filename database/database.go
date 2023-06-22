package database

import (
	"fmt"
	"log"
	"os"

	"github.com/dimasbayuseno/k-style-test/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func ConnectDb() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&models.Member{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(&models.Product{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(&models.ReviewProduct{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(&models.LikeReview{})
	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}
