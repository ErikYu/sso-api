package e

import "fmt"

type BaseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var (
	ErrorFetchParam = BaseError{
		Code:    40001,
		Message: "解析参数失败",
	}
)

func init() {
	fmt.Println(ErrorFetchParam)
}
