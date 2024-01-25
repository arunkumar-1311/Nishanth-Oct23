package adaptor

import (
	"blog_post/logger"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	log "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

// Get a new connection
func dbConn() {
	err := godotenv.Load(".env")
	if err != nil {
		logger.Logging().Error(err)
		fmt.Print("\nenv error")
		panic(err)
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

}

// Check if the connection is exist of helps to occur the connection
func GetConn() *gorm.DB {
	if db == nil {
		dbConn()
	}
	return db
}
