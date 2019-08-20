package nfs

import (
	"encoding/json"
	"fmt"
	"storage-api/pkg/api"
	"storage-api/pkg/sg"
	"strings"
)

type DirectoryList struct {
	sg.ErrorMsg
	Data struct {
		DetailFiles []*api.DetailFiles `json:"detail_files"`
		Total       int                `json:"total"`
	} `json:"result"`
}

// 文件创建，需要按目录层级结构，逐级创建
func (i *instance) CreateDirectory(path string) (ok bool, err error) {
	config := fmt.Sprintf(`{
	"path":"%s:%s",
	"posix_permission":"rwxrwxrwx",
	"auth_provider_id":"0",
	"owner_user_id":0,
	"owner_group_id":0,
	"owner_user_name":"root",
	"owner_group_name":"root"
	}`, i.common.StoreName, path)
	return i.createDirectory(config)
}

//指定目录权限
//rwxrwxrwx
//rwxr-xr-x
//用户    读取 写入 执行
//用户组  读取 写入 执行
//其他    读取 写入 执行
func (i *instance) CreateDirectoryWithPermission(path, permission string) (bool, error) {
	config := fmt.Sprintf(`{
	"path":"%s:%s",
	"posix_permission":"%s",
	"auth_provider_id":"0",
	"owner_user_id":0,
	"owner_group_id":0,
	"owner_user_name":"root",
	"owner_group_name":"root"
	}`, i.common.StoreName, path, permission)
	return i.createDirectory(config)
}
func (i *instance) createDirectory(config string) (ok bool, err error) {
	url := i.common.Command("/commands/create_file.action")
	params := make(map[string]string)
	params["params"] = config
	str, err := i.common.PostWithLoginSession(url, params)
	if err != nil {
		return false, err
	}
	result := &sg.ErrorMsg{}
	err = json.Unmarshal([]byte(str), result)
	if err != nil {
		return false, err
	}
	if result.ErrNo != 0 {
		return false, result.Error()
	}

	return true, nil
}

//非空目录不能删除，需要逐级删除
func (i *instance) DeleteDirectory(path string) (ok bool, err error) {
	url := i.common.Command("/commands/delete_file.action")
	params := make(map[string]string)
	params["params"] = fmt.Sprintf(`{
	"path":"%s:%s",
	}`, i.common.StoreName, path)

	str, err := i.common.PostWithLoginSession(url, params)
	if err != nil {
		return false, err
	}
	result := &sg.ErrorMsg{}
	err = json.Unmarshal([]byte(str), result)
	if err != nil {
		return false, err
	}
	if result.ErrNo != 0 {
		return false, result.Error()
	}

	return true, nil
}

//非空目录不能删除，需要逐级删除
func (i *instance) listDirectory(config string) ([]*api.DetailFiles, error) {
	url := i.common.Command("/commands/get_file_list.action")
	params := make(map[string]string)
	params["params"] = config

	str, err := i.common.PostWithLoginSession(url, params)
	if err != nil {
		return nil, err
	}
	result := &DirectoryList{}
	err = json.Unmarshal([]byte(str), result)
	if err != nil {
		return nil, err
	}
	if result.ErrNo != 0 {
		return nil, result.Error()
	}
	for _, v := range result.Data.DetailFiles {
		//去除存储服务器名称
		v.PosixPath = strings.TrimPrefix(v.Path, i.common.StoreName+":")
	}
	return result.Data.DetailFiles, nil
}

//列表显示目录
func (i *instance) ListDirectory(path string) ([]*api.DetailFiles, error) {

	config := fmt.Sprintf(`{
	"limit":1000000,
	"start":0,
	"sort":"",
	"path":"%s:%s",	
	"display_details":true,
	"type":"DIR",
	"searches":[{"searchKey":"name","searchValue":""}]
	}`, i.common.StoreName, path)

	return i.listDirectory(config)
}

//列表显示目录及文件
func (i *instance) ListDirectoryWithFiles(path string) ([]*api.DetailFiles, error) {
	config := fmt.Sprintf(`{
	"limit":1000000,
	"start":0,
	"sort":"",
	"path":"%s:%s",	
	"display_details":true,
	"searches":[{"searchKey":"name","searchValue":""}]
	}`, i.common.StoreName, path)
	return i.listDirectory(config)

}
