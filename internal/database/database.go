package database

import (
	"e-ticketing-api/config"
	"e-ticketing-api/internal/models"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	cfg := config.AppConfig
	dbinfo := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBName, cfg.DBPassword)

	DB, err = gorm.Open("postgres", dbinfo)
	if err != nil {
		log.Fatalf("Gagal terhubung ke database: %v", err)
	}

	log.Println("Koneksi database berhasil.")

	DB.AutoMigrate(&models.Admin{}, &models.Terminal{})
	log.Println("Migrasi database selesai.")
}