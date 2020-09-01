package restful

import (
	"fmt"
	emkRestful "github.com/emicklei/go-restful/v3"
	"path"

	"c6x.io/chassix.v2/config"
)

//Webservice web service
type WebService struct {
	name string
	emkRestful.WebService
}

func NewWebService(name string) *WebService {
	return &WebService{name: name}
}
func (ws *WebService) Name() string {
	return ws.name
}

func (ws *WebService) Route(builder *RouteBuilder) *WebService {
	ws.WebService.Route(builder.RouteBuilder)
	return ws
}

//func (ws *WebService)GET()

// GET is a shortcut for .Method("GET").Path(subPath)
func (ws *WebService) GET(subPath string) *RouteBuilder {
	rb := &RouteBuilder{RouteBuilder: ws.WebService.GET(subPath)}
	rb.subpath = subPath
	return rb
}

func Add(ws *WebService) {
	emkRestful.Add(&ws.WebService)
}

func (ws *WebService) AddDocs() {
	if config.Openapi().Enabled {
		apisCfg := config.Openapi().Resources
		if apisCfg != nil && len(apisCfg) > 0 {
			apisCfg.CopyToMap()
		} else {
			return
		}
		routes := ws.Routes()
		for i, route := range routes {
			fmt.Printf("route %+v\n", route)

			if routesFromConfig := apisCfg.Service(ws.Name()); routesFromConfig != nil && routesFromConfig.Routes != nil {

				if apiCfg := routesFromConfig.Route(path.Join(config.Openapi().GlobalApi.Root, route.Path)); apiCfg != nil {
					if apiCfg.Name != "" {
						routes[i].Doc = apiCfg.Name
					}

					//add param docs
					if len(apiCfg.Params) > 0 {
						for _, api := range apiCfg.Params {
							if api != nil && api.Name != "" {
								var param *emkRestful.Parameter
								if api.Type == paramTypeQuery {
									param = ws.QueryParameter(api.Name, api.Description)
								} else if api.Type == paramTypePath {
									param = ws.PathParameter(api.Name, api.Description)
								}
								if param != nil {
									param.Required(api.Required)
									param.DataType(api.DataType)
									param.DefaultValue(api.DefaultValue)
								}

								routes[i].ParameterDocs = append(routes[i].ParameterDocs, param)

							}
						}
					}

					if len(apiCfg.Returns) > 0 {
						resErrors := make(map[int]emkRestful.ResponseError)
						for _, resErr := range apiCfg.Returns {

							if resErr != nil {
								if resErr.Code > 0 {
									emkResErr := emkRestful.ResponseError{Code: resErr.Code, Message: resErr.Msg}
									if resErr.Model != "" {
										if m, ok := RegisteredModel(resErr.Model); ok {
											emkResErr.Model = m
										}
									}

									resErrors[resErr.Code] = emkResErr

								}

							}
						}
						routes[i].ResponseErrors = resErrors
					}

					//add tags
					tags := apiCfg.Tags
					if len(tags) > 0 {
						if route.Metadata == nil {
							routes[i].Metadata = map[string]interface{}{}
						}
						appendTags(&(routes[i]), tags)
					}
				}

				appendTags(&routes[i], routesFromConfig.Tags)
			}
			continue

		}
	}
}

func appendTags(route *emkRestful.Route, tags []string) {

	if route.Metadata == nil {
		route.Metadata = make(map[string]interface{})
	}
	existedTags, ok := route.Metadata[KeyOpenAPITags]
	if !ok {
		//route.Metadata[KeyOpenAPITags] = make([]string, 0)
		existedTags = make([]string, 0)
	}

	if existedStrTags, isStrArray := existedTags.([]string); isStrArray {
		existedStrTags := append(existedStrTags, tags...)
		route.Metadata[KeyOpenAPITags] = existedStrTags
	}
}
