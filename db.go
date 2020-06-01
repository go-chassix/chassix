package chassis

import (
	"fmt"
	"sync"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"pgxs.io/chassis/config"
	xLog "pgxs.io/chassis/log"

)

var (
	db     *gorm.DB
	dbOnce sync.Once
)

func mustConnectDB(dbCfg config.DatabaseConfig) *gorm.DB {
	log := xLog.New().Service("chassis").Category("gorm")
	dialect := dbCfg.Dialect
	if "" == dialect {
		dialect = "mysql"
	}
	db, err := gorm.Open(dialect, dbCfg.DSN)
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
	mdbs        *MultiDBSource
	dbsInitOnce sync.Once
)

//DBs get all custom database
func DBs() *MultiDBSource {
	dbsInitOnce.Do(func() {
		mdbs = new(MultiDBSource)
		mdbs.dbs = make(map[string]*gorm.DB)
	})
	return mdbs
}

func MDB(name string) *gorm.DB {
	multiCfg := config.MultiDatabase()
	if _, ok := multiCfg[name]; !ok {
		panic(fmt.Sprintf("Cant't find database configuration for db %s", name))
	}
	dbsInitOnce.Do(func() {
		mdbs = new(MultiDBSource)
		mdbs.dbs = make(map[string]*gorm.DB, len(multiCfg))
		mdbs.lock.Lock()
		defer mdbs.lock.Unlock()
		for k, v := range multiCfg {
			mdbs.dbs[k] = mustConnectDB(v)
		}
	})
	return mdbs.dbs[name]
}

//Close close multi database
func CloseMDB() error {
	for _, v := range mdbs.dbs {
		if err := v.Close(); nil != err {
			return err
		}
	}
	return nil
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
