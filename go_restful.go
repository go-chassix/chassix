package chassis

import (
	"net/http"
	"strconv"

	"github.com/emicklei/go-restful"
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/go-openapi/spec"
	log "gopkg.in/logger.v1"

	"pgxs.io/panguxs/pkg/chassis/config"
	restfilters "pgxs.io/panguxs/pkg/chassis/filters/rest"
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
const KeyOpenAPITags = restfulspec.KeyOpenAPITags

//newPostBuildOpenAPIObjectFunc open api api docs data
func newPostBuildOpenAPIObjectFunc() restfulspec.PostBuildSwaggerObjectFunc {
	return func(swo *spec.Swagger) {
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
					Name: "",
					URL:  "",
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

//Start rest webservice
func Start(svcs []*restful.WebService) {

	restful.Filter(restfilters.RequestID)
	restful.Filter(restfilters.MeasureTime)

	ucfg := config.Openapi().UI
	//定义swagger文档
	cfg := restfulspec.Config{
		WebServices:                   svcs, // you control what services are visible
		APIPath:                       ucfg.API,
		PostBuildSwaggerObjectHandler: newPostBuildOpenAPIObjectFunc()}
	restful.DefaultContainer.Add(restfulspec.NewOpenAPIService(cfg))
	http.Handle(ucfg.Entrypoint, http.StripPrefix(ucfg.Entrypoint, http.FileServer(http.Dir(ucfg.Dist))))
	//启动服务
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(config.Server().Port), nil))
}

//PageQueryParamters get page params from request
func PageQueryParamters(req *restful.Request) (pageIndex, pageSize uint) {
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
