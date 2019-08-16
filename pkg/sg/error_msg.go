package sg

import "fmt"

type ErrorMsg struct {
	DetailErrMsg interface{} `json:"detail_err_msg"`
	ErrMsg       interface{} `json:"err_msg"`
	ErrNo        int         `json:"err_no"`
	TimeStamp    int         `json:"timeStamp"`
}

func (s *ErrorMsg) ErrorString() string {
	return fmt.Sprintf("%d,%s,%s,%v", s.ErrNo, s.ErrMsg, s.DetailErrMsg, s.TimeStamp)
}
