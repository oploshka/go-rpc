package rpcInterface

type Response interface {
    GetRpcRequest() Request
    GetData() map[string]interface{}
    GetError() Error
    
    SetData(key string, value interface{}) Response
    
    SetError(error Error) Response
}
