package RpcOld

type iMethod interface {
	// GetDescription() string
	GetRequestSchema() map[string]RpcReformSchema
	// GetResponseSchema() string
	Run() *RpcResponse
	SetResponse(rRes *RpcResponse)
}

type IMethod interface {
	// GetDescription() string
	GetRequestSchema() map[string]RpcReformSchema
	// GetResponseSchema() string
	Run() *RpcResponse
	SetResponse(rRes *RpcResponse)
}