package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/kawamurasorachi/Gin-Postgres/entity"
)

var (
	db  *gorm.DB
	err error
)

func Init() {
	con := GetEnv()
	db, err = gorm.Open("postgres", con)
	if err != nil {
		panic(err)
	}
	autoMigration()
}

func GetDB() *gorm.DB {
	return db
}

func GetEnv() string {
	err := godotenv.Load(fmt.Sprintf("%spostgres.env", os.Getenv("GO_ENV")))
	if err != nil {
		panic(err)
	}
	cons := map[string]string{
		"host":     "0.0.0.0",
		"port":     "5432",
		"sslmode":  "disable",
		"user":     os.Getenv("POSTGRES_USER"),
		"password": os.Getenv("POSTGRES_PASSWORD"),
		"dbname":   os.Getenv("POSTGRES_DB"),
	}
	var con string
	for k, v := range cons {
		con += fmt.Sprintf("%s=%s ", k, v)
	}
	return con
}

func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}

func autoMigration() {
	db.AutoMigrate(&entity.User{})
}
