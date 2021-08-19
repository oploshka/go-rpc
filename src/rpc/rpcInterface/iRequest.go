package rpcInterface

type Request interface {
	GetRequestId() string
	GetMethodName() string
	GetData() map[string]interface{}
	GetLanguage() string
	GetVersion() string

	SetMethodName(methodName string)
}
