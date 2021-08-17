package reformHelper

import (
	"project-my-test/src/reform"
	"project-my-test/src/reform/reformInterface"
	"project-my-test/src/reform/reformItem"
)

const REFORM_INTEGER 	= "INTEGER"
const REFORM_STRING 	= "STRING"

func GetBaseReform() reformInterface.Reform {
	reform := reform.NewReform()
	reform.AddReformItem(REFORM_INTEGER, &reformItem.IntReformItem{})
	reform.AddReformItem(REFORM_STRING , &reformItem.StringReformItem{})

	return reform
}
