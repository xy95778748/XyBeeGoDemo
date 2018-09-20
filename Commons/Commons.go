package Commons

import "fmt"

/* 相应返回的 Model */
type ResponseModel struct {
	Msg string
	Code int
	Data interface{}
}

func (self *ResponseModel)Success(data interface{}) {

	self.Data = data
	self.Code = 0
	self.Msg = "请求成功"
}

func (self *ResponseModel)Failed(msg string) {

	self.Data = nil
	self.Msg = msg
	self.Code = 1
}

func CheckError(err error, msg string) bool {

	if err == nil {
		return true
	} else {
		fmt.Println( msg, " : ", err)
		return false
	}
}

