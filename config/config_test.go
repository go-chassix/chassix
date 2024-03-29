package config

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}
func setup() {
	LoadFromEnvFile()
}

func TestConfig(t *testing.T) {
	mails := Mails()
	assert.NotNil(t, mails)
	fmt.Println((mails)[0])
	assert.Len(t, mails, 1, "已设置1个邮箱配置")
	assert.Equal(t, "imap.example.com:993", mails[0].IMAPAddr, "测试邮箱IMAP地址应为imap.example.com:993")
	assert.Equal(t, "test", Openapi().Spec.License.Name)
	assert.Equal(t, "test", Openapi().Spec.License.URL)
	assert.True(t, Logging().NoColors)
	assert.Equal(t, 10, Database().MaxIdle)
	assert.Equal(t, 50, Database().MaxOpen)
	assert.Equal(t, 50, Database().MaxLifetime)

	if IsApolloEnable() {
		t.Logf("Apollo enable")
	} else {
		t.Logf("Apollo disable")
	}
}

func TestLoadFromEnvFile(t *testing.T) {
	fileName := os.Getenv(appConfigFileEnvKey)
	if err := LoadFromFile(fileName); err != nil {
		fmt.Printf("load file config error: %s\n", err)
		assert.NoError(t, err)
	}
	var cfg Config
	err := LoadCustomFromFile(fileName, &cfg)
	assert.NoError(t, err)
	assert.NotNil(t, cfg)
	assert.Equal(t, "1.1.0", cfg.App.Version)
	assert.Equal(t, "root:@tcp(database:3306)/test?parseTime=true", cfg.Databases[0].DSN)
	assert.Equal(t, ":memory:", cfg.Databases[1].DSN)
	assert.Equal(t, "postgres://postgres:123456@postgres:5432/test?sslmode=disable", cfg.Databases[2].DSN)
	assert.Equal(t, true, cfg.Databases[0].ShowSQL)
}

func TestRedis(t *testing.T) {
	redisCfg := Redis()
	assert.NotNil(t, Redis())
	assert.NotEmpty(t, Redis())
	assert.Equal(t, "simple", redisCfg.Mode)
	assert.Equal(t, "", redisCfg.Username)
	assert.Equal(t, "", redisCfg.Password)
	assert.Equal(t, 10, redisCfg.PoolSize)
	assert.Equal(t, 5, redisCfg.MinIdleConns)
	assert.Equal(t, "", redisCfg.MaxConnAge)
	assert.Equal(t, "4s", redisCfg.PoolTimeout)
	assert.Equal(t, "5m", redisCfg.IdleTimeout)
	assert.Equal(t, "1m", redisCfg.IdleCheckFrequency)
	assert.NotNil(t, redisCfg.Sentinel)
}
