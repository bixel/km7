package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Task struct {
	gorm.Model
	Name        string
	SeasonStart *time.Time
	SeasonEnd   *time.Time
	Frequency   *time.Duration
}

func main() {
	dbpath := os.Getenv("KM7_DBPATH")
	if dbpath == "" {
		dbpath = "database.sqlite"
	}

	db, err := gorm.Open("sqlite3", dbpath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.AutoMigrate(&Task{})

	tasks := make([]Task, 0)
	db.Find(&tasks)

	for _, task := range tasks {
		fmt.Println(task)
	}
}
