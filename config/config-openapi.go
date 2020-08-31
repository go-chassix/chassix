package config

//OpenAPIConfig open api config
type OpenAPIConfig struct {
	Enabled bool `yaml:"enabled"`
	Spec    struct {
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
	Tags      []OpenapiTagConfig     `yaml:",flow"`
	UI        OpenapiUIConfig        `yaml:"ui"`
	GlobalApi OpenapiGlobalApiConfig `yaml:"global-api"`
	APIs      OpenapiServiceConfig   `yaml:"apis"`
}

//OpenapiUIConfig swagger ui config
type OpenapiUIConfig struct {
	API        string `yaml:"api"`
	Dist       string `yaml:"dist"`
	Entrypoint string `yaml:"entrypoint"`
}

//OpenapiTagConfig openapi tag
type OpenapiTagConfig struct {
	Name        string
	Description string `yaml:"desc"`
}

type OpenapiServiceRouteConfig struct {
	Name    string
	Params  []*OpenapiParamConfig
	Returns []*OpenapiReturnConfig
	Tags    []string
}

type OpenapiGlobalApiConfig struct {
	Root string
}
type OpenapiParamConfig struct {
	Type         string
	Name         string
	Description  string
	DataType     string `yaml:"data-type"`
	Required     bool
	DefaultValue string `yaml:"default-value"`
}

type OpenapiReturnConfig struct {
	Code  int
	Msg   string
	Model string
}

type OpenapiServiceApisConfig struct {
	Routes map[string]*OpenapiServiceRouteConfig `yaml:"routes"`
	Tags   []string                              `yaml:"tags"`
}

type OpenapiServiceConfig map[string]*OpenapiServiceApisConfig

func (osc OpenapiServiceConfig) Service(key string) *OpenapiServiceApisConfig {
	return osc[key]
}

func (oac OpenapiServiceApisConfig) Route(key string) *OpenapiServiceRouteConfig {
	return oac.Routes[key]
}
