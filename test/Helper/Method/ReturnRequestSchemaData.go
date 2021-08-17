package Method

import (
	"project-my-test/src/Rpc"
)


type ReturnRequestSchemaData struct {
	Rpc.RpcMethod
	Data struct{
		Name  *string
		Email *string
	}
}

func (r *ReturnRequestSchemaData) GetRequestSchema() map[string]Rpc.RpcReformSchema {
	rs := make(map[string]Rpc.RpcReformSchema)

	rs["Name"] = Rpc.RpcReformSchema{
		Type: "STRING",
		Field: "name", // TODO: FieldLoad and FieldReturn
	}
	rs["Email"] = Rpc.RpcReformSchema{
		Type: "STRING",
		Field: "email", // TODO: FieldLoad and FieldReturn
	}
	//RequestSchema struct{ Name, Email string `mytag:"MyEmail"` }
	return rs
}



func (r *ReturnRequestSchemaData) Run() *Rpc.RpcResponse {
	r.Response.SetData("test::string", "string")
	r.Response.SetData("test::int", 20)
	r.Response.GetError().SetCode("ERROR")
	return r.Response
}


