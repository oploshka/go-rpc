package RpcOld

//
type RpcRequest struct {
	requestId  string
	methodName string
	data       string
	language   string
	version    string
}

func (rReq *RpcRequest) GetRequestId() string {
	return rReq.requestId
}
func (rReq *RpcRequest) GetMethodName() string {
	return rReq.methodName
}
func (rReq *RpcRequest) GetData() string {
	return rReq.data
}
func (rReq *RpcRequest) GetLanguage() string {
	return rReq.language
}
func (rReq *RpcRequest) GetVersion() string {
	return rReq.version
}

func (rReq *RpcRequest) SetMethodName(methodName string) *RpcRequest {
	rReq.methodName = methodName
	return rReq
}

func NewRpcRequest() *RpcRequest {
	return new(RpcRequest)
}
