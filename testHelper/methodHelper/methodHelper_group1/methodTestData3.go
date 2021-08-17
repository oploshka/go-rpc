package methodHelper_group1

import (
	"project-my-test/src/rpc"
	"project-my-test/src/rpc/rpcInterface"
	"project-my-test/src/rpc/rpcStruct"
)

type MethodTestData3 struct {
	rpc.RpcMethod
	Data struct{}
}

func (r *MethodTestData3) GetRequestSchema() map[string]rpcStruct.ReformSchema {
	return make(map[string]rpcStruct.ReformSchema)
}


func (r *MethodTestData3) Run() rpcInterface.Response {
	r.Response.SetData("test::string", "string")
	r.Response.SetData("test::int", 20)
	return r.Response
}
