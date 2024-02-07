package main

import (
	"fmt"
	"online-purchase/adaptor"
	"online-purchase/logger"
	"online-purchase/lookup"
	"online-purchase/router"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"gorm.io/gorm"
)

func main() {
	// Getting DB connection
	db, err := adaptor.NewDB_Connection()
	if err != nil {
		logger.ZapLog().Error(fmt.Sprint(err))
	}

	// Helps to load the lookup
	file, err := os.ReadDir("./lookup")
	if err != nil {
		logger.ZapLog().Error(fmt.Sprint(err))
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
					logger.ZapLog().Error(fmt.Sprint(err))
				}

				lookupVersion := lookup.Lookup{
					Name:    fileName,
					Version: nameAndVersion[1],
				}

				if err := db.Create(lookupVersion); err.Error != nil {
					logger.ZapLog().Error(fmt.Sprint(err))
					return
				}
			}
		}
	}

	router.Routes(adaptor.AcquireConnection(db))

}
