package reformInterface

import "encoding/json"

type Reform interface {
    AddReformItem(name string, item ReformItem)
    RunReformItem(name string, jsonFieldValue json.RawMessage) (interface{}, error)
}
