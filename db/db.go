package db

import (
	"fmt"
	"lfs-portal/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func NewDb() *gorm.DB {

	var err error
	db, err = gorm.Open("mysql", "root:@(localhost)/labonte?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}

	//Migrate the schema
	migrateModels()
	return db
}

func migrateModels() {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Customer{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	db.AutoMigrate(&models.Job{}).AddForeignKey("customer_id", "customers(id)", "RESTRICT", "RESTRICT")
}
