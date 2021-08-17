package reformItem

import (
	"encoding/json"
)

type BooleanReformItem struct {}

func (ri *BooleanReformItem) Validate(jsonFieldValue json.RawMessage) (interface{}, error) {
	var str bool
	err := json.Unmarshal(jsonFieldValue, &str)
	if err != nil {
		return nil, err
	}
	return &str, nil
}

// const FILTER_MIN = 'min';
// const FILTER_MAX = 'max';
// const FILTER = [
//   self::FILTER_MIN => 0,
//   self::FILTER_MAX => 1000000000,
// ];


//func (ri *intReformItem) Filter(jsonFieldValue json.RawMessage) (interface{}, error) {
//	return "asda", nil
//}