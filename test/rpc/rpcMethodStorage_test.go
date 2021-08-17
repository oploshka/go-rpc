package rpc

import (
	"github.com/stretchr/testify/assert"
	"project-my-test/src/Rpc"
	"project-my-test/testHelper/methodHelper/methodHelper_group2"
	"testing"
)

func TestRpcMethodStorage_correctWork(t *testing.T) {
	assert := assert.New(t)

	//t.Run("test1", func(t *testing.T) {
	//
	methodStore := Rpc.NewRpcMethodStore()

	methodStore.Add("test1", new(methodHelper_group2.ReturnRequestSchemaData), "group1")
	// methodStore.Add("test2", "class2", "group2")
	// methodStore.Add("test3", "class3", "group3")

	methodInfo, ok := methodStore.GetMethodInfo("test2")
	methodClass := methodInfo.GetClass()
	methodClass.Run()

	assert.Equal(ok, true, "MethodInfo test2 - empty")
	//assert.Equal(methodInfo.GetClass(), "class2", "MethodInfo test2 - not correct class")
	assert.Equal(methodInfo.GetGroup(), "group2", "MethodInfo test2 - not correct group")
	//})

}
