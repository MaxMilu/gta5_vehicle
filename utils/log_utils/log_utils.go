package log_utils

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
	"my_qor_test/config/consts"
	"my_qor_test/libraries/error_ext"
	"runtime/debug"
)

// 打印异常或错误日志
func PrintFault(r interface{}, err error) {
	if r != nil {
		log.Printf("%v\n%v\n", r, string(debug.Stack()))
	}
	if err != nil {
		log.Printf(consts.ERROR_STACK_FORMATTER, err)
	}
}

// 打印异常或错误日志并返回友好提示错误
func PrintFaultFriendly(r interface{}, err error) error {
	if r != nil || err != nil {
		PrintFault(r, err)
		return errors.New(consts.SYSTEM_ERROR)
	}
	return nil
}

// 获取异常或错误信息
func GetFaultMessage(r interface{}, err error) string {
	if r != nil {
		return fmt.Sprintf("%v\n%v\n", r, string(debug.Stack()))
	}
	if err != nil {
		return fmt.Sprintf(consts.ERROR_STACK_FORMATTER, err)
	}
	return consts.EMPTY
}

// 获取异常或错误信息并返回友好提示
func GetFaultMessageFriendly(r interface{}, err error) (string, error) {
	if faultMessage := GetFaultMessage(r, err); faultMessage != consts.EMPTY {
		return faultMessage, errors.New(consts.SYSTEM_ERROR)
	}
	return consts.EMPTY, nil
}

// 打印异常日志
func PrintException(r interface{}) {
	if r != nil {
		log.Printf("%v\n%v\n", r, string(debug.Stack()))
	}
}

// 打印异常日志并返回友好提示错误
func PrintExceptionFriendly(r interface{}) error {
	if r != nil {
		PrintException(r)
		return errors.New(consts.SYSTEM_ERROR)
	}
	return nil
}

// 获取异常信息
func GetExceptionMessage(r interface{}) string {
	if r != nil {
		return fmt.Sprintf("%v\n%v\n", r, string(debug.Stack()))
	}
	return consts.EMPTY
}

// 获取异常信息并返回友好提示
func GetExceptionMessageFriendly(r interface{}) (string, error) {
	if exceptionMessage := GetExceptionMessage(r); GetExceptionMessage(r) != consts.EMPTY {
		return exceptionMessage, errors.New(consts.SYSTEM_ERROR)
	}
	return consts.EMPTY, nil
}

// 打印错误日志
func PrintError(err error) {
	if err != nil {
		log.Printf(consts.ERROR_STACK_FORMATTER, err)
	}
}

// 打印错误日志并返回友好提示
func PrintErrorFriendly(err error) error {
	if err != nil {
		PrintError(err)
		return errors.New(consts.SYSTEM_ERROR)
	}
	return nil
}

// 打印错误日志堆栈（仅限立即打印错误情况）
func PrintErrorStackTrace(err error) {
	if err != nil {
		log.Printf(consts.ERROR_STACK_FORMATTER, errors.WithStack(err))
	}
}

// 打印错误日志堆栈并返回友好提示（仅限立即打印错误并返回的情况）
func PrintErrorStackTraceFriendly(err error) error {
	if err != nil {
		if error_ext.CheckError(err) {
			PrintErrorStackTrace(err)
			return err
		} else {
			PrintErrorStackTrace(err)
			return errors.New(consts.SYSTEM_ERROR)
		}
	}
	return nil
}

// 获取错误信息
func GetErrorMessage(err error) string {
	if err != nil {
		return fmt.Sprintf(consts.ERROR_STACK_FORMATTER, err)
	}
	return consts.EMPTY
}

// 获取错误信息并返回友好提示
func GetErrorMessageFriendly(err error) (string, error) {
	if errorMessage := GetErrorMessage(err); errorMessage != consts.EMPTY {
		return GetErrorMessage(err), errors.New(consts.SYSTEM_ERROR)
	}
	return consts.EMPTY, nil
}

func Println(logMessage interface{}) {
	log.Println(logMessage)
}
