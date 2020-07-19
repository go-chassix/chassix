package chassis

import (
	"errors"
	"io/ioutil"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql" //import mysql driver
	"github.com/golang-migrate/migrate/v4/source/go_bindata"

	"c6x.io/chassis/logx"
)

//Migrate Run new bindataInstance and UP
func Migrate(assetNames []string, afn bindata.AssetFunc, dbURL, dialect string) error {
	log := logx.New().Service("chassis").Category("migrate")
	// wrap assets in Resource
	s := bindata.Resource(assetNames, afn)

	d, err := bindata.WithInstance(s)
	if err != nil {
		log.Error(err)
		return errors.New("DB migrations build instance error")
	}

	databaseURL := ""
	if dialect == "" || dialect == "mysql" {
		databaseURL = "mysql://" + dbURL
	}
	if dialect == "sqlite3" {
		databaseURL = "sqlite3://" + dbURL
	}
	if dialect == "postgres" {
		databaseURL = "postgres://" + dbURL
	}
	m, err := migrate.NewWithSourceInstance("go-bindata", d, databaseURL)
	if err != nil {
		log.Error(err)
		return errors.New("DB migrations build bindata instance error")
	}

	//IF ENV NOT PROD IMPORT TEST DATA
	if !EnvIsProd() {
		if err := m.Down(); err != nil {
			log.Error("down: ", err)
		}
	}

	upErr := m.Up() // run migrations
	if upErr != nil && upErr != migrate.ErrNoChange {
		log.Errorf("Run DB migrations failed,error:%s", upErr.Error())
		return errors.New("DB migrations UP error " + upErr.Error())
	}

	//IF ENV NOT PROD IMPORT TEST DATA
	if !EnvIsProd() {
		fileName := os.Getenv(EnvPgTestDataFile)
		log.Debugf("import data file: %s", fileName)
		if fileName != "" {
			if file, err := os.Open(fileName); err == nil {
				// count := 0
				if data, err := ioutil.ReadAll(file); err == nil {
					if db, err := DB(); nil != err {
						return err
					} else {
						db.Exec(string(data))
					}
				}
			} else {
				log.Error(err)
			}
		} else {
			log.Error("import test data failed")
		}
	}
	return nil
}

//ExecSQLFile danger!!! just for importing test data.
func ExecSQLFile(fileName string) error {
	log := logx.New().Service("chassis").Category("execsql")
	if fileName != "" {
		if file, err := os.Open(fileName); err == nil {
			// count := 0
			if data, err := ioutil.ReadAll(file); err == nil {
				if db, err := DB(); nil != err {
					return err
				} else {
					err = db.Exec(string(data)).Error
					if err != nil {
						return err
					}
				}
			}
		} else {
			log.Error(err)
			return err
		}
	} else {
		log.Error("import test data failed")
		return errors.New("import test data failed, file cannot be empty")
	}
	return nil
}
