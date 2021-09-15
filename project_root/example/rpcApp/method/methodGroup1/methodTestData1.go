package methodGroup1

import (
    "project-my-test/example/rpcApp/methodRequestShemaItem"
    "project-my-test/src/rpc"
    "project-my-test/src/rpc/rpcInterface"
    "project-my-test/src/rpc/rpcStruct"
)

type MethodTestData1 struct {
    rpc.RpcMethod
    Data struct {
        TestUserString *string
        TestUserInt    *int
        TestUserBool   *bool
    }
    // di
    // Logger *iLogger
}

func (r *MethodTestData1) GetRequestSchema() map[string]rpcStruct.ReformSchema {
    rs := make(map[string]rpcStruct.ReformSchema)
    rs["TestUserString"]    = methodRequestShemaItem.GROUP1_TEST_USET_STRING()
    rs["TestUserInt"]       = methodRequestShemaItem.GROUP1_TEST_USET_INT()
    rs["TestUserBool"]      = methodRequestShemaItem.GROUP1_TEST_USET_BOOL()
    return rs
}

func (r *MethodTestData1) Run() rpcInterface.Response {
    r.Response.SetData("test::string", "MethodTestData1.go string")
    r.Response.SetData("test::json::full_name", r.Data.TestUserString)
    r.Response.SetData("test::json::number", r.Data.TestUserInt)
    r.Response.SetData("test::json::bool", r.Data.TestUserBool)
    return r.Response
}
