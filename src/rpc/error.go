package rpc

import "project-my-test/src/rpc/rpcInterface"

//
type error struct {
	code    string
	message string
	data    interface{}
}

func (rE *error) GetCode() string {
	return rE.code
}
func (rE *error) GetMessage() string {
	return rE.message
}
func (rE *error) GetData() interface{} {
	return rE.data
}

func (rE *error) SetCode(value string) rpcInterface.Error {
	rE.code = value
	return rE
}
func (rE *error) SetMessage(value string) rpcInterface.Error {
	rE.message = value
	return rE
}
func (rE *error) SetData(value interface{}) rpcInterface.Error {
	rE.data = value
	return rE
}

func NewRpcError() rpcInterface.Error {
	// code = ERROR_DEFAULT
	return new(error)
}
