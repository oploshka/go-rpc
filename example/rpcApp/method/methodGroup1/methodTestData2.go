package methodGroup1

import (
    "project-my-test/src/rpc"
    "project-my-test/src/rpc/rpcInterface"
)

type MethodTestData2 struct {
    rpc.RpcMethod
}

func (r *MethodTestData2) Run() rpcInterface.Response {
    r.Response.SetData("test::string", "string - MethodTestData2.go")
    r.Response.SetData("test::int", 222)
    return r.Response
}
