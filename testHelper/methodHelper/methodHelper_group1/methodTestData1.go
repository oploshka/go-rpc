package methodHelper_group1

import (
	"project-my-test/src/Rpc"
	"project-my-test/testHelper/reformHelper"
)

type MethodTestData1 struct {
	Rpc.RpcMethod
	Data struct{
		TestUserString 	*string
		TestUserInt 	*int
	}
	// di
	//Logger *iLogger
}

func (r *MethodTestData1) GetRequestSchema() map[string]Rpc.RpcReformSchema {
	rs := make(map[string]Rpc.RpcReformSchema)

	rs["TestUserString"] = Rpc.RpcReformSchema{
		Type: reformHelper.REFORM_STRING,
		Field: "full_name",
	}
	rs["TestUserInt"] = Rpc.RpcReformSchema{
		Type: reformHelper.REFORM_INTEGER,
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
