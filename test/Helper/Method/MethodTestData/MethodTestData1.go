package MethodTestData

import (
	"project-my-test/src/Rpc"
)

type MethodTestData1 struct {
	Rpc.RpcMethod
	Data struct{
		TestUserData *string
	}
}

func (r *MethodTestData1) GetRequestSchema() map[string]Rpc.RpcReformSchema {
	rs := make(map[string]Rpc.RpcReformSchema)

	rs["TestUserData"] = Rpc.RpcReformSchema{
		Type: "STRING",
		Field: "full_name", // TODO: FieldLoad and FieldReturn
	}
	return rs
}

func (r *MethodTestData1) Run() *Rpc.RpcResponse {
	r.Response.SetData("test::string", "MethodTestData1.go string")
	r.Response.SetData("test::json::full_name", r.Data.TestUserData)
	return r.Response
}
