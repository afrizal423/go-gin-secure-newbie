package configs

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// postgres database konektor
func GormPostgresConn() *gorm.DB {
	configDB := map[string]string{
		"DB_Username": Config("DB_USERNAME"),
		"DB_Password": Config("DB_PASSWORD"),
		"DB_Host":     Config("DB_ADDRESS"),
		"DB_Name":     Config("DB_NAME"),
	}

	openConnection := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Jakarta",
		configDB["DB_Host"],
		configDB["DB_Username"],
		configDB["DB_Password"],
		configDB["DB_Name"])

	// Open the connection
	db, err := gorm.Open(postgres.Open(openConnection), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	log.Println("DATABASE Successfully connected!")
	// return the connection
	return db
}
