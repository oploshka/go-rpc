package reformItem

import (
	"encoding/json"
)

type StringReformItem struct{}

func (ri *StringReformItem) Validate(jsonFieldValue json.RawMessage) (interface{}, error) {
	var str string
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
