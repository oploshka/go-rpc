package rpcInterface

import "encoding/json"

type Request interface {
	GetRequestId() string
	GetMethodName() string
	GetData() map[string]json.RawMessage // Это временное решение, пока не станет понимания как можно лучше
	GetLanguage() string
	GetVersion() string

	SetMethodName(methodName string)
}
