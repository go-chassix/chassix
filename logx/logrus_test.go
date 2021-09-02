package logx

import (
	"github.com/go-chassix/chassix/v2/config"
	"testing"
)

func Test_Logger(t *testing.T) {
	config.LoadFromEnvFile()

	New().Component("log").Category("test").Info("test log")
}
