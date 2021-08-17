package methodHelper_group1

import (
	"project-my-test/src/rpc"
	"project-my-test/src/rpc/rpcInterface"
	"project-my-test/src/rpc/rpcStruct"
)

type MethodTestData2 struct {
	rpc.RpcMethod
	Data struct{}
}

func (r *MethodTestData2) GetRequestSchema() map[string]rpcStruct.ReformSchema {
	return make(map[string]rpcStruct.ReformSchema)
}


func (r *MethodTestData2) Run() rpcInterface.Response {
	r.Response.SetData("test::string", "string - MethodTestData2.go")
	r.Response.SetData("test::int", 222)
	return r.Response
}
