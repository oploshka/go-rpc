package rpcInterface

import "net/http"



// получаем пользовательские данные в строку
type RpcRequestLoad interface {
    Load(req *http.Request) (*string, Error)
}
