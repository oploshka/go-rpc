package RpcOld

//
type RpcResponse struct {
	rpcRequest *RpcRequest
	data       map[string]interface{}
	error      *rpcError
}

func (rRes *RpcResponse) GetRpcRequest() *RpcRequest {
	return rRes.rpcRequest
}
func (rRes *RpcResponse) GetData() map[string]interface{} {
	return rRes.data
}
func (rRes *RpcResponse) GetError() *rpcError {
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
func (rRes *RpcResponse) SetData(key string, value interface{}) *RpcResponse {
	rRes.data[key] = value
	return rRes
}

func (rRes *RpcResponse) SetError(error *rpcError) *RpcResponse {
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

func NewRpcResponse() *RpcResponse {

	rRes := new(RpcResponse)
	rRes.rpcRequest = NewRpcRequest()
	rRes.data = make(map[string]interface{})
	rRes.error = NewRpcError()
	return rRes

	// $this->RpcRequest = $arr['RpcRequest'] ?? new RpcRequest(['methodName' => 'UNDEFINED']);
	// $this->Error      = new Error();
}
