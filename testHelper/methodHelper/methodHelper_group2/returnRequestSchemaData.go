package methodHelper_group2

import (
    "project-my-test/src/rpc"
    "project-my-test/src/rpc/rpcInterface"
    "project-my-test/src/rpc/rpcStruct"
    "project-my-test/testHelper/reformHelper"
)

type ReturnRequestSchemaData struct {
    rpc.RpcMethod
    Data struct {
        Name  *string
        Email *string
    }
}

func (r *ReturnRequestSchemaData) GetRequestSchema() map[string]rpcStruct.ReformSchema {
    rs := make(map[string]rpcStruct.ReformSchema)
    
    rs["Name"] = rpcStruct.ReformSchema{
        Type:  reformHelper.REFORM_STRING,
        Field: "name", // TODO: FieldLoad and FieldReturn
        Default: func() interface{} {
            return "UserNameDefault"
        },
    }
    rs["Email"] = rpcStruct.ReformSchema{
        Type:  reformHelper.REFORM_STRING,
        Field: "email", // TODO: FieldLoad and FieldReturn
    }
    // RequestSchema struct{ Name, Email string `mytag:"MyEmail"` }
    return rs
}

func (r *ReturnRequestSchemaData) Run() rpcInterface.Response {
    r.Response.SetData("test::string", "string")
    r.Response.SetData("test::int", 20)
    r.Response.GetError().SetCode("ERROR")
    return r.Response
}
