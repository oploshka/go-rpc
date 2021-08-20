package rpc

import (
    "encoding/json"
    "project-my-test/src/rpc/rpcInterface"
    "project-my-test/example/rpcApp/reform"
    "reflect"
)

// получить rpcRequest, вернуть rpcResponse
func (rc *RpcCore) RunMethodByRpcRequest(rpcRequest rpcInterface.Request) rpcInterface.Response {
    // # склеиваем в 1
    // # _runMethod
    // # _runMethodProcessing
    
    // подготавливаем среду для выполнения метода
    methodName := rpcRequest.GetMethodName()
    methodInfo, ok := rc.rpcMethodStore.GetMethodInfo(methodName)
    if !ok {
        rpcResponse := NewRpcResponse(rpcRequest)
        rpcResponse.SetError(NewRpcError(
            "ERROR_METHOD_NOT_FOUND",
            "method '"+methodName+"' undefined",
            nil,
        ))
        return rpcResponse
    }
    rpcMethod := methodInfo.GetClass()
    
    // вытаскиваем схему валидации данных requestSchema
    
    requestData := rpcRequest.GetData()
    // requestDataJsonRaw := make(map[string]json.RawMessage)
    //
    // for key, value := range requestData {
    //	//strKey := fmt.Sprintf("%v", key)
    //	//strValue := fmt.Sprintf("%v", value)
    //	requestDataJsonRaw[key] = value.(json.RawMessage)
    // }
    requestDataJsonRaw := requestData
    
    // валидация данных и DTO
    errReform := rc.RpcMethodDataInit(rpcMethod, requestDataJsonRaw)
    if errReform != nil {
        rpcResponse := NewRpcResponse(rpcRequest)
        rpcResponse.SetError(errReform)
        return rpcResponse
    }
    
    // init
    rpcMethod.SetResponse(NewRpcResponse(rpcRequest))
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
    
    var Reform = reform.GetBaseReform()
    
    for fieldName := range RequestSchema {
        
        jsonFieldValue, ok := jsonMap[RequestSchema[fieldName].Field]
        if !ok {
            return NewRpcError("VALIDATE_ERROR", "not require field: "+RequestSchema[fieldName].Field, nil)
        }
        
        res3, err3 := Reform.RunReformItem(RequestSchema[fieldName].Type, jsonFieldValue)
        if err3 != nil {
            // log.Println("ERROR ReformNew.RunReformItem", err3)
            return NewRpcError("VALIDATE_ERROR", err3.Error(), nil)
        }
        
        strSet := reflect.ValueOf(res3)
        dataStructLink.FieldByName(fieldName).Set(strSet)
    }
    
    return nil
}

func (rc *RpcCore) TestJsonMethodByRpcRequest(Json string, rpcRequest rpcInterface.Request) rpcInterface.Response {
    
    // load json
    var jsonMap map[string]json.RawMessage
    
    jsonByte := []byte(Json)
    err := json.Unmarshal(jsonByte, &jsonMap)
    if err != nil {
        res := NewRpcResponse(rpcRequest)
        res.SetError(NewRpcError(
            "ERROR_JSON_DECODE",
            err.Error(),
            err,
        ))
        return res
    }
    
    // # склеиваем в 1
    // # _runMethod
    // # _runMethodProcessing
    
    // подготавливаем среду для выполнения метода
    methodName := rpcRequest.GetMethodName()
    methodInfo, ok := rc.rpcMethodStore.GetMethodInfo(methodName)
    if !ok {
        rpcResponse := NewRpcResponse(rpcRequest)
        rpcResponse.SetError(NewRpcError(
            "ERROR_METHOD_NOT_FOUND",
            "method '"+methodName+"' undefined",
            nil,
        ))
        return rpcResponse
    }
    rpcMethod := methodInfo.GetClass()
    
    // вытаскиваем схему валидации данных requestSchema
    
    // валидация данных и DTO
    errReform := rc.RpcMethodDataInit(rpcMethod, jsonMap)
    if errReform != nil {
        rpcResponse := NewRpcResponse(rpcRequest)
        rpcResponse.SetError(errReform)
        return rpcResponse
    }
    
    // init
    rpcMethod.SetResponse(NewRpcResponse(rpcRequest))
    // выполняем метод
    rpcResponse := rpcMethod.Run()
    // отдаем ответ
    return rpcResponse
    
}
