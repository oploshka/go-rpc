package RpcOld

//
type rpcMethodInfo struct {
	name  string
	class iMethod
	group string
}

func (rMI *rpcMethodInfo) GetClass() iMethod {
	return rMI.class
}

func (rMI *rpcMethodInfo) GetGroup() string {
	return rMI.group
}

//
type rpcMethodStore struct {
	store map[string]rpcMethodInfo
}

func (rMS *rpcMethodStore) Add(name string, class iMethod, group string) {
	rMS.store[name] = rpcMethodInfo{name: name, class: class, group: group}
}

func (rMS *rpcMethodStore) GetMethodInfo(name string) (rpcMethodInfo, bool) {
	val, ok := rMS.store[name]
	return val, ok
}

func NewRpcMethodStore() *rpcMethodStore {
	rMS := new(rpcMethodStore)
	rMS.store = make(map[string]rpcMethodInfo)
	return rMS
}
