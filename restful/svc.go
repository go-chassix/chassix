package restful

import (
	emkRestful "github.com/emicklei/go-restful/v3"
)

//Webservice web service
type WebService struct {
	emkRestful.WebService
}

func (ws *WebService) Route(builder *RouteBuilder) *WebService {
	ws.WebService.Route(builder.RouteBuilder)
	return ws
}

//func (ws *WebService)GET()

// GET is a shortcut for .Method("GET").Path(subPath)
func (ws *WebService) GET(subPath string) *RouteBuilder {
	return &RouteBuilder{RouteBuilder: ws.WebService.GET(subPath)}
}

func Add(ws *WebService) {
	emkRestful.Add(&ws.WebService)
}
