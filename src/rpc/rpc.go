package rpc

import (
	"encoding/json"
	"project-my-test/src/rpc/rpcInterface"
	"project-my-test/testHelper/reformHelper"
	"reflect"
)

// получить rpcRequest, вернуть rpcResponse
func (rc *RpcCore) RunMethodByRpcRequest(rReq *request) rpcInterface.Response {
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
func (rc *RpcCore) RpcMethodDataInit(method rpcInterface.Method, jsonMap map[string]json.RawMessage) rpcInterface.Error {

	// get method schema
	RequestSchema := method.GetRequestSchema()

	// schema run
	objValue := reflect.ValueOf(method).Elem()
	dataStructLink := objValue.FieldByName("Data")

	var Reform = reformHelper.GetBaseReform()

	for fieldName := range RequestSchema {

		jsonFieldValue, ok := jsonMap[ RequestSchema[fieldName].Field ]
		if !ok {
			return NewRpcError("VALIDATE_ERROR", "not require field: " + RequestSchema[fieldName].Field, nil)
		}

		res3, err3 := Reform.RunReformItem( RequestSchema[fieldName].Type, jsonFieldValue)
		if err3 != nil {
			//log.Println("ERROR ReformNew.RunReformItem", err3)
			return NewRpcError("VALIDATE_ERROR", err3.Error(), nil)
		}

		strSet := reflect.ValueOf(res3)
		dataStructLink.FieldByName(fieldName).Set(strSet)
	}

	return nil
}


func (rc *RpcCore) TestJsonMethodByRpcRequest(Json string, rReq rpcInterface.Request) rpcInterface.Response {

	// load json
	var jsonMap map[string]json.RawMessage

	jsonByte := []byte(Json)
	err := json.Unmarshal(jsonByte, &jsonMap)
	if err != nil {
		res :=  NewRpcResponse()
		return res
	}

	// # склеиваем в 1
	// # _runMethod
	// # _runMethodProcessing

	// подготавливаем среду для выполнения метода
	methodName := rReq.GetMethodName()
	methodInfo, ok := rc.rpcMethodStore.GetMethodInfo(methodName)
	if !ok {
		rpcResponse := NewRpcResponse()
		rpcResponse.SetError(NewRpcError(
			"ERROR_METHOD_NOT_FOUND",
			"method '" + methodName + "' undefined",
			nil,
		))
		return rpcResponse
	}
	rpcMethod := methodInfo.GetClass()

	// вытаскиваем схему валидации данных requestSchema

	// валидация данных и DTO
	errReform := rc.RpcMethodDataInit(rpcMethod, jsonMap)
	if errReform != nil {
		rpcResponse := NewRpcResponse()
		rpcResponse.SetError(errReform)
		return rpcResponse
	}


	// init
	rpcMethod.SetResponse(NewRpcResponse())
	// выполняем метод
	rpcResponse := rpcMethod.Run()
	// отдаем ответ
	return rpcResponse

}
