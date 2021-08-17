package RpcOld

import (
	"encoding/json"
	"log"
	"project-my-test/testHelper/reformHelper"
	"reflect"
)

// получить rpcRequest, вернуть rpcResponse
func (rc *RpcCore) RunMethodByRpcRequest(rReq *RpcRequest) *RpcResponse {
	// # склеиваем в 1
	// # _runMethod
	// # _runMethodProcessing

	// подготавливаем среду для выполнения метода
	methodName := rReq.GetMethodName()
	methodInfo, ok := rc.rpcMethodStore.GetMethodInfo(methodName)
	if !ok {
		return NewRpcResponse()
	}
	rpcMethod := methodInfo.GetClass()

	// вытаскиваем схему валидации данных requestSchema

	// валидация данных

	// init
	rpcMethod.SetResponse(NewRpcResponse())
	// выполняем метод
	rpcResponse := rpcMethod.Run()
	// отдаем ответ
	return rpcResponse
}


// RpcMethodDataInit
// это магический метод DTO входных данных - который закидывает данные внутрь метода
// данный метод находится в процессе разработки (нет отлова ошибок и доп проверок)
func (rc *RpcCore) RpcMethodDataInit(method iMethod, jsonMap map[string]json.RawMessage) {

	// get method schema
	RequestSchema := method.GetRequestSchema()

	// schema run
	objValue := reflect.ValueOf(method).Elem()
	dataStructLink := objValue.FieldByName("Data")

	var Reform = reformHelper.GetBaseReform()

	for fieldName := range RequestSchema {

		jsonFieldValue, ok := jsonMap[ RequestSchema[fieldName].Field ]
		log.Print(ok)

		res3, err3 := Reform.RunReformItem( RequestSchema[fieldName].Type, jsonFieldValue)
		if err3 != nil {
			log.Println("ERROR ReformNew.RunReformItem", err3)
		}

		strSet := reflect.ValueOf(res3)
		dataStructLink.FieldByName(fieldName).Set(strSet)
	}
}


func (rc *RpcCore) TestJsonMethodByRpcRequest(Json string, rReq *RpcRequest) *RpcResponse {

	// load json
	var jsonMap map[string]json.RawMessage

	jsonByte := []byte(Json)
	err := json.Unmarshal(jsonByte, &jsonMap)
	if err != nil {
		return NewRpcResponse()
	}

	// # склеиваем в 1
	// # _runMethod
	// # _runMethodProcessing

	// подготавливаем среду для выполнения метода
	methodName := rReq.GetMethodName()
	methodInfo, ok := rc.rpcMethodStore.GetMethodInfo(methodName)
	if !ok {
		return NewRpcResponse()
	}
	rpcMethod := methodInfo.GetClass()

	// вытаскиваем схему валидации данных requestSchema

	// валидация данных и DTO
	rc.RpcMethodDataInit(rpcMethod, jsonMap)

	// init
	rpcMethod.SetResponse(NewRpcResponse())
	// выполняем метод
	rpcResponse := rpcMethod.Run()
	// отдаем ответ
	return rpcResponse

}