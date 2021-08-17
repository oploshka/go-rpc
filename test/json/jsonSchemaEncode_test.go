package json

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"net/url"
	"reflect"
	"testing"
)


var testJsonString = `{
"say":"Hello",
"sendMsg":{
    "user":"ANisus",
    "msg":"Trying to send a message"
   }
}`
var testJsonArray = "[" + testJsonString + "," + testJsonString + "]"


type RequestSchema struct{
	Name  string `my-require:"true"`
	Email string `mytag:"MyEmail"`
}


func ShowStructure(s interface{}) {
	a := reflect.ValueOf(s)
	numfield := reflect.ValueOf(s).Elem().NumField()
	if a.Kind() != reflect.Ptr {
		log.Fatal("wrong type struct")
	}
	for x := 0; x < numfield; x++ {
		fmt.Printf("Name field: `%s`  Type: `%s`\n", reflect.TypeOf(s).Elem().Field(x).Name,
			reflect.ValueOf(s).Elem().Field(x).Type())
	}
}

func TestJson_schema(t *testing.T) {
	assert := assert.New(t)

	rS := RequestSchema{Name:"test", Email: "asdasd@asdasd.sd"}


	val := reflect.ValueOf(rS)
	for i := 0; i < val.Type().NumField(); i++ {
		as := val.Type().Field(i)
		fmt.Println(as)
		// as.Name  - имя файла

		fmt.Println(val.Type().Field(i).Tag.Get("json"))
	}

	//ShowStructure(rS)
	//fieldName := "Name"

	assert.NotNil(rS)

	//field, found := t.FieldByName(fieldName)
	//if !found {
	//	continue
	//}
	//fmt.Printf("\nField: User.%s\n", fieldName)
	//fmt.Printf("\tWhole tag value : %q\n", field.Tag)
	//fmt.Printf("\tValue of 'mytag': %q\n", field.Tag.Get("mytag"))
}

func TestJson_correctWork(t *testing.T) {
	assert := assert.New(t)

	jsonByte := []byte(testJsonString)

	var objmap map[string]json.RawMessage
	err := json.Unmarshal(jsonByte, &objmap)

	assert.Nil(err)

	for k := range objmap {
		log.Println(k)
	}
	log.Print(objmap)
}


func TestJson_url_params_to_json(t *testing.T) {
	assert := assert.New(t)

	u, err := url.Parse("https://example.org/?method=MethodTestData2&a=1&a=2&b=&=3&&&&")
	if err != nil {
		log.Fatal(err)
	}
	q := u.Query()

	assert.NotNil(q)

	jsonByte, _ := json.Marshal(q)
	jsonString := string(jsonByte)
	assert.NotNil(jsonString)

	keys, ok := q["method"]
	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'method' is missing")
	}
}
