package config

const (
	configFileEnvKey = "PG_CONF_FILE"
)

func init() {
	//todo check config
	loadFromEnv()
}

func initLog() {

}

//Config all config
type Config struct {
	App      AppConfig
	Database DatabaseConfig
	Openapi  OpenapiConfig
	Server   ServerConfig
	Logging  LoggingConfig
	Mails    []*MailConfig `yaml:"mail,flow"`
}

//LoggingConfig log config
type LoggingConfig struct {
	Level        uint32
	ReportCaller bool `yaml:"report-caller"`
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
	DSN         string `yaml:"dsn"`
	MaxIdle     int
	MaxOpen     int
	MaxLifetime int
}

//OpenapiConfig open api config
type OpenapiConfig struct {
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
	Entrypoint string
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

var config Config

//App app config
func App() *AppConfig {
	return &config.App
}

//Server server config
func Server() *ServerConfig {
	return &config.Server
}

//Openapi openapi config
func Openapi() *OpenapiConfig {
	return &config.Openapi
}

//Database DB config
func Database() *DatabaseConfig {
	return &config.Database
}

//Logging log config
func Logging() *LoggingConfig {
	return &config.Logging
}

//Mails mail settings
func Mails() []*MailConfig {
	return config.Mails
}
