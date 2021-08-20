package methodGroup2

import (
    "project-my-test/src/rpc"
    "project-my-test/src/rpc/rpcInterface"
    "project-my-test/src/rpc/rpcStruct"
    "project-my-test/example/rpcApp/methodRequestShemaItem"
)

type MethodMyTest struct {
    rpc.RpcMethod
    Data struct {
        Name  *string
        Email *string
    }
}

func (r *MethodMyTest) GetRequestSchema() map[string]rpcStruct.ReformSchema {
    rs := make(map[string]rpcStruct.ReformSchema)
    
    rs["Name"]  = methodRequestShemaItem.GROUP2_NAME()
    rs["Email"] = methodRequestShemaItem.GROUP2_EMAIL()
    return rs
}

func (r *MethodMyTest) Run() rpcInterface.Response {
    r.Response.SetData("test::string", "string")
    r.Response.SetData("test::int", 20)
    r.Response.GetError().SetCode("ERROR")
    return r.Response
}
