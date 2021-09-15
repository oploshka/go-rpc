package temp

import (
    "github.com/stretchr/testify/assert"
    "project-my-test/example/rpcApp/method/methodGroup1"
    "project-my-test/example/rpcApp/method/methodGroup2"
  "project-my-test/src/service"
    "project-my-test/src/rpc/rpcCore"
    "testing"
)

func TestMethodDiSet(t *testing.T) {
    assert := assert.New(t)
    logger := service.CreateLogger()
    // get method schema
    rpcMethod := new(methodGroup2.MethodMyTest)
    
    assert.Nil(rpcMethod.Logger)
    
    rpcCore.RpcMethodDiInit(rpcMethod, logger)
    
    assert.NotNil(rpcMethod.Logger)
    rpcMethod.Logger.Info("test Info")
    rpcMethod.Logger.Warning("test Warning")
    rpcMethod.Logger.Debug("test Debug")
}


func TestMethodDiEmptySet(t *testing.T) {
    assert := assert.New(t)
    logger := service.CreateLogger()
    
    // get method schema
    rpcMethod := new(methodGroup1.MethodTestData3)
    
    rpcCore.RpcMethodDiInit(rpcMethod, logger)
    
    assert.NotNil(true)
}