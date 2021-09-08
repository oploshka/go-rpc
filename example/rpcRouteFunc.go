package example

import (
    "encoding/json"
    "log"
    "net/http"
    "project-my-test/example/rpcApp"
    "project-my-test/src/rpc"
    "project-my-test/src/rpc/plugin/rpcRequestLoad"
    "project-my-test/src/rpc/plugin/rpcStructure"
    "project-my-test/src/rpc/rpcInterface"
)


func returnHttpJson(w http.ResponseWriter, rResp rpcInterface.Response) {
    
    responseStruct := rpcStructure.MultipartJsonRpcResponseEncode(rResp)
    
    js, err := json.Marshal(responseStruct)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    w.Write(js)
}

func TestHttpRpc(w http.ResponseWriter, r *http.Request) {
    // TODO: delete rpc Init (need for test)
    rpcClient := rpcApp.CreateTestRpc()
    
    
    // temp code for method test by get params
    getMap := r.URL.Query()
    // jsonByte, _ := json.Marshal(getMap)
    // jsonString := string(jsonByte)
    
    methodName := "MethodTestData1"
    keys, ok := getMap["method"]
    if !ok || len(keys[0]) < 1 {
        log.Println("Url Param 'method' is missing")
    } else {
        methodName = keys[0]
    }
    
    rpcRequestLoader := rpcRequestLoad.NewPostMultipartFormDataField("data")
    jsonString, errLoad := rpcRequestLoader.Load(r)
    
    if errLoad != nil {
        rpcResponse := rpc.NewRpcResponse(rpc.NewRpcRequest("", "", nil, "", ""))
        rpcResponse.SetError(errLoad)
        
        returnHttpJson(w, rpcResponse)
        return
    }
    
    
    // load json
    var jsonMap map[string]json.RawMessage
    
    jsonByte := []byte(*jsonString)
    err := json.Unmarshal(jsonByte, &jsonMap)
    if err != nil {
        // res := NewRpcResponse(rpcRequest)
        // res.SetError(NewRpcError(
        //     "ERROR_JSON_DECODE",
        //     err.Error(),
        //     err,
        // ))
        // return res
        w.Write([]byte("JSON ERROR 500"))
        return
    }
    
    //
    rpcRequestTestMethod := rpc.NewRpcRequest("", methodName, jsonMap, "", "")
    //
    rpcResponse := rpcClient.RunMethodByRpcRequest(rpcRequestTestMethod)
    
    // TODO: временное решение
    // log.Println("=================================")
    responseStruct := rpcStructure.MultipartJsonRpcResponseEncode(rpcResponse)
    
    js, err := json.Marshal(responseStruct)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    w.Write(js)
}