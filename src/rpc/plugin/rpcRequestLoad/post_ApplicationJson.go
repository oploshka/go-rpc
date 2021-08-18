package rpcRequestLoad

import (
  "net/http"
  "project-my-test/src/rpc"
  "project-my-test/src/rpc/rpcInterface"
)

//
type postApplicationJson struct {
  field string
  maxMemory int64
}

func (paj *postApplicationJson) Load(req *http.Request) (*string, rpcInterface.Error) {

  if req.Method != "POST" {
    return nil, rpc.NewRpcError("ERROR_REQUEST_LOAD", "#10 request method type is not 'POST'", nil)
  }

  // https://www.alexedwards.net/blog/how-to-properly-parse-a-json-request-body

  // req.Body - request body to string
  // https://stackoverflow.com/questions/38673673/access-http-response-as-string-in-go

  str := "{ test json}"
  return &str, nil
}



func NewPostApplicationJson(field string) *postMultipartFormDataField {
  paj := new(postMultipartFormDataField)
  paj.field = field
  paj.maxMemory = 32 << 20
  return paj
}