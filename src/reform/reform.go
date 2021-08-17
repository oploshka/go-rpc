package reform

import (
	"encoding/json"
	"errors"
	"project-my-test/src/reform/reformInterface"
)

type reform struct {
	store map[string]reformInterface.ReformItem
}

func (r *reform) AddReformItem(name string, item reformInterface.ReformItem ) {
	r.store[name] = item
}

func (r *reform) RunReformItem(name string, jsonFieldValue json.RawMessage) (interface{}, error) {
	val, ok := r.store[name]
	if !ok {
		return nil, errors.New("[SERVER_ERROR] Not correct validate type: " + name)
	}
	res, err := val.Validate(jsonFieldValue)
	return res, err
}


func NewReform() *reform {
	r := new(reform)
	r.store = make(map[string]reformInterface.ReformItem)
	return r
}
