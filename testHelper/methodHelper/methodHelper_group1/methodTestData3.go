package methodHelper_group1

import (
	"project-my-test/src/rpc"
	"project-my-test/src/rpc/rpcInterface"
)

type MethodTestData3 struct {
	rpc.RpcMethod
}

func (r *MethodTestData3) Run() rpcInterface.Response {
	r.Response.SetData("test::string", "string")
	r.Response.SetData("test::int", 20)
	return r.Response
}
