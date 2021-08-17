package Json

import (
	"encoding/json"
	"fmt"
	"github.com/oleiade/reflections"
	"github.com/stretchr/testify/assert"
	"log"
	"project-my-test/src/Rpc"
	"project-my-test/test/Helper/Method"
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
	rpcMethod := Method.ReturnRequestSchemaData{}
	RequestSchema := rpcMethod.GetRequestSchema()

	// schema run
	dataStructLink := &rpcMethod.Data

	keys := make([]Rpc.RpcReformSchema, 0, len(RequestSchema))
	for fieldName := range RequestSchema {
		keys = append(keys, RequestSchema[fieldName])

		jsonFieldValue, ok := jsonMap[ RequestSchema[fieldName].Field ]
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


func RpcMethodDataInit(method Rpc.IMethod, jsonMap map[string]json.RawMessage) {
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
	for fieldName := range RequestSchema {

		jsonFieldValue, ok := jsonMap[ RequestSchema[fieldName].Field ]
		log.Print(ok)

		var str string
		err := json.Unmarshal(jsonFieldValue, &str)
		log.Print(err)

		//err2 := reflections.SetField(dataStructLink, fieldName, &str)
		//log.Print(err2)
		//iStr := str
		strSet := reflect.ValueOf(&str)
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
	rpcMethod := Method.ReturnRequestSchemaData{}

	//rpcMethod.GetRequestSchema()
	RpcMethodDataInit(&rpcMethod, jsonMap)

	assert.Equal(*rpcMethod.Data.Name, "Hello")
	assert.Equal(*rpcMethod.Data.Email, "test@mail.ru")
}
