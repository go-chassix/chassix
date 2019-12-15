package config

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Config(t *testing.T) {
	mails := Mails()
	assert.NotNil(t, mails)
	fmt.Println((mails)[0])
	assert.Len(t, mails, 1, "已设置1个邮箱配置")
	assert.Equal(t, "imap.example.com:993", mails[0].IMAPAddr, "测试邮箱IMAP地址应为imap.example.com:993")

}
