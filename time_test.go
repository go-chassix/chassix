package chassis

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type User struct {
	Date JSONDate `json:"date,omitempty"`
}

func Test_JSONDate(t *testing.T) {

	var date = `
	{
		"date": "2020-10-10"
	}
	`
	var u User
	json.Unmarshal([]byte(date), &u)
	t2 := time.Time(u.Date).Local().Format(dateFormat)
	assert.Equal(t, "2020-10-10", t2)
}
