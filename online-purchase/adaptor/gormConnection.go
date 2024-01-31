package adaptor

import (
	"fmt"
	"online-purchase/repository"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// Contains all needed methods to manipulate db operations
type Database interface {
}

// Helps to get new DB connection
func NewDB_Connection() (*gorm.DB, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	dbConn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("user"), os.Getenv("password"), os.Getenv("host"), os.Getenv("port"), os.Getenv("dbname"))
	DB, err := gorm.Open(postgres.Open(dbConn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}
	return DB, nil
}

// Set the db connection to the interface
func AcquireConnection(db *gorm.DB) Database {
	return &repository.GORM_Connection{DB: db}

}
