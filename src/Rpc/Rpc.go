package Rpc

import (
	"encoding/json"
	"log"
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

	var Reform = make(map[string]func(jsonFieldValue json.RawMessage) (interface{}, error))
	Reform["STRING"] = func(jsonFieldValue json.RawMessage) (interface{}, error) {
		var str string
		err := json.Unmarshal(jsonFieldValue, &str)
		if err != nil {
			log.Print(err)
			return nil, err
		}
		return &str, nil
	}
	Reform["INT"] = func(jsonFieldValue json.RawMessage) (interface{}, error) {
		var str int
		err := json.Unmarshal(jsonFieldValue, &str)
		if err != nil {
			log.Print(err)
			return nil, err
		}
		return &str, nil
	}

	for fieldName := range RequestSchema {

		jsonFieldValue, ok := jsonMap[ RequestSchema[fieldName].Field ]
		log.Print(ok)

		//var str string
		//err := json.Unmarshal(jsonFieldValue, &str)
		//log.Print(err)
		var funcReform = Reform[ RequestSchema[fieldName].Type ]
		res, _ := funcReform(jsonFieldValue)

		strSet := reflect.ValueOf(res)
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