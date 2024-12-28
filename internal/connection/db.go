package connection

import (
	"log"

	"github.com/midoo/go_ecommerce/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDbConnection(cnf configs.AppConfig) *gorm.DB {
	dialect := mysql.Open(cnf.Dsn)
	db, err := gorm.Open(dialect, &gorm.Config{})
	if err != nil {
		log.Fatalf("failed create connection to DB : %s", err.Error())
	}

	return db
}
