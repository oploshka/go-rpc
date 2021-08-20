package reformInterface

import "encoding/json"

type ReformItem interface {
	Validate(jsonFieldValue json.RawMessage) (interface{}, error)
}
