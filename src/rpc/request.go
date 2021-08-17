package rpc

//
type request struct {
	requestId  string
	methodName string
	data       string
	language   string
	version    string
}

func (rReq *request) GetRequestId() string {
	return rReq.requestId
}
func (rReq *request) GetMethodName() string {
	return rReq.methodName
}
func (rReq *request) GetData() string {
	return rReq.data
}
func (rReq *request) GetLanguage() string {
	return rReq.language
}
func (rReq *request) GetVersion() string {
	return rReq.version
}

func (rReq *request) SetMethodName(methodName string) {
	rReq.methodName = methodName
}

func NewRpcRequest() *request {
	return new(request)
}
