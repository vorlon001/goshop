package dbs

import (
	"github.com/quangdangfit/gocommon/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"fmt"
	"goshop/internal/app/models"
	"goshop/internal/config"
)

var Database *gorm.DB

func Init() {

	cfg := config.GetConfig()
	fmt.Printf("DNS:Init:cfg:%#v\n",cfg)
	database, err := gorm.Open(postgres.Open(cfg.DatabaseURI), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Warn),
	})

	if err != nil {
		logger.Fatal("Cannot connect to database", err)
	}

	// Set up connection pool
	sqlDB, err := database.DB()
	if err != nil {
		logger.Fatal("Cannot connect to database", err)
	}
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(200)
	Database = database
        fmt.Printf("DNS:Init:Database:%#v\n",Database)

	Migrate()
}

func Migrate() {
	User := models.User{}
	Product := models.Product{}
	Order := models.Order{}
	OrderLine := models.OrderLine{}

	Database.AutoMigrate(&Product, &User, Order, OrderLine)
}
