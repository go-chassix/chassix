package config

import (
	"gopkg.in/apollo.v0"
)

var (
	configFileEnvKey = "PG_CONF_FILE"
	config           *Config
)

func init() {
	config = new(Config)
}

//Config all config
type Config struct {
	App       AppConfig
	Databases []*DatabaseConfig `yaml:"databases"`
	OpenAPI   OpenAPIConfig     `yaml:"openapi"`
	Server    ServerConfig
	Logging   LoggingConfig
	Mails     []MailConfig `yaml:"mail,flow"`
	Apollo    ApolloConfig
}

//SetLoadFileEnvKey set env var name for read the config file path, default:PG_CONF_FILE
func SetLoadFileEnvKey(key string) {
	configFileEnvKey = key
}

//LoadFileEnvKey load file env key
func LoadFileEnvKey() string {
	return configFileEnvKey
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
	MaxIdle     int
	MaxOpen     int
	MaxLifetime int
	ShowSQL     bool `yaml:"showSQL"`
}

//OpenAPIConfig open api config
type OpenAPIConfig struct {
	Spec struct {
		Title       string
		Description string `yaml:"desc"`
		Contact     struct {
			Name  string
			Email string
			URL   string
		} `yaml:"contact"`
		License struct {
			Name string
			URL  string
		} `yaml:"license"`
		Version string
	}
	Tags []OpenapiTagConfig `yaml:",flow"`
	UI   OpenapiUIConfig    `yaml:"ui"`
}

//OpenapiUIConfig swagger ui config
type OpenapiUIConfig struct {
	API        string `yaml:"api"`
	Dist       string
	Entrypoint string `yaml:"entrypoint"`
}

//OpenapiTagConfig openapi tag
type OpenapiTagConfig struct {
	Name        string
	Description string `yaml:"desc"`
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
func Openapi() OpenAPIConfig {
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
func IsNil() bool {
	return config == nil
}
func NotNil() bool {
	return config != nil
}
