package rpc

import (
    "encoding/json"
    "github.com/stretchr/testify/assert"
    "project-my-test/example/rpcApp/method/methodGroup2"
    "project-my-test/src/rpc"
    "project-my-test/src/rpc/plugin/rpcStructureRequest"
    "project-my-test/src/rpc/rpcCore"
    "testing"
)

func TestRpc_correctWork(t *testing.T) {
    assert := assert.New(t)
    rpcClient := rpcCore.NewRpcCore()
    
    _tMethodName1 := "test2"
    _tMethodClass1 := new(methodGroup2.MethodMyTest)
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
    
    rpcRequest := rpc.NewRpcRequest("", "", nil, "", "")
    rpcRequest.SetMethodName(_tMethodName1)
    
    rpcResponse := rpcClient.RunMethodByRpcRequest(rpcRequest)
    assert.Equal(rpcResponse.GetError().GetCode(), "ERROR", "MethodInfo test - empty")
    
}

func TestRpc_correctWork2as(t *testing.T) {
    assert := assert.New(t)
    assert.NotNil(true)
    /*
    	foods := map[string]interface{}{
    		"bacon": "delicious",
    		"eggs": struct {
    			source string
    			price  float64
    		}{"chicken", 1.75},
    		"steak": true,
    	}
    */
    //
    // methodRun := &methodGroup2.MethodMyTest{
    //	Data: struct {
    //		Name  *string
    //		Email *string
    //	}{Name: "str", Email: "str"},
    // }
    
}
func TestRpc_correctWorkasdf(t *testing.T) {
    assert := assert.New(t)
    assert.NotNil(true)
    
    jsonString := `{
  "specification": "multipart-json-rpc",
  "specificationVersion" : "0.1.0",

  "version": "1",
  "language": "en",

  "request" : {
    "id"   : "9423234",
    "name" : "MethodTest1",
    "data" : {
      "full_name" : "work",
      "number" : 40,
      "bool": false
    }
  }
}`
    
    jsonByte := []byte(jsonString)
    
    // декодируем в структуру
    reqStruct := new(rpcStructureRequest.MultipartJsonRpcRequest)
    
    err := json.Unmarshal(jsonByte, reqStruct)
    assert.Nil(err)
    
    // получаем объект запроса
    rpcRequest := reqStruct.ConvertToRpcRequest()
    assert.NotNil(rpcRequest)
}
