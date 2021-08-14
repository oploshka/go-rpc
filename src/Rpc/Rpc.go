package Rpc

// получить rpcRequest, вернуть rpcResponse
func (rc *RpcCore) RunMethodByRpcRequest(rReq *RpcRequest) *RpcResponse {
	// # склеиваем в 1
	// # _runMethod
	// # _runMethodProcessing

	// подготавливаем среду для выполнения метода
	methodName := rReq.GetMethodName()
	methodInfo, ok := rc.rpcMethodStore.GetMethodInfo(methodName)
	if !ok {
		return NewRpcResponse()
	}
	rpcMethod := methodInfo.GetClass()

	// вытаскиваем схему валидации данных requestSchema

	// валидация данных

	// init
	rpcMethod.SetResponse(NewRpcResponse())
	// выполняем метод
	rpcResponse := rpcMethod.Run()
	// отдаем ответ
	return rpcResponse
}
