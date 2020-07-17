package chassis

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"pgxs.io/chassis/config"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func TestDBs(t *testing.T) {
	//defer CloseAllDB()
	// given
	config.LoadFromEnvFile()
	// when
	dbs, _ := DBs()
	// then
	assert.NotNil(t, dbs[1])
	assert.Nil(t, dbs[1].DB().Ping())
}
