package migrates

import (
	"crud_demo/app/models"
	"crud_demo/config"
)

func MigrateTableUser() {
	if !config.DB.HasTable(&models.User{}) {
		if err := config.DB.CreateTable(&models.User{}).Error; err != nil {
			panic(err)
		}
	}
}


