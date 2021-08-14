package Rpc

type RpcCore struct {
	// # di                    string // di container
	// # reform                string // validate and convert data  [ReformDebug]
	rpcMethodStore *rpcMethodStore // method store

	//
	// # rpcRequestLoad        string // Post_MultipartFormData_Field - вытаскиваем данные из http запроса
	// # rpcRequestFormatter   string // json_decode и json_encode (тут json->object)
	// # rpcRequestStructure   string // вытаскиваем из структуры (что прислал пользователь) - данные
	// # rpcResponseFormatter  string // json_decode и json_encode (тут object->json)
	// # rpcResponseStructure  string // создаем структуру для отправки данных
}

func (rC *RpcCore) GetRpcMethodStore() *rpcMethodStore {
	return rC.rpcMethodStore
}

func NewRpcCore() *RpcCore {
	rC := new(RpcCore)
	rC.rpcMethodStore = NewRpcMethodStore()
	return rC
}

// # // getters
// # public function getReform()               { return $this->Reform;               }
// # public function getRpcMethodStorage()     { return $this->RpcMethodStorage;     }
// # //
// # public function getRpcRequestLoad()       { return $this->RpcRequestLoad;       }
// # public function getRpcRequestFormatter()  { return $this->RpcRequestFormatter;  }
// # public function getRpcRequestStructure()  { return $this->RpcRequestStructure;  }
// # public function getRpcResponseFormatter() { return $this->RpcResponseFormatter; }
// # public function getRpcResponseStructure() { return $this->RpcResponseStructure; }
// #
// # // setters TODO: fix
// # public function setReform($obj)               { return $this->Reform           = $obj; }
// # public function setRpcMethodStorage($obj)     { return $this->RpcMethodStorage = $obj; }
// # //
// # public function setRpcRequestLoad($obj)       { $this->RpcRequestLoad       = $obj; }
// # public function setRpcRequestFormatter($obj)  { $this->RpcRequestFormatter  = $obj; }
// # public function setRpcRequestStructure($obj)  { $this->RpcRequestStructure  = $obj; }
// # public function setRpcResponseFormatter($obj) { $this->RpcResponseFormatter = $obj; }
// # public function setRpcResponseStructure($obj) { $this->RpcResponseStructure = $obj; }

// TODO: add settings Go ???
// TODO: add settings http headers ???
