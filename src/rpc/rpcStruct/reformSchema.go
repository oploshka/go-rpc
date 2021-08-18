package rpcStruct


type ReformSchema struct {
	Type		string
	Field 		string
	Req   		bool
	Default 	func()interface{}
	Validate 	func()error
}