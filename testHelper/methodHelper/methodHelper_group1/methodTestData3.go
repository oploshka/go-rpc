package methodHelper_group1

import (
	"project-my-test/src/Rpc"
)

type MethodTestData3 struct {
	Rpc.RpcMethod
	Data struct{}
}

func (r *MethodTestData3) GetRequestSchema() map[string]Rpc.RpcReformSchema {
	return make(map[string]Rpc.RpcReformSchema)
}


func (r *MethodTestData3) Run() *Rpc.RpcResponse {
	r.Response.SetData("test::string", "string")
	r.Response.SetData("test::int", 20)
	return r.Response
}
