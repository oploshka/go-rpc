package rpcInterface

type MethodStoreInfo interface {
	GetClass() Method
	GetGroup() string
}

type MethodStore interface {
	Add(name string, class Method, group string)
	GetMethodInfo(name string) (MethodStoreInfo, bool)
}
