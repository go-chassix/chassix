package logx

import (
	"c6x.io/chassis/config"
	"testing"
)

func Test_Logger(t *testing.T) {
	config.LoadFromEnvFile()

	New().Component("log").Category("test").Info("test log")
}
