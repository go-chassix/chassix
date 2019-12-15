package chassis

import (
	"errors"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql" //import mysql driver
	bindata "github.com/golang-migrate/migrate/v4/source/go_bindata"
	xLog "pgxs.io/chassis/log"
)


//Run new bindataInstance and UP
func Run(assetNames []string, afn bindata.AssetFunc, dbURL string) error {
	log := xLog.New().Service("chassis").Category("migrate")

	// wrap assets in Resource
	s := bindata.Resource(assetNames, afn)

	d, err := bindata.WithInstance(s)
	if err != nil {
		log.Error(err)
		return errors.New("DB migrations build instance error")
	}

	m, err := migrate.NewWithSourceInstance("go-bindata", d, "mysql://"+dbURL)
	if err != nil {
		log.Error(err)
		return errors.New("DB migrations build bindata instance error")
	}
	upErr := m.Up() // run your migrations
	if upErr != nil && upErr != migrate.ErrNoChange {
		log.Errorf("数据库迁移异常.错误:,%s", upErr.Error())
		return errors.New("DB migrations UP error " + upErr.Error())
	}
	return nil
}
