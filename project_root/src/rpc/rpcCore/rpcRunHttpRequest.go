package rpcCore

import (
    "encoding/json"
    "net/http"
    "project-my-test/src/rpc"
    "project-my-test/src/rpc/plugin/rpcRequestLoad"
    "project-my-test/src/rpc/plugin/rpcStructureRequest"
    "project-my-test/src/rpc/plugin/rpcStructureResponse"
    "project-my-test/src/rpc/rpcInterface"
)

func (rc *RpcCore) RunByHttpRequest(r *http.Request) rpcInterface.Response {
    
    rpcRequestLoader := rpcRequestLoad.NewPostMultipartFormDataField("data")
    jsonString, errLoad := rpcRequestLoader.Load(r)
    
    if errLoad != nil {
        rpcResponse := rpc.NewRpcResponse(rpc.NewRpcRequest("", "", nil, "", ""))
        rpcResponse.SetError(errLoad)
        return rpcResponse
    }
    jsonByte := []byte(*jsonString)
    
    // декодируем в структуру
    reqStruct := new(rpcStructureRequest.MultipartJsonRpcRequest)
    if err := json.Unmarshal(jsonByte, reqStruct); err != nil {
        rpcResponse := rpc.NewRpcResponse(rpc.NewRpcRequest("", "", nil, "", ""))
        rpcResponse.SetError(rpc.NewRpcError("STRUCTURE_ERROR", err.Error(), nil))
        
        return rpcResponse
    }
    
    // получаем объект запроса
    rpcRequest := reqStruct.ConvertToRpcRequest()
    
    //
    rpcResponse := rc.RunMethodByRpcRequest(rpcRequest)
    
    return rpcResponse
}


func returnHttpJson(w http.ResponseWriter, rResp rpcInterface.Response) {
    
    responseStruct := rpcStructureResponse.MultipartJsonRpcResponseEncode(rResp)
    
    js, err := json.Marshal(responseStruct)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    w.Write(js)
}


func (rc *RpcCore) RunHttp(w http.ResponseWriter, r *http.Request) {
    rpcResponse := rc.RunByHttpRequest(r)
    returnHttpJson(w, rpcResponse)
}