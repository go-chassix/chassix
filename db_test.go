package chassis

import (
	"testing"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/stretchr/testify/assert"

	"c6x.io/chassis/config"
)

func TestDBs(t *testing.T) {
	//defer CloseAllDB()
	// given
	config.LoadFromEnvFile()
	dbCfg := config.Databases()
	assert.NotEmpty(t, dbCfg)
	// when
	dbs, _ := DBs()
	// then
	assert.NotNil(t, dbs[1])
	assert.Nil(t, dbs[1].DB().Ping())
}
