package MethodTestData

import (
	"project-my-test/src/Rpc"
)

type MethodTestData2 struct {
	Rpc.RpcMethod
}

func (r *MethodTestData2) Run() *Rpc.RpcResponse {
	r.Response.SetData("test::string", "string - MethodTestData2.go")
	r.Response.SetData("test::int", 222)
	return r.Response
}
