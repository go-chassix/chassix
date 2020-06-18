package chassis

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"pgxs.io/chassis/config"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func TestDBs(t *testing.T) {
	//defer CloseAllDB()
	// given
	fileName := os.Getenv("PG_CONF_FILE")
	if "" == fileName {
		fileName = "configs/app.yml"
	}
	if err := config.LoadFromFile(fileName); err != nil {
		fmt.Printf("load file config error: %s\n", err)
		assert.NoError(t, err)
	}
	// when
	dbs, _ := DBs()
	// then
	assert.NotNil(t, dbs[1])
	assert.Nil(t, dbs[1].DB().Ping())
}
