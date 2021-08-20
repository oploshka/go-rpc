package rpcHelper

import (
    "project-my-test/src/rpc"
    "project-my-test/testHelper/methodHelper/methodHelper_group1"
    "project-my-test/testHelper/methodHelper/methodHelper_group2"
)

func CreateTestRpc() *rpc.RpcCore {
    
    rpcClient := rpc.NewRpcCore()
    rpcMethodStore := rpcClient.GetRpcMethodStore()
    //
    rpcMethodStore.Add("testMethod", new(methodHelper_group2.ReturnRequestSchemaData), "")
    rpcMethodStore.Add("MethodTestData1", new(methodHelper_group1.MethodTestData1), "TestData")
    rpcMethodStore.Add("MethodTestData2", new(methodHelper_group1.MethodTestData2), "TestData")
    rpcMethodStore.Add("MethodTestData3", new(methodHelper_group1.MethodTestData3), "TestData")
    
    return rpcClient
}
