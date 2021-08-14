package Method

import (
	"project-my-test/src/Rpc"
)

type ReturnRequestSchemaData struct {
	Rpc.RpcMethod
}

func (r *ReturnRequestSchemaData) Run() *Rpc.RpcResponse {
	r.Response.SetData("test::string", "string")
	r.Response.SetData("test::int", 20)
	r.Response.GetError().SetCode("ERROR")
	return r.Response
}
