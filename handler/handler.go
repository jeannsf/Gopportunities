package handler

import (
	"github.com/jeannsf/Gopportunities/config"
	"gorm.io/gorm"
)


var (
	logger *config.Logger
	db *gorm.DB
)

func InitializeHandler(){
	logger = config.GetLogger("Handler")
	db = config.GetSQLite()
}



