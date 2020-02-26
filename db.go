package chassis

import (
	"sync"
	"time"

	"github.com/jinzhu/gorm"

	"pgxs.io/chassis/config"
	xLog "pgxs.io/chassis/log"
)

var (
	db     *gorm.DB
	dbOnce sync.Once
)

func mustConnectDB(dbCfg config.DatabaseConfig) *gorm.DB {
	log := xLog.New().Service("chassis").Category("gorm")
	db, err := gorm.Open("mysql", dbCfg.DSN)
	if err != nil {
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
	return db
}

//DB get *Db
func DB() *gorm.DB {
	dbOnce.Do(func() {
		db = mustConnectDB(config.Database())
	})
	return db
}

//Close close db
func CloseDB() error {
	return db.Close()
}

type MultiDBSource struct {
	lock sync.RWMutex
	dbs  map[string]*gorm.DB
}

type DBSource struct {
	cfg       *config.DatabaseConfig
	dbs       *gorm.DB
	connected bool
}

var (
	dbs         *MultiDBSource
	dbsInitOnce sync.Once
)

//DBs get all custom database
func DBs() *MultiDBSource {
	dbsInitOnce.Do(func() {
		dbs = new(MultiDBSource)
		dbs.dbs = make(map[string]*gorm.DB)
	})
	return dbs
}

//Set add a new or reset  database source for app
func (s MultiDBSource) Set(name string, cfg config.DatabaseConfig) {
	s.lock.Lock()
	defer s.lock.Unlock()
	db := mustConnectDB(cfg)
	s.dbs[name] = db
}

//Get get db by name
func (s MultiDBSource) Get(name string) (db *gorm.DB, ok bool) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	db = s.dbs[name]
	if db != nil {
		ok = true
	}
	return
	ok = false
	return
}

//Size get db size
func (s MultiDBSource) Size() int {
	s.lock.RUnlock()
	defer s.lock.RUnlock()
	return len(s.dbs)
}
