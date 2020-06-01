package chassis

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"pgxs.io/chassis/config"
	"testing"
)

func TestMDB(t *testing.T)  {
	defer CloseMDB()
	// given
	fileName := os.Getenv("PG_CONF_FILE")
	if err := config.LoadFromFile(fileName); err != nil {
		fmt.Printf("load file config error: %s\n", err)
		assert.NoError(t, err)
	}
	// when
	db1 := MDB("db1")
	assert.NotNil(t, db1)
	assert.Nil(t, db1.DB().Ping())
}