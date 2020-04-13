package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"os"
)

var db *gorm.DB //база данных

func init()  {
	env := godotenv.Load(".envir")
	if env != nil {
		fmt.Print(env)
	}

	userName := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")

	dbUri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, userName, dbName, password)
	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Println(err)
	}
	db = conn
	db.Debug().AutoMigrate(&Account{}, &Contact{})
}

func GetDB() *gorm.DB {
	return db
}