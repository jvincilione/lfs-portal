package db

import (
	"fmt"
	"lfs-portal/models"
	"lfs-portal/utils"

	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
)

var db *gorm.DB

func NewDb() *gorm.DB {
	var err error
	cfg := utils.Config
	connectString := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", cfg.DB_USER, cfg.DB_PASS, cfg.DB_HOST, cfg.DB_PORT, cfg.DB_NAME)
	if cfg.GIN_MODE == "release" {
		connectString = fmt.Sprintf("%s:%s@unix(/cloudsql/%s)/%s?parseTime=True", cfg.DB_USER, cfg.DB_PASS, cfg.DB_HOST, cfg.DB_NAME)
	}
	db, err = gorm.Open("mysql", connectString)

	if err != nil {
		logrus.Info(connectString)
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}

	//Migrate the schema
	migrateModels()
	logrus.Info("Successfully connected and migrated database.")
	return db
}

func migrateModels() {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Company{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	db.AutoMigrate(&models.Job{}).AddForeignKey("company_id", "companies(id)", "RESTRICT", "RESTRICT")
}
