package db

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

type DbConfig struct {
	username string
	password string
	host     string
	port     string
	name     string
}

func InitDb() *gorm.DB {
	Db = connectDB()
	return Db
}

func connectDB() *gorm.DB {
	con := *loadConfig()

	dsn := con.username + ":" + con.password + "@tcp" + "(" + con.host + ":" + con.port + ")/" + con.name + "?" + "parseTime=true&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Error db connect: ", err)
		return nil
	}

	return db
}

func loadConfig() *DbConfig {
	return &DbConfig{
		username: os.Getenv("DB_USERNAME"),
		password: os.Getenv("DB_PASSWORD"),
		host:     os.Getenv("DB_HOST"),
		port:     os.Getenv("DB_PORT"),
		name:     os.Getenv("DB_NAME"),
	}
}
