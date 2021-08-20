package rpcStructure

import "project-my-test/src/rpc/rpcInterface"

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

type MultipartJsonRpcRequest struct {
    Specification        string                           `json:"specification"`
    SpecificationVersion string                           `json:"specificationVersion"`
    Version              string                           `json:"version"`
    Language             string                           `json:"language"`
    Request              *MultipartJsonRpcRequest_request `json:"response"`
}
type MultipartJsonRpcRequest_request struct {
    Id   interface{} `json:"id"`
    Name string      `json:"name"`
    Data interface{} `json:"data"`
}

func MultipartJsonRpcRequestDecode() {
    // if(
    //   !is_array($arr)
    //   || !isset( $arr['request'] )
    //   || !is_array($arr['request'])
    //   || !isset($arr['request']['name']) || !is_string($arr['request']['name'])
    //   || !isset($arr['request']['data']) || !is_array($arr['request']['data'])
    // ){
    //   throw new \Oploshka\RpcException\ReformException('ERROR_REQUEST_STRUCTURE_DECODE');
    // }
    //
    // return new \Oploshka\Rpc\RpcRequest([
    //   'requestId'   => $arr['request']['id'] ?? null,
    //   'methodName'  => $arr['request']['name'],
    //   'data'        => $arr['request']['data'],
    //   // TODO
    //   // 'language'    => $arr['request']['data'],
    //   // 'version'     => $arr['request']['data'],
    // ]);
    // return &resp
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
