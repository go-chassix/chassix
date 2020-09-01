package config

import (
	"gopkg.in/apollo.v0"
)

var (
	appConfigFileEnvKey = "CHASSIX_APP_CONF"
	apiConfigFileEnvKey = "CHASSIX_API_CONF"
	config              *Config
)

func init() {
	config = new(Config)
}

//Config all config
type Config struct {
	App       AppConfig
	Databases []*DatabaseConfig `yaml:"databases,flow"`
	OpenAPI   *OpenAPIConfig    `yaml:"openapi"`
	Server    ServerConfig      `yaml:"server"`
	Logging   LoggingConfig     `yaml:"logging"`
	Mails     []MailConfig      `yaml:"mail,flow"`
	Apollo    ApolloConfig      `yaml:"apollo"`
	Redis     RedisConfig       `yaml:"redis"`
}

//SetLoadFileEnvKey set env var name for read the config file path, default:PG_CONF_FILE
func SetLoadFileEnvKey(key string) {
	appConfigFileEnvKey = key
}

//LoadFileEnvKey load file env key
func LoadFileEnvKey() string {
	return appConfigFileEnvKey
}

//LoggingConfig log config
type LoggingConfig struct {
	Level        uint32
	ReportCaller bool `yaml:"report-caller"`
	NoColors     bool `yaml:"no-colors"`
	CallerFirst  bool `yaml:"caller-first"`
}

//AppConfig application config
type AppConfig struct {
	Name    string
	Version string
	Env     string
}

//ServerConfig server config
type ServerConfig struct {
	Port int
}

//DatabaseConfig db config
type DatabaseConfig struct {
	Dialect     string `yaml:"dialect"`
	DSN         string `yaml:"dsn"`
	MaxIdle     int    `yaml:"maxIdle"`
	MaxOpen     int    `yaml:"maxOpen"`
	MaxLifetime int    `yaml:"maxLifetime"`
	ShowSQL     bool   `yaml:"showSQL"`
}

//MailConfig mail config
type MailConfig struct {
	IMAPAddr string `yaml:"imap-addr"`
	SMTPAddr string `yaml:"smtp-addr"`
	TLS      bool   `yaml:"tls"`
	Username string
	Password string
}

// IsApolloEnable is apollo enable
func IsApolloEnable() bool {
	return config.Apollo.Enable
}

// ApolloConfig apollo config
type ApolloConfig struct {
	Enable bool        `yaml:"enable"`
	Conf   apollo.Conf `yaml:"conf"`
}

//App app config
func App() AppConfig {
	return config.App
}

//Server server config
func Server() ServerConfig {
	return config.Server
}

//OpenAPI openapi config
func Openapi() *OpenAPIConfig {
	return config.OpenAPI
}

//Databases Multi Database config
func Databases() []*DatabaseConfig {
	return config.Databases
}

//Database first database config
func Database() *DatabaseConfig {
	if len(config.Databases) > 0 {
		return config.Databases[0]
	}
	return nil
}

//Logging log config
func Logging() LoggingConfig {
	return config.Logging
}

//Mails mail settings
func Mails() []MailConfig {
	return config.Mails
}

//IsNil check config init status
func IsNil() bool {
	return config == nil
}

//NotNil check config not nil
func NotNil() bool {
	return config != nil
}

//Redis return redis config
func Redis() RedisConfig {
	return config.Redis
}
