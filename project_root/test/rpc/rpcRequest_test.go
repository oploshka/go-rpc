package rpc

import (
    "github.com/stretchr/testify/assert"
    "project-my-test/src/rpc"
    "testing"
)

func TestRpcRequest_correctWork(t *testing.T) {
    assert := assert.New(t)
    
    // t.Run("test1", func(t *testing.T) {
    //
    rpcRequest := rpc.NewRpcRequest("", "", nil, "", "")
    
    assert.Equal(rpcRequest.GetRequestId(), "", "rpcRequest error")
    assert.Equal(rpcRequest.GetMethodName(), "", "rpcRequest error")
    // TODO: fix
    // assert.Equal(rpcRequest.GetData(), nil, "rpcRequest error")
    assert.Equal(rpcRequest.GetLanguage(), "", "rpcRequest error")
    assert.Equal(rpcRequest.GetVersion(), "", "rpcRequest error")
    // })
    
}
