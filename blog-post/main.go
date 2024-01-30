package main

import (
	"blog_post/adaptor"
	"blog_post/logger"
	"blog_post/lookup"
	routing "blog_post/router"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func main() {
	// Acquiring Database Connection
	db := adaptor.NewDB_Connection()

	// Helps to load the lookup
	file, err := os.ReadDir("./lookup")
	if err != nil {
		logger.Logging().Print(err)
		return
	}

	for _, value := range file {
		if value.Name() != "master.go" {
			fileName := strings.TrimSuffix(value.Name(), filepath.Ext(value.Name()))
			nameAndVersion := strings.Split(fileName, "_")
			emp := lookup.Empty{}
			if lookUp := db.Model(&lookup.Lookup{}).Where("version = ?", nameAndVersion[1]).Find(&emp); lookUp.RowsAffected == 0 {

				method := reflect.ValueOf(&emp).MethodByName(fileName).Interface().(func(*gorm.DB))
				method(db)

				lookupVersion := lookup.Lookup{
					Name:    fileName,
					Version: nameAndVersion[1],
				}

				if err := db.Create(lookupVersion); err.Error != nil {
					logger.Logging().Print(err)
					return
				}
			}
		}
	}

	// Helps to initiate the fiber configration and router
	router := fiber.New(fiber.Config{
		CaseSensitive: true,
		AppName:       "Blog Post v1.0.1",
	})

	routes := routing.Routes(router, db)
	if err := routes.Listen(":8000"); err != nil {
		fmt.Println("Can't listen to the server ", err)
		logger.Logging().Print(err)
		return
	}
}
