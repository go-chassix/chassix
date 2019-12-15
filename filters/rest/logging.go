package restfilters

import (
	"time"

	"github.com/emicklei/go-restful"
	uuid "github.com/satori/go.uuid"

	xLog "pgxs.io/chassis/log"
)

var log *xLog.Entry

func init() {
	log = xLog.New().Service("chassis").Category("filter")

}

//RequestID Filter
func RequestID(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	if reqID := uuid.NewV4(); &reqID != nil && reqID.String() != "" {
		req.SetAttribute("reqId", reqID.String())
		req.Request.Header.Add("X-Request-Id", reqID.String())
		resp.AddHeader("X-Request-Id", reqID.String())
	}

	chain.ProcessFilter(req, resp)
}

//MeasureTime time filter
func MeasureTime(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	now := time.Now()
	chain.ProcessFilter(req, resp)
	log.SetReqInfo(req).Infof("time: %v", time.Now().Sub(now))
}
