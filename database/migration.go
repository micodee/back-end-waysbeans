package database

import (
	"waysbeans/models"
	"waysbeans/pkg/mysql"
	"fmt"
)

func RunMigration() {
	err := mysql.ConnDB.AutoMigrate(&models.Product{})

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}