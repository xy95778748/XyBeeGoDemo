package Commons

import "fmt"

/* 相应返回的 Model */
type ResponseModel struct {
	Msg string
	Code int
	Data interface{}
}

func MakeSuccessResponse (data interface{}) ResponseModel {

	return ResponseModel{"请求成功", 0, data}
}

func MakeFailedResponse (data interface{}, msg string) ResponseModel {

	return ResponseModel{msg, 1, data}
}

func CheckError(err error, msg string) bool {

	if err == nil {
		return true
	} else {
		fmt.Println( msg, " : ", err)
		return false
	}
}

