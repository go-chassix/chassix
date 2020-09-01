package main

import (
	"fmt"

	"c6x.io/chassix.v2"
	"c6x.io/chassix.v2/apierrors"
	"c6x.io/chassix.v2/config"
	"c6x.io/chassix.v2/restful"
)

func main() {

	config.LoadFromEnvFile()
	ws := restful.NewWebService("users")
	ws.Path("users")
	ws.Produces(restful.MediaTypeApplicationJson)
	ws.Consumes(restful.MediaTypeApplicationJson)

	ws.Route(ws.GET("").To(queryUsers))
	ws.Route(ws.GET("/{id}").To(getUser))

	//strings web service
	ws2 := restful.NewWebService("strings")
	ws2.Path("strings")
	ws2.Produces(restful.MediaTypeApplicationJson)
	ws2.Consumes(restful.MediaTypeApplicationJson)
	ws2.Route(ws2.GET("").To(strings))

	restful.RegisterModels(registerModels)

	ws.AddDocs()
	ws2.AddDocs()

	restful.Add(ws)
	restful.Add(ws2)
	chassix.ServeRestful()
}

func registerModels() {
	restful.RegisterModel(UserDTO{})
	restful.RegisterModel(&UserDTO{})
	restful.RegisterModelName("userDTOPageRes", UserPageRes{})
	restful.RegisterModel(stringArray{})
}

func strings(ctx *restful.Context) {
	strs := []string{"a", "b", "c"}
	ctx.Response.WriteAsJson(strs)
}

func queryUsers(ctx *restful.Context) {
	ctx.Request.QueryParameter("uid")
	//UserPageRes{Data: apierrors.APIError{}}
	//ctx.Response.WriteAsJson()
	chassix.RestResponse(ctx).Ok([]UserDTO{{Name: "test", Age: 18}})
}
func getUser(ctx *restful.Context) {
	id := ctx.Request.PathParameter("1")
	fmt.Println("id: " + id)
	ctx.Response.WriteAsJson(UserDTO{Name: "test", Age: 18})
}

type stringArray []string

func (sa stringArray) ModelName() string {
	return "stringArray"
}

type UserDTO struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (u UserDTO) ModelName() string {
	return "userDTO"
}

type UserPageRes struct {
	apierrors.APIError
	Data struct {
		chassix.PageDTO
		List []UserDTO `json:"list"`
	} `json:"data"`
}
