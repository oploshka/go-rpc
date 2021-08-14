package RpcStructure

import "project-my-test/src/Rpc"

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

func MultipartJsonRpcResponseEncode(rResp *Rpc.RpcResponse) *MultipartJsonRpcResponse {

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

///**
// * @param array $arr
// * @return \Oploshka\Rpc\RpcResponse
// * @throws \Exception
// */
//func MultipartJsonRpcResponse_decode(value interface{}) *Rpc.RpcResponse {
//    //if(
//    //  !is_array($arr)
//    //  || !isset( $arr['request'] )
//    //  || !is_array($arr['request'])
//    //  || !isset($arr['request']['name']) || !is_string($arr['request']['name'])
//    //  || !isset($arr['request']['data']) || !is_array($arr['request']['data'])
//    //){
//    //  throw new \Oploshka\RpcException\ReformException('ERROR_RESPONSE_STRUCTURE_DECODE');
//    //}
//    //
//    //return new \Oploshka\Rpc\RpcResponse([
//    //  'requestId'   => $arr['request']['id'] ?? null,
//    //  'methodName'  => $arr['request']['name'],
//    //  'data'        => $arr['request']['data'],
//    //  // TODO
//    //  // 'language'    => $arr['request']['data'],
//    //  // 'version'     => $arr['request']['data'],
//    //]);
//
//  return Rpc.NewRpcResponse()
//}
