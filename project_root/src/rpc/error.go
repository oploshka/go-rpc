package rpc

import "project-my-test/src/rpc/rpcInterface"

// TODO rename
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

func NewRpcError(code string, message string, data interface{}) rpcInterface.Error {
    // code = ERROR_DEFAULT
    return &error{
        code:    code,
        message: message,
        data:    data,
    }
}
