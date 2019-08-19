//job任务管理
package sg

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

var (
	jobStateStopped = "STOPPED"
	jobStateRunning = "RUNNING"
)

type TraceResult struct {
	TraceID        string `json:"trace_id"`
	DetailErrMsg   string `json:"detail_err_msg"`
	TimeStamp      int    `json:"time_stamp"`
	ErrMsg         string `json:"err_msg"`
	TimeZoneOffset int    `json:"time_zone_offset"`
	Sync           bool   `json:"sync"`
	ErrNo          int    `json:"err_no"`
}
type JobResult struct {
	ID               int64        `json:"id"`
	Name             string       `json:"name"`
	Progress         int          `json:"progress"`
	State            string       `json:"state"` //STOPPED RUNNING 运行状态
	TraceResult      *TraceResult `json:"result"`
	ResultType       string       `json:"result_type"`
	EndTime          int          `json:"end_time"`
	EndTimeForPerf   int          `json:"end_time_for_perf"`
	StartTime        int64        `json:"start_time"`
	StartTimeForPerf int64        `json:"start_time_for_perf"`
}
type JobResultWrapper struct {
	ErrorMsg
	Data           *JobResult `json:"result"`
	Sync           bool       `json:"sync"`
	TimeStamp      int64      `json:"time_stamp"`
	TimeZoneOffset int        `json:"time_zone_offset"`
	TraceID        string     `json:"trace_id"`
}

type JobID struct {
	JobID    int64  `json:"job_id"`
	JobIDStr string `json:"job_id_str"`
}
type JobIDResult struct {
	ErrorMsg
	Data           *JobID `json:"result"`
	Sync           bool   `json:"sync"`
	TimeStamp      int64  `json:"time_stamp"`
	TimeZoneOffset int    `json:"time_zone_offset"`
}

//获取JOB
//POST
//params: {"job_id_str":"1603768355463880"}
//https://192.168.3.60:6080/commands/get_job_by_id.action?cmd_id=0.9346451830352056&user_name=optadmin&uuid=9fdc9c55-cb34-4e40-9da9-ada6d5334a6c
func (r *Robot) GetJobById(jobID string) (*JobResult, error) {
	url := r.fullURL("/commands/get_job_by_id.action?user_name=" + r.Username + "&uuid=" + r.uuid)
	params := make(map[string]string, 0)
	params["params"] = fmt.Sprintf("{\"job_id_str\":\"%s\"}", jobID)
	jsonStr, err := r.PostWithLoginSession(url, params)
	if err != nil {
		return nil, err
	}
	wrapper := &JobResultWrapper{}
	err = json.Unmarshal([]byte(jsonStr), wrapper)
	if err != nil {
		return nil, err
	}
	if wrapper.ErrNo != 0 {
		return nil, errors.New(wrapper.ErrorString())
	}
	return wrapper.Data, nil
}

func (r *Robot) IsJobDone(jobID string) (bool, error) {
	jobResult, err := r.GetJobById(jobID)
	if err != nil {
		return false, err
	}
	if jobResult.State == jobStateRunning {
		time.Sleep(time.Millisecond * 500)
		return r.IsJobDone(jobID)
	}
	if jobResult.State == jobStateStopped {
		traceResult := jobResult.TraceResult
		if traceResult.ErrNo == 0 {
			return true, nil
		} else {
			return false, errors.New(traceResult.ErrMsg + traceResult.DetailErrMsg)
		}
	}
	return false, errors.New("未知错误")
}
