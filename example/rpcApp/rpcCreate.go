package rpcApp

import (
    "project-my-test/src/rpc"
    "project-my-test/example/rpcApp/method/methodGroup1"
    "project-my-test/example/rpcApp/method/methodGroup2"
)

func CreateTestRpc() *rpc.RpcCore {
    
    rpcClient := rpc.NewRpcCore()
    rpcMethodStore := rpcClient.GetRpcMethodStore()
    //
    rpcMethodStore.Add("testMethod", new(methodGroup2.MethodMyTest), "")
    rpcMethodStore.Add("MethodTestData1", new(methodGroup1.MethodTestData1), "TestData")
    rpcMethodStore.Add("MethodTestData2", new(methodGroup1.MethodTestData2), "TestData")
    rpcMethodStore.Add("MethodTestData3", new(methodGroup1.MethodTestData3), "TestData")
    
    return rpcClient
}
