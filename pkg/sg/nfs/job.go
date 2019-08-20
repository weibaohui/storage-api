//job任务管理
package nfs

import (
	"encoding/json"
	"errors"
	"fmt"
	"nfs-api/pkg/sg"
	"time"
)

var (
	jobStateStopped = "STOPPED"
	jobStateRunning = "RUNNING"
	jobStateReady   = "READY"
)

type jobResult struct {
	ID               int64        `json:"id"`
	Name             string       `json:"name"`
	Progress         int          `json:"progress"`
	State            string       `json:"state"` //STOPPED RUNNING,READY 运行状态
	ErrorMsg         *sg.ErrorMsg `json:"result"`
	ResultType       string       `json:"result_type"`
	EndTime          int          `json:"end_time"`
	EndTimeForPerf   int          `json:"end_time_for_perf"`
	StartTime        int64        `json:"start_time"`
	StartTimeForPerf int64        `json:"start_time_for_perf"`
}
type jobResultWrapper struct {
	sg.ErrorMsg
	Data *jobResult `json:"result"`
}

type jobID struct {
	JobID    int64  `json:"job_id"`
	JobIDStr string `json:"job_id_str"`
}
type jobIDResult struct {
	sg.ErrorMsg
	Data *jobID `json:"result"`
}

//获取JOB
//POST
//params: {"job_id_str":"1603768355463880"}
//https://192.168.3.60:6080/commands/get_job_by_id.action?cmd_id=0.9346451830352056&user_name=optadmin&uuid=9fdc9c55-cb34-4e40-9da9-ada6d5334a6c
func (r *instance) getJobById(jobID string) (*jobResult, error) {
	url := r.common.Command("/commands/get_job_by_id.action")
	params := make(map[string]string, 0)
	params["params"] = fmt.Sprintf("{\"job_id_str\":\"%s\"}", jobID)
	jsonStr, err := r.common.PostWithLoginSession(url, params)
	if err != nil {
		return nil, err
	}
	wrapper := &jobResultWrapper{}
	err = json.Unmarshal([]byte(jsonStr), wrapper)
	if err != nil {
		return nil, err
	}
	if wrapper.ErrNo != 0 {
		return nil, errors.New(wrapper.ErrorString())
	}
	return wrapper.Data, nil
}

func (r *instance) isJobDone(jobID string) (bool, error) {
	jobResult, err := r.getJobById(jobID)
	if err != nil {
		return false, err
	}
	switch jobResult.State {
	case jobStateReady:
		return true, nil
	case jobStateStopped:
		traceResult := jobResult.ErrorMsg
		if traceResult.ErrNo == 0 {
			return true, nil
		} else {
			return false, errors.New(traceResult.ErrorString())
		}
	case jobStateRunning:
		time.Sleep(time.Millisecond * 500)
		return r.isJobDone(jobID)
	}
	return false, errors.New("未知错误jobResult.State=" + jobResult.State)
}
