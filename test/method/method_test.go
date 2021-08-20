package rpc

import (
    "github.com/fatih/structs"
    "github.com/oleiade/reflections"
    "github.com/stretchr/testify/assert"
    "project-my-test/src/rpc/rpcInterface"
    "project-my-test/example/rpcApp/method/methodGroup2"
    "testing"
)

func TestMethod_correctWork(t *testing.T) {
    assert := assert.New(t)
    
    methodRpc := methodGroup2.MethodMyTest{}
    MethodValidate(assert, &methodRpc)
}

func MethodValidate(assert *assert.Assertions, methodRpc rpcInterface.Method) {
    methodData, _ := reflections.GetField(methodRpc, "Data")
    methodRequestSchema := methodRpc.GetRequestSchema()
    
    fields := structs.Fields(methodData)
    fieldsMap := make(map[string]bool)
    for _, field := range fields {
        fieldsMap[field.Name()] = true
        _, ok := methodRequestSchema[field.Name()]
        assert.Equal(ok, true, "[MethodName]: Data field '"+field.Name()+"' is not add to RequestSchema")
    }
    
    for key := range methodRequestSchema {
        _, ok := fieldsMap[key]
        assert.Equal(ok, true, "[MethodName]: RequestSchema field '"+key+"' is not add to Data")
        
        // TODO: RequestSchema.Type есть в Reform
    }
    
}
