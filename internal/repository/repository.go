package repository

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase(g *gorm.DB) error {
	dsn := "host=nattily-apt-kingfisher.data-1.use1.tembo.io user=postgres password=R3uJQJheg2H7BsTk dbname=postgres port=5432 sslmode=require TimeZone=Asia/Shanghai"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	fmt.Printf("Successfully connected to %v", dsn)
	return nil
}
