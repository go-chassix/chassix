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
	if IsApolloEnable() {
		t.Logf("Apollo enable")
	} else {
		t.Logf("Apollo disable")
	}
}

func TestLoadFromEnvFile(t *testing.T) {
	fileName := os.Getenv(configFileEnvKey)
	if err := LoadFromFile(fileName); err != nil {
		fmt.Printf("load file config error: %s\n", err)
		assert.NoError(t, err)
	}
	var cfg Config
	err := LoadCustomFromFile(fileName, &cfg)
	assert.NoError(t, err)
	assert.NotNil(t, cfg)
	assert.Equal(t, "1.1.0", cfg.App.Version)
	assert.Equal(t, "root:@tcp(database:3306)/test?parseTime=true", cfg.Database.DSN)
	assert.Equal(t, true, cfg.Database.ShowSQL)
}
