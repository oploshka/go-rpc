package rpc

import "project-my-test/src/rpc/rpcInterface"

//
type methodStoreInfo struct {
	name  string
	class rpcInterface.Method
	group string
}

func (rMI *methodStoreInfo) GetClass() rpcInterface.Method {
	return rMI.class
}

func (rMI *methodStoreInfo) GetGroup() string {
	return rMI.group
}

//
type methodStore struct {
	store map[string]methodStoreInfo
}

func (rMS *methodStore) Add(name string, class rpcInterface.Method, group string) {
	rMS.store[name] = methodStoreInfo{name: name, class: class, group: group}
}

func (rMS *methodStore) GetMethodInfo(name string) (methodStoreInfo, bool) {
	val, ok := rMS.store[name]
	return val, ok
}

func NewMethodStore() *methodStore {
	rMS := new(methodStore)
	rMS.store = make(map[string]methodStoreInfo)
	return rMS
}
