package Helper

//import (
//	"github.com/stretchr/testify/assert"
//	"project-my-test/src/Rpc"
//	"project-my-test/test/Helper/Method"
//)
//
//func GetTestRpc() Rpc.RpcCore {
//
//	rpcRequest := Rpc.NewRpcCore()
//
//	_tMethodName1 	:= "test2"
//	_tMethodClass1 	:= "class2"
//	_tMethodGroup1 	:= "group2"
//
//	methodStore1 := rpcRequest.GetRpcMethodStore()
//	assert.NotNil(methodStore1,  "RpcMethodStorage not nil")
//
//	methodStore1.Add(_tMethodName1, new(Method.ReturnRequestSchemaData), _tMethodGroup1)
//
//	methodStore2 := rpcRequest.GetRpcMethodStore()
//	methodInfo, ok := methodStore2.GetMethodInfo(_tMethodName1)
//}
