package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	GetEnvConfig()
	dbUser := viper.Get("DB_USER")
	dbPassword := viper.Get("DB_PASSWORD")
	dbHost := viper.Get("DB_HOST")
	dbName := viper.Get("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?tls=true&charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("Current database connection: %s", dbName)

	return db, nil
}
