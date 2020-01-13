package chassis

import (
	"time"

	"github.com/jinzhu/gorm"

	"pgxs.io/chassis/config"
	xLog "pgxs.io/chassis/log"
)

var db *gorm.DB

func connectDB() {
	log := xLog.New().Service("chassis").Category("gorm")
	dbCfg := config.Database()
	var err error
	db, err = gorm.Open("mysql", dbCfg.DSN)
	if err != nil {
		//todo
		log.Fatalln(err)
	}
	db.LogMode(dbCfg.ShowSQL)

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
