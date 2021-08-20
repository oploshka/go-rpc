package reform

import (
	"encoding/json"
	"fmt"
	"github.com/oleiade/reflections"
	"github.com/stretchr/testify/assert"
	"log"
	"project-my-test/src/rpc/rpcInterface"
	"project-my-test/src/rpc/rpcStruct"
	"project-my-test/testHelper/methodHelper/methodHelper_group2"
	"project-my-test/testHelper/reformHelper"
	"reflect"
	"testing"
)

var testJsonString22 = `{
"name":"Hello",
"email": "test@mail.ru"
}`

func TestReform_112a(t *testing.T) {
	assert := assert.New(t)

	// load json
	jsonByte := []byte(testJsonString22)
	var jsonMap map[string]json.RawMessage
	err := json.Unmarshal(jsonByte, &jsonMap)
	assert.Nil(err)

	// get method schema
	rpcMethod := methodHelper_group2.ReturnRequestSchemaData{}
	RequestSchema := rpcMethod.GetRequestSchema()

	// schema run
	dataStructLink := &rpcMethod.Data

	keys := make([]rpcStruct.ReformSchema, 0, len(RequestSchema))
	for fieldName := range RequestSchema {
        keys = append(keys, RequestSchema[fieldName])
        
        jsonFieldValue, ok := jsonMap[RequestSchema[fieldName].Field]
        assert.Equal(ok, true)
        
        var str string
        err := json.Unmarshal(jsonFieldValue, &str)
        assert.Nil(err)
        
        err2 := reflections.SetField(dataStructLink, fieldName, &str)
        assert.Nil(err2)
    }

	assert.Equal(*rpcMethod.Data.Name, "Hello")
	assert.Equal(*rpcMethod.Data.Email, "test@mail.ru")
	assert.NotNil(keys)
}

func RpcMethodDataInit(method rpcInterface.Method, jsonMap map[string]json.RawMessage) {
	// get method schema
	RequestSchema := method.GetRequestSchema()

	// schema run
	//dataStructLink, _ := reflections.GetField(method, "Data")
	objValue := reflect.ValueOf(method).Elem()
	fmt.Println(reflect.ValueOf(objValue))

	dataStructLink := objValue.FieldByName("Data")

	//dd := &dataStructLink
	//structValue := reflect.ValueOf(dd).Elem()
	//structFieldValue := structValue.FieldByName("Name")
	//log.Print(structFieldValue)

	var ReformNew = reformHelper.GetBaseReform()

	for fieldName := range RequestSchema {
        
        jsonFieldValue, ok := jsonMap[RequestSchema[fieldName].Field]
        log.Print(ok)
        
        var str string
        err := json.Unmarshal(jsonFieldValue, &str)
        log.Print(err)
        
        // err2 := reflections.SetField(dataStructLink, fieldName, &str)
        // log.Print(err2)
        // iStr := str
        
        // var Reform = reformHelper.GetOldReform()
        // var funcReform = Reform[ RequestSchema[fieldName].Type ]
        // res, _ := funcReform(jsonFieldValue)
        // strSet := reflect.ValueOf(&str)
        
        res3, err3 := ReformNew.RunReformItem(RequestSchema[fieldName].Type, jsonFieldValue)
        if err3 != nil {
            log.Println("ERROR ReformNew.RunReformItem", err3)
        }
        
        strSet := reflect.ValueOf(res3)
        
        dataStructLink.FieldByName(fieldName).Set(strSet)
    }
}

func TestReform_1312x(t *testing.T) {
	assert := assert.New(t)

	// load json
	jsonByte := []byte(testJsonString22)
	var jsonMap map[string]json.RawMessage
	err := json.Unmarshal(jsonByte, &jsonMap)
	assert.Nil(err)

	// get method schema
	rpcMethod := methodHelper_group2.ReturnRequestSchemaData{}

	//rpcMethod.GetRequestSchema()
	RpcMethodDataInit(&rpcMethod, jsonMap)

	assert.Equal(*rpcMethod.Data.Name, "Hello")
	assert.Equal(*rpcMethod.Data.Email, "test@mail.ru")
}
