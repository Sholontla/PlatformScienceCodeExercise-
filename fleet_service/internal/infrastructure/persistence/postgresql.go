package persistence

import (
	"fmt"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

const (
	host = "127.0.0.1"
	port = 5435
)

var err error

func Connect() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, "test", "test", "test", port)
	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully connected to the database")
	}
}

var DbPool = &sync.Pool{
	New: func() interface{} {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, "test", "test", "test", port)
		Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err != nil {
			panic(err)
		}
		return Database
	},
}

func Conn() *gorm.DB {
	client := DbPool.Get().(*gorm.DB)
	return client
}

func ReleaseDB(client *gorm.DB) {
	DbPool.Put(client)
}
