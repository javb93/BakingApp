package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func ConnectDB() (*gorm.DB, error) {
	dsn := "user=" + os.Getenv("DBUSERLOCAL") + " dbname=bakingapp sslmode=disable password=" + os.Getenv("DBPASSWORDLOCAL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil

}
