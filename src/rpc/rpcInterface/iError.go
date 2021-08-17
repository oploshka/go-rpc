package rpcInterface

type Error interface {
	GetCode() string
	GetMessage() string
	GetData() interface{}

	SetCode(value string) Error
	SetMessage(value string) Error
	SetData(value interface{}) Error
}