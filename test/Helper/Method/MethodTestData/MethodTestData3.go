package MethodTestData

import (
	"project-my-test/src/Rpc"
)

type MethodTestData3 struct {
	Rpc.RpcMethod
}

func (r *MethodTestData3) Run() *Rpc.RpcResponse {
	r.Response.SetData("test::string", "string")
	r.Response.SetData("test::int", 20)
	return r.Response
}
