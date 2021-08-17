package RpcStructure

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
  Specification         string                            `json:"specification"`
  SpecificationVersion  string                            `json:"specificationVersion"`
  Version               string                            `json:"version"`
  Language              string                            `json:"language"`
  Request               *MultipartJsonRpcRequest_request  `json:"response"`
}
type MultipartJsonRpcRequest_request struct {
  Id       interface{}              `json:"id"`
  Name     string                   `json:"name"`
  Data     map[string]interface{}   `json:"data"`
}


/*
class MultipartJsonRpcRequest implements \Oploshka\RpcInterface\RpcStructure {
  
  public function decode($arr) {
  
    if(
      !is_array($arr)
      || !isset( $arr['request'] )
      || !is_array($arr['request'])
      || !isset($arr['request']['name']) || !is_string($arr['request']['name'])
      || !isset($arr['request']['data']) || !is_array($arr['request']['data'])
    ){
      throw new \Oploshka\RpcException\ReformException('ERROR_REQUEST_STRUCTURE_DECODE');
    }
    
    return new \Oploshka\Rpc\RpcRequest([
      'requestId'   => $arr['request']['id'] ?? null,
      'methodName'  => $arr['request']['name'],
      'data'        => $arr['request']['data'],
      // TODO
      // 'language'    => $arr['request']['data'],
      // 'version'     => $arr['request']['data'],
    ]);
  
  }
  
  // @param \Oploshka\Rpc\RpcRequest $RpcMethodInfo
  public function encode($RpcMethodInfo){

    return [
      'specification'         => 'multipart-json-rpc',
      'specificationVersion'  => '0.1.0',
      'version'               => $RpcMethodInfo->getVersion(),
      'language'              => $RpcMethodInfo->getLanguage(),
  
      'request'               => [
        'id'        => $RpcMethodInfo->getRequestId(),
        'name'      => $RpcMethodInfo->getMethodName(),
        'data'      => $RpcMethodInfo->getData(),
      ],
    ];
  }
  
}
 */
