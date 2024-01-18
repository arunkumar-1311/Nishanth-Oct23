package adaptor

import (
	"blog_post/logger"
	"blog_post/models"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	log "gorm.io/gorm/logger"
)

// Get a new connection
func dbConn() (db *gorm.DB) {
	err := godotenv.Load(".env")
	if err != nil {
		logger.Logging().Error(err)
		fmt.Print("\nenv error")

	}

	dbConn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("user"), os.Getenv("password"), os.Getenv("host"), os.Getenv("port"), os.Getenv("dbname"))
	db, err = gorm.Open(postgres.Open(dbConn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: log.Default.LogMode(log.Silent),
	})
	if err != nil {
		logger.Logging().Error(err)
		panic(err)
	}
	db.AutoMigrate(&models.Post{}, &models.Category{}, &models.Roles{}, &models.Users{}, &models.Comments{})

	return
}

// Check if the connection is exist of helps to occur the connection
func GetConn() (db *gorm.DB) {
	if db == nil {
		db = dbConn()
	}
	return db
}
