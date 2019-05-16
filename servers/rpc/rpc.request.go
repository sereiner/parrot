package rpc

import (
	"github.com/sereiner/parrot/servers/rpc/pb"
	"github.com/sereiner/library/jsons"
	"github.com/sereiner/library/types"
)

type Request struct {
	*pb.RequestContext
	header map[string]string
	input  map[string]interface{}
}

func (r *Request) GetHeader() map[string]string {
	if r.header == nil {
		hm, _ := jsons.Unmarshal([]byte(r.RequestContext.Header))
		r.header, _ = types.ToStringMap(hm)
	}
	return r.header
}
func (r *Request) GetForm() map[string]interface{} {
	if r.input == nil {
		r.input, _ = jsons.Unmarshal([]byte(r.RequestContext.Input))
	}
	return r.input
}
