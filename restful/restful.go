package restful

import (
	"c6x.io/chassix.v2/config"
	restFilters "c6x.io/chassix.v2/filters/rest"
	"c6x.io/chassix.v2/logx"
	restfulSpec "github.com/emicklei/go-restful-openapi/v2"
	emkRestful "github.com/emicklei/go-restful/v3"
	"net/http"
	"reflect"
	"strconv"
)

const MediaTypeApplicationJson = "application/json"

type Model interface {
	ModelName() string
}

var registeredModels = make(map[string]interface{})

func RegisterModel(m Model) {
	RegisterModelName(m.ModelName(), m)
}
func RegisterModelName(name string, m interface{}) {
	v := reflect.ValueOf(m)
	t := v.Type()
	if v.Kind() == reflect.Ptr {
		t = v.Elem().Type()
		//return
	}

	if v := reflect.Zero(t); v.CanInterface() {
		registeredModels[name] = v.Interface()
	}
}
func RegisteredModel(dtoName string) (interface{}, bool) {
	t, ok := registeredModels[dtoName]
	return t, ok
}
func StartService() {
	log := logx.New().Service("chassix").Category("restful")

	emkRestful.Filter(restFilters.RequestID)
	emkRestful.Filter(restFilters.MeasureTime)

	openApiCfg := config.Openapi()
	//if enable openapi setting. register swagger ui and apidocs json API.
	if openApiCfg.Enabled {
		swaggerUICfg := config.Openapi().UI
		//定义swagger文档
		cfg := restfulSpec.Config{
			WebServices:                   emkRestful.RegisteredWebServices(), // you control what services are visible
			APIPath:                       swaggerUICfg.API,
			PostBuildSwaggerObjectHandler: newPostBuildOpenAPIObjectFunc()}
		emkRestful.DefaultContainer.Add(restfulSpec.NewOpenAPIService(cfg))
		http.Handle(swaggerUICfg.Entrypoint, http.StripPrefix(swaggerUICfg.Entrypoint, http.FileServer(http.Dir(swaggerUICfg.Dist))))
	}
	//启动服务
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(config.Server().Port), nil))
}
