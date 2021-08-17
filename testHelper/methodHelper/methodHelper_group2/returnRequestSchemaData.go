package methodHelper_group2

import (
	"project-my-test/src/Rpc"
	"project-my-test/testHelper/reformHelper"
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
		Type: reformHelper.REFORM_STRING,
		Field: "name", // TODO: FieldLoad and FieldReturn
	}
	rs["Email"] = Rpc.RpcReformSchema{
		Type: reformHelper.REFORM_STRING,
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


