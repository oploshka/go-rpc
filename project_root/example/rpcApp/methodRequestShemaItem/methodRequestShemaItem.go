package methodRequestShemaItem

import (
    "project-my-test/example/rpcApp/reform"
    "project-my-test/src/rpc/rpcStruct"
)

// group 1:

func GROUP1_TEST_USET_STRING() rpcStruct.ReformSchema {
    return rpcStruct.ReformSchema{
        Type:  reform.REFORM_STRING,
        Field: "full_name",
    }
}
func GROUP1_TEST_USET_INT() rpcStruct.ReformSchema {
    return rpcStruct.ReformSchema{
        Type:  reform.REFORM_INTEGER,
        Field: "number",
    }
}
func GROUP1_TEST_USET_BOOL() rpcStruct.ReformSchema {
    return rpcStruct.ReformSchema{
        Type:  reform.REFORM_BOOLEAN,
        Field: "bool",
    }
}

// group 2:

func GROUP2_NAME() rpcStruct.ReformSchema {
    return rpcStruct.ReformSchema{
        Type:  reform.REFORM_STRING,
        Field: "name", // TODO: FieldLoad and FieldReturn
        Default: func() interface{} {
            return "UserNameDefault"
        },
    }
}
func GROUP2_EMAIL() rpcStruct.ReformSchema {
    return rpcStruct.ReformSchema{
        Type:  reform.REFORM_STRING,
        Field: "email", // TODO: FieldLoad and FieldReturn
    }
}