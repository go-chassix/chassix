package chassis

import (
	"strconv"

	"github.com/emicklei/go-restful"
)

//BaseController has some basic methods
type BaseController struct {
}

//ValidateResourceID validate REST resource ID in pathparam and return it if validated.
func (bc BaseController) ValidateResourceID(req *restful.Request, res *restful.Response, IDKey string) (ID uint, ok bool) {
	intID, err := strconv.Atoi(req.PathParameter(IDKey))
	if err != nil {
		NewResponse(res).Error(400, ErrIDInvalid)
		return 0, false
	}
	return uint(intID), true
}

//ValidatePageableParams validate pagination params return paable struct
func (bc BaseController) ValidatePageableParams(req *restful.Request, res *restful.Response) (pageable *Pageable, ok bool) {
	pi, err := strconv.Atoi(req.QueryParameter("page_index"))
	ps, err := strconv.Atoi(req.QueryParameter("page_size"))
	if err != nil {
		NewResponse(res).Error(400, ErrPageParamsInvalid)
		return nil, false
	}
	//todo sort
	return &Pageable{Page: uint(pi), Size: uint(ps)}, true
}
