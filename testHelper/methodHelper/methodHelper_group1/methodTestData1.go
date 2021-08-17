package methodHelper_group1

import (
	"project-my-test/src/rpc"
	"project-my-test/src/rpc/rpcInterface"
	"project-my-test/src/rpc/rpcStruct"
	"project-my-test/testHelper/reformHelper"
)

type MethodTestData1 struct {
	rpc.RpcMethod
	Data struct{
		TestUserString 	*string
		TestUserInt 	*int
	}
	// di
	//Logger *iLogger
}

func (r *MethodTestData1) GetRequestSchema() map[string]rpcStruct.ReformSchema {
	rs := make(map[string]rpcStruct.ReformSchema)

	rs["TestUserString"] = rpcStruct.ReformSchema{
		Type: reformHelper.REFORM_STRING,
		Field: "full_name",
	}
	rs["TestUserInt"] = rpcStruct.ReformSchema{
		Type: reformHelper.REFORM_INTEGER,
		Field: "number",
	}
	return rs
}

func (r *MethodTestData1) Run() rpcInterface.Response {
	r.Response.SetData("test::string", "MethodTestData1.go string")
	r.Response.SetData("test::json::full_name"	, r.Data.TestUserString)
	r.Response.SetData("test::json::number"		, r.Data.TestUserInt)
	return r.Response
}
