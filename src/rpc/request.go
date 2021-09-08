package rpc

import "encoding/json"

//
type request struct {
    requestId  interface{}
    methodName string
    // dataType => json
    data     map[string]json.RawMessage // Это временное решение, пока не станет понимания как можно лучше
    language string
    version  string
}

func (rReq *request) GetRequestId() interface{} {
	return rReq.requestId
}
func (rReq *request) GetMethodName() string {
	return rReq.methodName
}
func (rReq *request) GetData() map[string]json.RawMessage {
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

func NewRpcRequest(
    requestId interface{},
    methodName string,
    data map[string]json.RawMessage,
    language string,
    version string,
) *request {
    
    rpcRequest := request{
        requestId:  requestId,
        methodName: methodName,
        data:       data,
        language:   language,
        version:    version,
    }
    return &rpcRequest
}
