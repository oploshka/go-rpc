package rpc

import (
	"github.com/stretchr/testify/assert"
	"project-my-test/src/rpc"
	"project-my-test/testHelper/methodHelper/methodHelper_group2"
	"testing"
)

func TestRpc_correctWork(t *testing.T) {
	assert := assert.New(t)
	rpcClient := rpc.NewRpcCore()

	_tMethodName1 := "test2"
	_tMethodClass1 := new(methodHelper_group2.ReturnRequestSchemaData)
	_tMethodGroup1 := "group2"

	methodStore1 := rpcClient.GetRpcMethodStore()
	assert.NotNil(methodStore1, "RpcMethodStorage not nil")

	methodStore1.Add(_tMethodName1, _tMethodClass1, _tMethodGroup1)

	methodStore2 := rpcClient.GetRpcMethodStore()
	methodInfo, ok := methodStore2.GetMethodInfo(_tMethodName1)

	assert.Equal(ok, true, "MethodInfo test - empty")
	assert.Equal(methodInfo.GetClass(), _tMethodClass1, "MethodInfo - not correct class")
	assert.Equal(methodInfo.GetGroup(), _tMethodGroup1, "MethodInfo - not correct group")

	_, ok2 := methodStore2.GetMethodInfo("a124sz1dcssa")
	assert.NotEqual(ok2, true, "MethodInfo test - empty")

	rpcRequest := rpc.NewRpcRequest()
	rpcRequest.SetMethodName(_tMethodName1)

	rpcResponse := rpcClient.RunMethodByRpcRequest(rpcRequest)
	assert.Equal(rpcResponse.GetError().GetCode(), "ERROR", "MethodInfo test - empty")

}
