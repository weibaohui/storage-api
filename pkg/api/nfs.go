package api

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
type NFSApi interface {
	//创建目录
	CreateDirectory(path string) (ok bool, err error)
	//删除目录
	DeleteDirectory(path string) (ok bool, err error)
	//创建配额
	CreateQuota(path string, ips, ops, readBw, writeBw int) (ok bool, quotaID string, err error)
	//删除配额
	DeleteQuota(id string) (ok bool, err error)

	//列表显示文件夹及文件
	ListDirectoryWithFiles(path string) ([]*DetailFiles, error)
	//列表显示文件夹
	ListDirectory(path string) ([]*DetailFiles, error)
}
