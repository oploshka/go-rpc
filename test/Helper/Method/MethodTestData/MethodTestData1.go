package MethodTestData

import (
	"project-my-test/src/Rpc"
)

type MethodTestData1 struct {
	Rpc.RpcMethod
	Data struct{
		TestUserString 	*string
		TestUserInt 	*int
	}
	// di
	// Logger *iLogger
}

func (r *MethodTestData1) GetRequestSchema() map[string]Rpc.RpcReformSchema {
	rs := make(map[string]Rpc.RpcReformSchema)

	rs["TestUserString"] = Rpc.RpcReformSchema{
		Type: "STRING",
		Field: "full_name",
	}
	rs["TestUserInt"] = Rpc.RpcReformSchema{
		Type: "INT",
		Field: "number",
	}
	return rs
}

func (r *MethodTestData1) Run() *Rpc.RpcResponse {
	r.Response.SetData("test::string", "MethodTestData1.go string")
	r.Response.SetData("test::json::full_name"	, r.Data.TestUserString)
	r.Response.SetData("test::json::number"		, r.Data.TestUserInt)
	return r.Response
}
