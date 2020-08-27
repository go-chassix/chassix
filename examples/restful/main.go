package main

import (
	"net/http"

	"c6x.io/chassix.v2/restful"
)

func main() {

	ws := new(restful.WebService)
	ws.Route(ws.GET("/hello").To(hello))
	restful.Add(ws)
	http.ListenAndServe(":9000", nil)
}

func hello(ctx *restful.Context) {
	ctx.Request.QueryParameter("")
	names := []string{"a", "b", "c"}
	ctx.Response.WriteAsJson(names)
}
