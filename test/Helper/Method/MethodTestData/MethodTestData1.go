package MethodTestData

import (
	"project-my-test/src/Rpc"
)

type MethodTestData1 struct {
	Rpc.RpcMethod
}

func (r *MethodTestData1) Run() *Rpc.RpcResponse {
	r.Response.SetData("test::string", "MethodTestData1.go string")
	r.Response.SetData("test::int", 111)
	return r.Response
}
