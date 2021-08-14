package Rpc

type iMethod interface {
	// GetDescription() string
	// GetRequestSchema() string
	// GetResponseSchema() string
	Run() *RpcResponse
	SetResponse(rRes *RpcResponse)
}
