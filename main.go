package main

import (
	"encoding/json"
	"log"
	"net/http"
	"project-my-test/src/rpc"
	"project-my-test/src/rpc/plugin/rpcStructure"
	"project-my-test/testHelper/rpcHelper"
)

func main() {

	rpcClient := rpcHelper.CreateTestRpc()

	//
	//
	//
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//
		rpcRequestTestMethod := rpc.NewRpcRequest()
		rpcRequestTestMethod.SetMethodName("MethodTestData1")
		//
		rpcResponse := rpcClient.TestJsonMethodByRpcRequest("{\"full_name\": \"Andrey\", \"number\": 17 }", rpcRequestTestMethod)

		// TODO: временное решение
		log.Println("=================================")
		responseStruct := rpcStructure.MultipartJsonRpcResponseEncode(rpcResponse)

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
		rpcRequestTestMethod := rpc.NewRpcRequest()
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
		rpcRequestTestMethod := rpc.NewRpcRequest()
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
