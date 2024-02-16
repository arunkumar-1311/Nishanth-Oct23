package main

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"to-do/adaptor"
	"to-do/logger"
	"to-do/lookup"
	"to-do/router"
	"to-do/service"

	"gorm.io/gorm"
)

func main() {
	// helps to get new connection
	db := adaptor.NewConnection()

	redis, err := adaptor.RedisConnection()
	if err != nil {
		logger.ZeroLogger().Msg(err.Error())
		return
	}

	// Helps to load the lookup
	file, err := os.ReadDir("./lookup")
	if err != nil {
		logger.ZeroLogger().Msg(err.Error())
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
					logger.ZeroLogger().Msg(err.Error())
				}

				lookupVersion := lookup.Lookup{
					Name:    fileName,
					Version: nameAndVersion[1],
				}

				if err := db.Create(lookupVersion).Error; err != nil {
					logger.ZeroLogger().Msg(err.Error())
					return
				}
			}
		}
	}

	if err := router.Router(adaptor.AcquireConnection(db, redis), service.AcquireService()); err != nil {
		fmt.Print("Can't connect to the server ", err)
		return
	}
}
