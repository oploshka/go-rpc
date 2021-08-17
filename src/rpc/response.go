package rpc

import "project-my-test/src/rpc/rpcInterface"

//
type response struct {
	rpcRequest rpcInterface.Request
	data       map[string]interface{}
	error      rpcInterface.Error
}

func (rRes *response) GetRpcRequest() rpcInterface.Request {
	return rRes.rpcRequest
}
func (rRes *response) GetData() map[string]interface{} {
	return rRes.data
}
func (rRes *response) GetError() rpcInterface.Error {
	return rRes.error
}

///**  @return string */
//public function getErrorCode(){
//  return $this->Error->getCode();
//}
///**  @return string */
//public function getErrorMessage(){
//  return $this->Error->getMessage();
//}
///**  @return array */
//public function getErrorData(){
//  return $this->Error->getData();
//}

// setters
func (rRes *response) SetData(key string, value interface{}) rpcInterface.Response {
	rRes.data[key] = value
	return rRes
}

func (rRes *response) SetError(error rpcInterface.Error) rpcInterface.Response {
	rRes.error = error
	return rRes
}

//public function setErrorCode($code){
//$this->Error->setCode($code);
//return $this;
//}
//public function setErrorMessage($message){
//$this->Error->setMessage($message);
//return $this;
//}
//public function setErrorData($data){
//$this->Error->setData($data);
//return $this;
//}
//public function error($error = null) {
//$error && $this->error = $error;
//throw new \Oploshka\RpcException\MethodEndException('');
//}

func NewRpcResponse() rpcInterface.Response {

	rRes := new(response)
	rRes.rpcRequest = NewRpcRequest()
	rRes.data = make(map[string]interface{})
	rRes.error = NewRpcError("ERROR_DEFAULT", "", nil)
	return rRes

	// $this->RpcRequest = $arr['RpcRequest'] ?? new RpcRequest(['methodName' => 'UNDEFINED']);
	// $this->Error      = new Error();
}
