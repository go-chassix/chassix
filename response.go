package chassix

import (
	emkRestful "github.com/emicklei/go-restful/v3"
	"github.com/go-chassix/chassix/v2/restful"

	"github.com/go-chassix/chassix/v2/apierrors"
)

//Entity response entity
type Entity struct {
	*apierrors.APIError
	Data interface{} `json:"data,omitempty"`
}

//ResponseEntity response entity for go-restful Writes(ResponseEntity{Data: Type{}})
type ResponseEntitySample struct {
	ErrCode int         `json:"err_code,omitempty"`
	ErrMsg  string      `json:"err_msg,omitempty"`
	ErrDesc string      `json:"err_desc,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

//Response rest response
type Response struct {
	body       Entity
	httpStatus int
	res        *emkRestful.Response
}

//NewResponse new return
func NewResponse(res *emkRestful.Response) *Response {
	return &Response{res: res}
}
func RestResponse(ctx *restful.Context) *Response {
	return &Response{res: ctx.Response}
}

//Ok 200 return
func (re *Response) Ok(entity interface{}) {
	re.body.Data = entity
	re.httpStatus = 200
	re.body.APIError = apierrors.DefaultAPIError
	re.writeHeaderAndEntity()
}

//Created 201 return
func (re *Response) Created(entity interface{}) {
	re.body.Data = entity
	re.httpStatus = 201
	re.body.APIError = apierrors.DefaultAPIError
	re.writeHeaderAndEntity()
}

//Error error response
func (re *Response) Error(status int, err *apierrors.APIError) {
	re.body.APIError = err
	re.httpStatus = status
	re.writeHeaderAndEntity()
}
func (re *Response) Status(statusCode int) *Response {
	re.httpStatus = statusCode
	return re
}
func (re *Response) Entity(entity interface{}) {
	re.body.Data = entity
	re.writeHeaderAndEntity()
}
func (re *Response) writeHeaderAndEntity() {
	re.res.WriteHeaderAndEntity(re.httpStatus, re.body)
}
