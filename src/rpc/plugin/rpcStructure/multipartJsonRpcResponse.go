package rpcStructure

import (
	"project-my-test/src/rpc"
	"project-my-test/src/rpc/rpcInterface"
)

/*
{
  "specification": "multipart-json-rpc",
  "specificationVersion": "0.1.0",
  "version": "1.0.0",
  "language": "en",
  "response": {
    "requestId": "9423234",
    "error": {
      "code": "ERROR_NO_METHOD_INFO",
      "message": "",
      "data": []
    },
    "data": {}
  }
}
*/

type MultipartJsonRpcResponse struct {
	Specification        string                             `json:"specification"`
	SpecificationVersion string                             `json:"specificationVersion"`
	Version              string                             `json:"version"`
	Language             string                             `json:"language"`
	Response             *MultipartJsonRpcResponse_response `json:"response"`
}
type MultipartJsonRpcResponse_response struct {
	RequestId interface{}                              `json:"requestId"`
	Error     *MultipartJsonRpcResponse_response_error `json:"error"`
	Data      map[string]interface{}                   `json:"data"`
}
type MultipartJsonRpcResponse_response_error struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func MultipartJsonRpcResponseEncode(rResp rpcInterface.Response) *MultipartJsonRpcResponse {

	resp := MultipartJsonRpcResponse{
		Specification:        "multipart-json-rpc",
		SpecificationVersion: "0.1.0",
		Version:              "1.0.0",
		Language:             "en",
		//
		Response: &MultipartJsonRpcResponse_response{
			RequestId: rResp.GetRpcRequest().GetRequestId(),
			Error: &MultipartJsonRpcResponse_response_error{
				Code:    rResp.GetError().GetCode(),
				Message: rResp.GetError().GetMessage(),
				Data:    rResp.GetError().GetData(),
			},
			Data: rResp.GetData(),
		},
	}

	return &resp
}

//
// MultipartJsonRpcResponseDecode - это вспомогательный блок (нужен для удобства написания тестов)
//
func MultipartJsonRpcResponseDecode(mjr *MultipartJsonRpcResponse) rpcInterface.Response{
	//rpcError    := rpc.NewRpcError(mjr.Response.Error.Code, mjr.Response.Error.Message, mjr.Response.Error.Data )
	//rpcRequest  := rpc.NewRpcRequest()
	rpcResponse := rpc.NewRpcResponse()

	return rpcResponse
}