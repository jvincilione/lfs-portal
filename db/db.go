package db

import (
	"fmt"
	"lfs-portal/models"
	"lfs-portal/utils"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func NewDb() *gorm.DB {
	var err error
	cfg := utils.Config
	connectString := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", cfg.DB_USER, cfg.DB_PASS, cfg.DB_HOST, cfg.DB_PORT, cfg.DB_NAME)
	db, err = gorm.Open("mysql", connectString)

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
