package chassis

import (
	"time"

	"github.com/jinzhu/gorm"
	log "gopkg.in/logger.v1"

	"pgxs.io/panguxs/pkg/chassis/config"
)

var db *gorm.DB

func connectDB() {
	dbCfg := config.Database()
	var err error
	db, err = gorm.Open("mysql", dbCfg.DSN)
	if err != nil {
		//todo
		log.Fatalln(err)
	}
	db.LogMode(true)

	if dbCfg.MaxIdle > 0 {
		db.DB().SetMaxIdleConns(dbCfg.MaxIdle)
	}
	if dbCfg.MaxOpen > 0 && dbCfg.MaxOpen > dbCfg.MaxIdle {
		db.DB().SetMaxOpenConns(100)
	}
	if dbCfg.MaxLifetime > 0 {
		db.DB().SetConnMaxLifetime(time.Duration(dbCfg.MaxLifetime) * time.Second)
	}
}

//DB get *Db
func DB() *gorm.DB {
	return db
}

//Close close db
func Close() error {
	return db.Close()
}
