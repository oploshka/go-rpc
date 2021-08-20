package rpcInterface

import (
    "project-my-test/src/rpc/rpcStruct"
)

type Method interface {
    GetRequestSchema() map[string]rpcStruct.ReformSchema
    Run() Response
    SetResponse(rRes Response)
    
    // GetDescription() string
    // GetResponseSchema() string
}
