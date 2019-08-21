package sg

import (
	"errors"
	"fmt"
)

type ErrorMsg struct {
	DetailErrMsg   interface{} `json:"detail_err_msg"`
	ErrMsg         interface{} `json:"err_msg"`
	ErrNo          int         `json:"err_no"`
	TimeStamp      int         `json:"timeStamp"`
	Sync           bool        `json:"sync"`
	TimeZoneOffset int         `json:"time_zone_offset"`
	TraceID        string      `json:"trace_id"`
}

func (s *ErrorMsg) ErrorString() string {
	if s.ErrNo != 0 {
		return ""
	}
	return fmt.Sprintf("%d,%s,%s,%v", s.ErrNo, s.ErrMsg, s.DetailErrMsg, s.TimeStamp)
}
func (s *ErrorMsg) Error() error {
	return errors.New(s.ErrorString())
}
