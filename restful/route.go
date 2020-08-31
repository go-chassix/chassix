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
