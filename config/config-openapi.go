package config

import "strings"

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
	Tags         []OpenapiTagConfig     `yaml:",flow"`
	UI           OpenapiUIConfig        `yaml:"ui"`
	GlobalApi    OpenapiGlobalApiConfig `yaml:"global-api"`
	Resources    OpenapiResourcesConfig `yaml:"resources,flow"`
	resourcesMap map[string]*OpenapiServiceConfig
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

type OpenapiResourcesConfig []*OpenapiServiceConfig
type OpenapiServiceRouteConfig struct {
	Path    string
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

type OpenapiServiceConfig struct {
	Name     string
	Routes   []*OpenapiServiceRouteConfig `yaml:"routes"`
	Tags     []string                     `yaml:"tags"`
	routeMap map[string]*OpenapiServiceRouteConfig
}

//type OpenapiServiceConfig []*OpenapiServiceConfig

func (osc OpenapiResourcesConfig) Service(key string) *OpenapiServiceConfig {
	if config != nil && len(Openapi().Resources) > 0 {
		if svc, ok := Openapi().resourcesMap[key]; ok {
			return svc
		}
	}

	return nil
}

func (osc OpenapiServiceConfig) Route(key string) *OpenapiServiceRouteConfig {
	return osc.routeMap[key]
}

var resourcesMap map[string]*OpenapiServiceConfig

func (osc OpenapiResourcesConfig) _copyResourcesToMap() {
	for _, resource := range osc {
		if resource != nil {
			if config.OpenAPI.resourcesMap == nil {
				config.OpenAPI.resourcesMap = make(map[string]*OpenapiServiceConfig)
			}
			if resource.Name != "" {
				config.OpenAPI.resourcesMap[resource.Name] = resource
			}
		}
	}
}
func (osc OpenapiResourcesConfig) _copyResourceRoutesToMap() {
	for _, svc := range osc {
		if svc.routeMap == nil {
			svc.routeMap = make(map[string]*OpenapiServiceRouteConfig)
		}
		for _, route := range svc.Routes {
			if strings.Trim(route.Path, " ") != "" {

				svc.routeMap[route.Path] = route
			}
		}
	}
}
func (osc OpenapiResourcesConfig) CopyToMap() {
	osc._copyResourcesToMap()
	osc._copyResourceRoutesToMap()
}
