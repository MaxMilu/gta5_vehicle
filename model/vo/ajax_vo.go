package vo

import (
	"my_qor_test/config/consts"
	"my_qor_test/utils/json_utils"
)

var AjaxSystemError = &AjaxResult{
	Status:  consts.FAILURE,
	Message: consts.SYSTEM_ERROR,
	Result:  consts.SYSTEM_ERROR,
}

type AjaxResult struct {
	Status  string
	Message string
	Result  interface{}
}

func (result AjaxResult) Marshal() []byte {
	if resultData, err := json_utils.Encode(result); err != nil {
		return []byte{}
	} else {
		return resultData
	}
}

// 返回带自定义信息的错误Ajax
func AjaxFailureMessage(message string) *AjaxResult {
	return &AjaxResult{
		Status:  consts.FAILURE,
		Message: message,
		Result:  consts.SYSTEM_ERROR,
	}
}
