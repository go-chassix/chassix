package restful

import (
	emkRestful "github.com/emicklei/go-restful/v3"
)

//Context context request and response
type Context struct {
	Request  *emkRestful.Request
	Response *emkRestful.Response
}

type RouteFunction func(ctx *Context)

type Route struct {
	*emkRestful.Route
}

type RouteBuilder struct {
	*emkRestful.RouteBuilder
}

// If this route is matched with the incoming Http Request then call this function with the *Request,*Response pair. Required.
func (b *RouteBuilder) To(function RouteFunction) *RouteBuilder {
	b.RouteBuilder.To(convert(function))
	return b
}
func convert(function RouteFunction) func(request *emkRestful.Request, response *emkRestful.Response) {
	return func(request *emkRestful.Request, response *emkRestful.Response) {
		ctx := &Context{
			Request:  request,
			Response: response,
		}
		function(ctx)
	}
}

//func ()
//// RouteBuilder is a helper to construct Routes.
//type RouteBuilder struct {
//	rootPath    string
//	currentPath string
//	produces    []string
//	consumes    []string
//	httpMethod  string        // required
//	function    RouteFunction // required
//	filters     []FilterFunction
//	conditions  []RouteSelectionConditionFunction
//
//	typeNameHandleFunc TypeNameHandleFunction // required
//
//	// documentation
//	doc                     string
//	notes                   string
//	operation               string
//	readSample, writeSample interface{}
//	parameters              []*Parameter
//	errorMap                map[int]ResponseError
//	defaultResponse         *ResponseError
//	metadata                map[string]interface{}
//	deprecated              bool
//	contentEncodingEnabled  *bool
//}
