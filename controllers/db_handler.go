package controllers

import (
	"database/sql"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func connect() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/db_latihan_pbp?parseTime=true&loc=Asia%2FJakarta")

	if err != nil {
		log.Fatal(err)
		return nil
	}

	return db
}

func connectGORM() *gorm.DB{
	dsn := "root:@tcp(localhost:3306)/db_latihan_pbp?parseTime=true&loc=Asia%2FJakarta"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
		return nil
	}

	return db
}