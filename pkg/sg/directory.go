package sg

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type DetailFiles struct {
	AccessTime      int64  `json:"access_time"`
	CreateTime      int64  `json:"create_time"`
	ModifyTime      int64  `json:"modify_time"`
	Name            string `json:"name"`
	OwnerGroupName  string `json:"owner_group_name"`
	OwnerUserName   string `json:"owner_user_name"`
	Path            string `json:"path"`       // ParaStor300S:/test/123
	PosixPath       string `json:"posix_path"` // /test/123
	PosixPermission string `json:"posix_permission"`
	Size            int    `json:"size"`
	Type            string `json:"type"` //DIR FILE
}
type DirectoryList struct {
	ErrorMsg
	Data struct {
		DetailFiles []*DetailFiles `json:"detail_files"`
		Total       int            `json:"total"`
	} `json:"result"`
}

// 文件创建，需要按目录层级结构，逐级创建
//https://192.168.3.60:6080/commands/create_file.action?user_name=optadmin&uuid=9fdc9c55-cb34-4e40-9da9-ada6d5334a6c
//rand:
//params: {"path":"ParaStor300S:/ddd","posix_permission":"rwxr-xr-x","auth_provider_id":"0","owner_user_id":0,"owner_group_id":0,"owner_user_name":"root","owner_group_name":"root"}

func (r *Robot) CreateDirectory(path string) (bool, error) {
	config := fmt.Sprintf(`{
	"path":"%s:%s",
	"posix_permission":"rwxrwxrwx",
	"auth_provider_id":"0",
	"owner_user_id":0,
	"owner_group_id":0,
	"owner_user_name":"root",
	"owner_group_name":"root"
	}`, r.storeName, path)
	return r.createDirectory(config)
}

//指定目录权限
//rwxrwxrwx
//rwxr-xr-x
//用户    读取 写入 执行
//用户组  读取 写入 执行
//其他    读取 写入 执行
func (r *Robot) CreateDirectoryWithPermission(path, permission string) (bool, error) {
	config := fmt.Sprintf(`{
	"path":"%s:%s",
	"posix_permission":"%s",
	"auth_provider_id":"0",
	"owner_user_id":0,
	"owner_group_id":0,
	"owner_user_name":"root",
	"owner_group_name":"root"
	}`, r.storeName, path, permission)
	return r.createDirectory(config)
}
func (r *Robot) createDirectory(config string) (bool, error) {
	url := r.fullURL("/commands/create_file.action?user_name=" + r.Username + "&uuid=" + r.uuid)
	params := make(map[string]string)
	params["params"] = config
	str, err := r.PostWithLoginSession(url, params)
	if err != nil {
		return false, err
	}
	result := &ErrorMsg{}
	err = json.Unmarshal([]byte(str), result)
	if err != nil {
		return false, err
	}
	if result.ErrNo != 0 {
		return false, errors.New(result.ErrorString())
	}

	return true, nil
}

//非空目录不能删除，需要逐级删除
//POST
//https://192.168.3.60:6080/commands/delete_file.action?user_name=optadmin&uuid=9fdc9c55-cb34-4e40-9da9-ada6d5334a6c
//rand:
//params: {"path":"ParaStor300S:/test"}
func (r *Robot) DeleteDirectory(path string) (bool, error) {
	url := r.fullURL("/commands/delete_file.action?user_name=" + r.Username + "&uuid=" + r.uuid)
	params := make(map[string]string)
	params["params"] = fmt.Sprintf(`{
	"path":"%s:%s",
	}`, r.storeName, path)

	str, err := r.PostWithLoginSession(url, params)
	if err != nil {
		return false, err
	}
	result := &ErrorMsg{}
	err = json.Unmarshal([]byte(str), result)
	if err != nil {
		return false, err
	}
	if result.ErrNo != 0 {
		return false, errors.New(result.ErrorString())
	}

	return true, nil
}

//非空目录不能删除，需要逐级删除
//POST
//https://192.168.3.60:6080/commands/get_file_list.action?cmd_id=0.7323753691986996&user_name=optadmin&uuid=9fdc9c55-cb34-4e40-9da9-ada6d5334a6c
func (r *Robot) listDirectory(config string) ([]*DetailFiles, error) {
	url := r.fullURL("/commands/get_file_list.action?user_name=" + r.Username + "&uuid=" + r.uuid)
	params := make(map[string]string)
	params["params"] = config

	str, err := r.PostWithLoginSession(url, params)
	if err != nil {
		return nil, err
	}
	result := &DirectoryList{}
	err = json.Unmarshal([]byte(str), result)
	if err != nil {
		return nil, err
	}
	if result.ErrNo != 0 {
		return nil, errors.New(result.ErrorString())
	}
	for _, v := range result.Data.DetailFiles {
		//去除存储服务器名称
		v.PosixPath = strings.TrimPrefix(v.Path, r.storeName+":")
	}
	return result.Data.DetailFiles, nil
}

//列表显示目录
func (r *Robot) ListDirectory(path string) ([]*DetailFiles, error) {

	config := fmt.Sprintf(`{
	"limit":1000000,
	"start":0,
	"sort":"",
	"path":"%s:%s",	
	"display_details":true,
	"type":"DIR",
	"searches":[{"searchKey":"name","searchValue":""}]
	}`, r.storeName, path)

	return r.listDirectory(config)
}

//列表显示目录及文件
func (r *Robot) ListDirectoryWithFiles(path string) ([]*DetailFiles, error) {
	config := fmt.Sprintf(`{
	"limit":1000000,
	"start":0,
	"sort":"",
	"path":"%s:%s",	
	"display_details":true,
	"searches":[{"searchKey":"name","searchValue":""}]
	}`, r.storeName, path)
	return r.listDirectory(config)

}
