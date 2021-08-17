package rpc

import (
	"github.com/stretchr/testify/assert"
	"project-my-test/src/Rpc"
	"testing"
)

func TestRpcRequest_correctWork(t *testing.T) {
	assert := assert.New(t)

	//t.Run("test1", func(t *testing.T) {
	//
	rpcRequest := Rpc.NewRpcRequest()

	assert.Equal(rpcRequest.GetRequestId(), "", "rpcRequest error")
	assert.Equal(rpcRequest.GetMethodName(), "", "rpcRequest error")
	assert.Equal(rpcRequest.GetData(), "", "rpcRequest error")
	assert.Equal(rpcRequest.GetLanguage(), "", "rpcRequest error")
	assert.Equal(rpcRequest.GetVersion(), "", "rpcRequest error")
	//})

}
