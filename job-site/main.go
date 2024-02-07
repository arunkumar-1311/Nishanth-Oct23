package main

import (
	"fmt"
	"job-post/adaptor"
	"job-post/logger"
	"job-post/lookup"
	"job-post/router"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/go-kit/log/level"
	"gorm.io/gorm"
)

func main() {
	// helps to get new connection
	db := adaptor.NewConnection()

	// Helps to load the lookup
	file, err := os.ReadDir("./lookup")
	if err != nil {
		level.Debug(logger.GokitLogger(err)).Log()
		return
	}

	for _, value := range file {
		if value.Name() != "master.go" {
			fileName := strings.TrimSuffix(value.Name(), filepath.Ext(value.Name()))
			nameAndVersion := strings.Split(fileName, "_")
			emp := lookup.Empty{}
			if lookUp := db.Model(&lookup.Lookup{}).Where("version = ?", nameAndVersion[1]).Find(&emp); lookUp.RowsAffected == 0 {

				method := reflect.ValueOf(&emp).MethodByName(fileName).Interface().(func(*gorm.DB) error)
				if err := method(db); err != nil {
					level.Debug(logger.GokitLogger(err)).Log()
				}

				lookupVersion := lookup.Lookup{
					Name:    fileName,
					Version: nameAndVersion[1],
				}

				if err := db.Create(lookupVersion).Error; err != nil {
					level.Debug(logger.GokitLogger(err)).Log()
					return
				}
			}
		}
	}

	if err := router.Router(adaptor.AcquireConnection(db)); err != nil {
		fmt.Print("Can't connect to the server ", err)
		return
	}
	// level.Debug(logger.GokitLogger(fmt.Errorf("error"))).Log()
}
