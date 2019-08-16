//job任务管理
package sg

import "fmt"

//https://192.168.3.60:6080/commands/get_job_by_id.action?cmd_id=0.9346451830352056&user_name=optadmin&uuid=9fdc9c55-cb34-4e40-9da9-ada6d5334a6c
func (r *Robot) Job(uuid, jobID string) {
	url := r.fullURL("/commands/get_job_by_id.action?user_name=" + r.Username + "&uuid=" + uuid)
	fmt.Println(url)
}
