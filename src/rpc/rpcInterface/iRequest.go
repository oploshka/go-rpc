package rpcInterface

type Request interface {
	GetRequestId() string
	GetMethodName() string
	GetData() string
	GetLanguage() string
	GetVersion() string

	SetMethodName(methodName string)
}
