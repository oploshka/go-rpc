package rpcRequestLoad

import (
    "net/http"
    "project-my-test/src/rpc"
    "project-my-test/src/rpc/rpcInterface"
)

//
type postMultipartFormDataField struct {
    field     string
    maxMemory int64
}

func (pmf *postMultipartFormDataField) Load(req *http.Request) (*string, rpcInterface.Error) {
    
    if req.Method != "POST" {
        return nil, rpc.NewRpcError("ERROR_REQUEST_LOAD", "#10 request method type is not 'POST'", nil)
    }
    
    // https://golangbyexample.com/http-mutipart-form-body-golang/
    err := req.ParseMultipartForm(pmf.maxMemory) // maxMemory 32MB
    if err != nil {
        return nil, rpc.NewRpcError("ERROR_REQUEST_LOAD", "#20 multipart/form-data load error", err)
    }
    
    keys, ok := req.MultipartForm.Value[pmf.field]
    if !ok {
        return nil, rpc.NewRpcError("ERROR_REQUEST_LOAD", "#30 load '"+pmf.field+"' field error", nil)
    }
    if len(keys[0]) < 1 {
        return nil, rpc.NewRpcError("ERROR_REQUEST_LOAD", "#40 load '"+pmf.field+"' field len error", err)
    }
    
    return &keys[0], nil
}

func NewPostMultipartFormDataField(field string) *postMultipartFormDataField {
    pmf := new(postMultipartFormDataField)
    pmf.field = field
    pmf.maxMemory = 32 << 20
    return pmf
}
