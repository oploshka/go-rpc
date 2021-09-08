package reformHelper

import (
    "encoding/json"
    "log"
)


func GetOldReform() map[string]func(jsonFieldValue json.RawMessage) (interface{}, error) {
    
    var Reform = make(map[string]func(jsonFieldValue json.RawMessage) (interface{}, error))
    Reform["STRING"] = func(jsonFieldValue json.RawMessage) (interface{}, error) {
        var str string
        err := json.Unmarshal(jsonFieldValue, &str)
        if err != nil {
            log.Print(err)
            return nil, err
        }
        return &str, nil
    }
    Reform["INT"] = func(jsonFieldValue json.RawMessage) (interface{}, error) {
        var str int
        err := json.Unmarshal(jsonFieldValue, &str)
        if err != nil {
            log.Print(err)
            return nil, err
        }
        return &str, nil
    }
    
    return Reform
}
