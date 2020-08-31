package restful

import (
	emkRestful "github.com/emicklei/go-restful/v3"
)

type RouteBuilder struct {
	*emkRestful.RouteBuilder
	subpath string
}

const paramTypeQuery = "query"
const paramTypePath = "path"

// If this route is matched with the incoming Http Request then call this function with the *Context. Required.
func (b *RouteBuilder) To(function RouteFunction) *RouteBuilder {
	b.RouteBuilder.To(convert(function))
	return b
}

// Convert route call function with the *Request,*Response pair to *Context
func convert(function RouteFunction) func(request *emkRestful.Request, response *emkRestful.Response) {
	return func(request *emkRestful.Request, response *emkRestful.Response) {
		ctx := &Context{
			Request:  request,
			Response: response,
		}
		function(ctx)
	}
}
