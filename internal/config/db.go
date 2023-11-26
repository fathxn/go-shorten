package config

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func InitDB() *gorm.DB {
	username := viper.GetString("DB_USER")
	password := viper.GetString("DB_PASS")
	host := viper.GetString("DB_HOST")
	port := viper.GetString("DB_PORT")
	dbName := viper.GetString("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True&loc=Local", username, password, host, port, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("error connecting to database: %s", err)
	}
	return db
}
