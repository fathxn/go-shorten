package config

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func NewDB(cfg Config) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// func InitDB() *gorm.DB {
// 	username := viper.GetString("DB_USER")
// 	password := viper.GetString("DB_PASS")
// 	host := viper.GetString("DB_HOST")
// 	port := viper.GetString("DB_PORT")
// 	dbName := viper.GetString("DB_NAME")

// 	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True&loc=Local", username, password, host, port, dbName)
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatalf("error connecting to database: %s", err)
// 	}
// 	return db
// }
