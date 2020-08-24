package chassis

import (
	"net/http"
	"strconv"

	restfulSpec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/go-openapi/spec"

	"c6x.io/chassis/config"
	restFilters "c6x.io/chassis/filters/rest"
	"c6x.io/chassis/logx"
)

const (
	//PageIndexKey pagination page index default 0
	PageIndexKey = "page_index"
	//DefaultPageIndexValue pagination page index default 0
	DefaultPageIndexValue = 0
	//PageSizeKey pagination page size default 10
	PageSizeKey = "page_size"
	//DefaultPageSizeValue pagination page size default 10
	DefaultPageSizeValue = 10
)

// KeyOpenAPITags is a Metadata key for a restful Route
const KeyOpenAPITags = restfulSpec.KeyOpenAPITags

//newPostBuildOpenAPIObjectFunc open api api docs data
func newPostBuildOpenAPIObjectFunc() restfulSpec.PostBuildSwaggerObjectFunc {
	return func(swo *spec.Swagger) {
		swo.Host = config.Openapi().Host
		swo.BasePath = config.Openapi().BasePath
		swo.Schemes = config.Openapi().Schemas
		swo.Info = &spec.Info{
			InfoProps: spec.InfoProps{
				Title:       config.Openapi().Spec.Title,
				Description: config.Openapi().Spec.Description,
				Contact: &spec.ContactInfo{
					Name:  config.Openapi().Spec.Contact.Name,
					Email: config.Openapi().Spec.Contact.Email,
					URL:   config.Openapi().Spec.Contact.URL,
				},

				License: &spec.License{
					Name: config.Openapi().Spec.License.Name,
					URL:  config.Openapi().Spec.License.URL,
				},
				Version: config.Openapi().Spec.Version,
			},
		}

		var nTags []spec.Tag
		for _, tag := range config.Openapi().Tags {
			nTag := spec.Tag{TagProps: spec.TagProps{Name: tag.Name, Description: tag.Description}}

			nTags = append(nTags, nTag)
		}
		swo.Tags = nTags
	}
}

//Serve rest webservice
func Serve(svc []*restful.WebService) {
	log := logx.New().Service("chassis").Category("restful")

	restful.Filter(restFilters.RequestID)
	restful.Filter(restFilters.MeasureTime)

	//if enable openapi setting. register swagger ui and apidocs json API.
	if config.Openapi().Enabled {
		swaggerUICfg := config.Openapi().UI
		//定义swagger文档
		cfg := restfulSpec.Config{
			WebServices:                   svc, // you control what services are visible
			APIPath:                       swaggerUICfg.API,
			PostBuildSwaggerObjectHandler: newPostBuildOpenAPIObjectFunc()}
		restful.DefaultContainer.Add(restfulSpec.NewOpenAPIService(cfg))
		http.Handle(swaggerUICfg.Entrypoint, http.StripPrefix(swaggerUICfg.Entrypoint, http.FileServer(http.Dir(swaggerUICfg.Dist))))
	}
	//启动服务
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(config.Server().Port), nil))
}

//PageQueryParams get page params from request
func PageQueryParams(req *restful.Request) (pageIndex, pageSize uint) {
	// var err error
	if pi, err := strconv.Atoi(req.QueryParameter(PageIndexKey)); err == nil {
		pageIndex = uint(pi)
	} else {
		pageIndex = DefaultPageIndexValue

	}
	if pz, err := strconv.Atoi(req.QueryParameter(PageSizeKey)); err == nil {
		pageSize = uint(pz)
	} else {
		pageSize = DefaultPageSizeValue

	}
	return
}

//AddMetaDataTags add metadata tags to Webservice all routes
func AddMetaDataTags(ws *restful.WebService, tags []string) {
	routes := ws.Routes()
	for i, route := range routes {
		if route.Metadata == nil {
			routes[i].Metadata = map[string]interface{}{}
		}
		routeTags := routes[i].Metadata[KeyOpenAPITags]
		if routeTags != nil {
			existedTags, ok := routeTags.([]string)
			if ok {
				existedTags = append(existedTags, tags...)
				routes[i].Metadata[KeyOpenAPITags] = existedTags
			}
			continue
		}
		routes[i].Metadata[KeyOpenAPITags] = tags
	}
}

//AddWriteSample setting a webservice all routes to  write sample with ResponseEntity
func AddWriteSample(ws *restful.WebService, entityType interface{}) {
	routes := ws.Routes()
	for i := range routes {
		routes[i].WriteSample = ResponseEntitySample{Data: entityType}
	}
}

// AddMetaDataTagsAndWriteSample  AddMetaDataTags() and  WriteSample
func AddMetaDataTagsAndWriteSample(ws *restful.WebService, tags []string, entityType interface{}) {
	AddMetaDataTags(ws, tags)
	AddWriteSample(ws, entityType)
}

//NewWriteSample new write sample
func NewWriteSample(entity interface{}) ResponseEntitySample {
	return ResponseEntitySample{
		Data: entity,
	}
}
