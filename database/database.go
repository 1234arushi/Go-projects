package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)
var(
	DBConn *gorm.DB//go lang orm to help connect with database


)