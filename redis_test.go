package chassix

import (
	"context"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"

	"github.com/go-chassix/chassix/v2/config"
)

func TestReadOptions(t *testing.T) {
	config.LoadFromEnvFile()

	var opts redis.Options
	readRedisOptions(&opts)
	assert.NotEmpty(t, opts)
	assert.NotNil(t, &opts)
	assert.Equal(t, "redis:6379", opts.Addr)
}
func TestRedisClient(t *testing.T) {
	config.LoadFromEnvFile()
	ctx := context.Background()
	RedisClient().Set(ctx, "test", "123", 5*time.Minute)

	val := RedisClient().Get(ctx, "test").Val()
	assert.NotEmpty(t, val)
	assert.Equal(t, "123", val)
}
