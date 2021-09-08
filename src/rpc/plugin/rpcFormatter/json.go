package rpcFormatter


//
type jsonFormatter struct {

}

func (jf *jsonFormatter) Decode(str string) interface{} {
    
    return nil
}

func (jf *jsonFormatter) Encode(obj interface{}) string {
    return "{}"
}


func NewJsonFormatter() *jsonFormatter {
    jf := new(jsonFormatter)
    return jf
}
