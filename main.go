package main

import (
	"encoding/json"
	"log"
	"net/http"
	"project-my-test/src/Rpc"
	"project-my-test/src/Rpc/Plugin/RpcStructure"
	"project-my-test/test/Helper/Method"
	"project-my-test/test/Helper/Method/MethodTestData"
)

func main() {

	rpcClient := Rpc.NewRpcCore()
	rpcMethodStore := rpcClient.GetRpcMethodStore()
	//
	rpcMethodStore.Add("testMethod", new(Method.ReturnRequestSchemaData), "")
	rpcMethodStore.Add("MethodTestData1", new(MethodTestData.MethodTestData1), "TestData")
	rpcMethodStore.Add("MethodTestData2", new(MethodTestData.MethodTestData2), "TestData")
	rpcMethodStore.Add("MethodTestData3", new(MethodTestData.MethodTestData3), "TestData")

	//
	//
	//
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//
		rpcRequestTestMethod := Rpc.NewRpcRequest()
		rpcRequestTestMethod.SetMethodName("MethodTestData1")
		//
		rpcResponse := rpcClient.TestJsonMethodByRpcRequest("{\"full_name\": \"Andrey\", \"number\": 17 }", rpcRequestTestMethod)

		// TODO: временное решение
		log.Println("=================================")
		responseStruct := RpcStructure.MultipartJsonRpcResponseEncode(rpcResponse)

		js, err := json.Marshal(responseStruct)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	})
	//
	http.HandleFunc("/1", func(w http.ResponseWriter, r *http.Request) {
		//
		rpcRequestTestMethod := Rpc.NewRpcRequest()
		rpcRequestTestMethod.SetMethodName("MethodTestData1")
		//
		rpcResponse := rpcClient.RunMethodByRpcRequest(rpcRequestTestMethod)

		// TODO: временное решение
		log.Println("=================================")
		log.Println(rpcResponse)

		js, err := json.Marshal(rpcResponse.GetData())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	})
	//
	http.HandleFunc("/2", func(w http.ResponseWriter, r *http.Request) {
		//
		rpcRequestTestMethod := Rpc.NewRpcRequest()
		rpcRequestTestMethod.SetMethodName("MethodTestData2")
		//
		rpcResponse := rpcClient.RunMethodByRpcRequest(rpcRequestTestMethod)

		// TODO: временное решение
		log.Println("=================================")
		log.Println(rpcResponse)

		js, err := json.Marshal(rpcResponse.GetData())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
