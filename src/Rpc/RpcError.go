package Rpc

//
type rpcError struct {
	code    string
	message string
	data    interface{}
}

func (rE *rpcError) GetCode() string {
	return rE.code
}
func (rE *rpcError) GetMessage() string {
	return rE.message
}
func (rE *rpcError) GetData() interface{} {
	return rE.data
}

func (rE *rpcError) SetCode(value string) *rpcError {
	rE.code = value
	return rE
}
func (rE *rpcError) SetMessage(value string) *rpcError {
	rE.message = value
	return rE
}
func (rE *rpcError) SetData(value interface{}) *rpcError {
	rE.data = value
	return rE
}

func NewRpcError() *rpcError {
	// code = ERROR_DEFAULT
	return new(rpcError)
}
