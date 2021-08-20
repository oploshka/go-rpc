package jsonStringHelper

func GetLargeNestingJsonString() string {
	return `{
"say":"Hello",
"sendMsg":{
    "user":"ANisus",
    "msg":"Trying to send a message"
	"sendMsg":{
		"user":"ANisus",
		"msg":"Trying to send a message"
	}
   }
}`
}