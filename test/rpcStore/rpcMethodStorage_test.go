package rpcStore

import (
	"github.com/stretchr/testify/assert"
	"project-my-test/src/rpc"
	"project-my-test/testHelper/methodHelper/methodHelper_group1"
	"testing"
)

func TestRpcMethodStorage_correctWork(t *testing.T) {
	assert := assert.New(t)

	//t.Run("test1", func(t *testing.T) {
	//
	methodStore := rpc.NewMethodStore()

	methodStore.Add("test2", new(methodHelper_group1.MethodTestData3), "group2")
	// methodStore.Add("test2", "class2", "group2")
	// methodStore.Add("test3", "class3", "group3")

	methodInfo, ok := methodStore.GetMethodInfo("test2")
	assert.Equal(ok, true, "MethodInfo test2 - empty")

	methodClass := methodInfo.GetClass()
	rpcResp := rpc.NewRpcResponse()
	methodClass.SetResponse(rpcResp)
	assert.NotEmpty(methodClass)

	//assert.Equal(methodInfo.GetClass(), "class2", "MethodInfo test2 - not correct class")
	assert.Equal(methodInfo.GetGroup(), "group2", "MethodInfo test2 - not correct group")
	//})

}
