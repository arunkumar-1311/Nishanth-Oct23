package adaptor

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"job-post/repository"
	"os"
)

// Helps to occur the new connection
func NewConnection() *gorm.DB {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	dbConnection := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("user"), os.Getenv("password"), os.Getenv("host"), os.Getenv("port"), os.Getenv("dbname"))
	db, err := gorm.Open(postgres.Open(dbConnection), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(err)
	}
	return db
}

// Helps to acquire the connection
func AcquireConnection(db *gorm.DB, redis *redis.Client) Database {
	return &repository.DB_Connection{DB: db, Redis: redis}
}
