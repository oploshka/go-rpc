package rpcStructureRequest

import (
    "encoding/json"
    "project-my-test/src/rpc"
    "project-my-test/src/rpc/rpcInterface"
)

/*
{
  "specification": "multipart-json-rpc",
  "specificationVersion" : "0.1.0",

  "version": "1",
  "language": "en",

  "request" : {
    "id"   : "9423234",
    "name" : "MethodTest1",
    "data" : { "userId" : 1 }
  }
}
*/

type MultipartJsonRpcRequest_request struct {
    Id   interface{} `json:"id"`
    Name string      `json:"name"`
    
    // TODO: fix
    // Data interface{} `json:"data"`
    Data map[string]json.RawMessage `json:"data"`
}

type MultipartJsonRpcRequest struct {
    Specification        string                           `json:"specification"`
    SpecificationVersion string                           `json:"specificationVersion"`
    Version              string                           `json:"version"`
    Language             string                           `json:"language"`
    Request              *MultipartJsonRpcRequest_request `json:"request"`
}

// TODO: add return error
func (obj *MultipartJsonRpcRequest) ConvertToRpcRequest() rpcInterface.Request {
    return rpc.NewRpcRequest(
        obj.Request.Id, // requestId string,
        obj.Request.Name, // methodName string,
        obj.Request.Data, // data map[string]json.RawMessage,
        obj.Language, // language string,
        obj.Version, // version string,
    )
}




//
// MultipartJsonRpcRequestEncode - это вспомогательный блок (нужен для удобства написания тестов)
//
func MultipartJsonRpcRequestEncode(rResp rpcInterface.Response) *MultipartJsonRpcRequest {
    
    resp := MultipartJsonRpcRequest{
        Specification:        "multipart-json-rpc",
        SpecificationVersion: "0.1.0",
        Version:              "1.0.0",
        Language:             "en",
        //
        Request: &MultipartJsonRpcRequest_request{
            Id:   rResp.GetRpcRequest().GetRequestId(),
            Name: rResp.GetRpcRequest().GetMethodName(),
            Data: rResp.GetRpcRequest().GetData(),
        },
    }
    
    return &resp
}
