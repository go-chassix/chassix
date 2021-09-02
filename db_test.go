package chassix

import (
	"testing"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/stretchr/testify/assert"

	"github.com/go-chassix/chassix/v2/config"
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
