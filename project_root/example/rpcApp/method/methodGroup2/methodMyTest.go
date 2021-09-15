package methodGroup2

import (
    "project-my-test/example/rpcApp/methodRequestShemaItem"
    "project-my-test/src/rpc"
    "project-my-test/src/rpc/rpcInterface"
    "project-my-test/src/rpc/rpcStruct"
)

type MethodMyTest struct {
    rpc.RpcMethod
    Data struct {
        Name  *string
        Email *string
    }
    Logger rpcInterface.Logger
}

func (r *MethodMyTest) GetRequestSchema() map[string]rpcStruct.ReformSchema {
    rs := make(map[string]rpcStruct.ReformSchema)
    
    rs["Name"]  = methodRequestShemaItem.GROUP2_NAME()
    rs["Email"] = methodRequestShemaItem.GROUP2_EMAIL()
    return rs
}

func (r *MethodMyTest) Run() rpcInterface.Response {
    
    r.Logger.Info("test Info")
    r.Logger.Warning("test Warning")
    r.Logger.Error("test Error")
    r.Logger.Debug("test Debug")
    
    r.Response.SetData("test::string", "string")
    r.Response.SetData("test::int", 20)
    r.Response.GetError().SetCode("ERROR")
    return r.Response
}
