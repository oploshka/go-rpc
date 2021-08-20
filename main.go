package main

import (
    "encoding/json"
    "io/ioutil"
    "log"
    "net/http"
    "project-my-test/example/rpcApp"
    "project-my-test/src/rpc"
    "project-my-test/src/rpc/plugin/rpcRequestLoad"
    "project-my-test/src/rpc/plugin/rpcStructure"
)

func main() {
    
    rpcClient := rpcApp.CreateTestRpc()
    
    //
    //
    //
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        
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
            
            responseStruct := rpcStructure.MultipartJsonRpcResponseEncode(rpcResponse)
            
            js, err := json.Marshal(responseStruct)
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
            
            w.Header().Set("Content-Type", "application/json")
            w.Write(js)
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
    })
    
    // api page
    http.HandleFunc("/api-form", func(w http.ResponseWriter, r *http.Request) {
        dat, err := ioutil.ReadFile("./static/api.html")
        if err != nil {
            w.Write([]byte("Не удалось считать файл"))
            return
        }
        w.Header().Set("Content-Type", "text/html")
        w.Write(dat)
    })
    // browser favicon fix
    http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte(""))
    })
    
    log.Fatal(http.ListenAndServe(":8080", nil))
}
