package MethodTestData

import (
	"project-my-test/src/Rpc"
)

type MethodTestData2 struct {
	Rpc.RpcMethod
	Data struct{}
}

func (r *MethodTestData2) GetRequestSchema() map[string]Rpc.RpcReformSchema {
	return make(map[string]Rpc.RpcReformSchema)
}


func (r *MethodTestData2) Run() *Rpc.RpcResponse {
	r.Response.SetData("test::string", "string - MethodTestData2.go")
	r.Response.SetData("test::int", 222)
	return r.Response
}
