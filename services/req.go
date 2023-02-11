package services

type UserListReq struct {
	Current  int64  `json:"current"`
	PageSize int64  `json:"pageSize"`
	Name     string `json:"name"`
	NickName string `json=nickName"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email"`
	Status   int64  `json:"status"`
	DeptId   int64  `json:"dept_id"`
}
type ListUserReq struct {
	Current  int64  `json:"current,default=1"`
	PageSize int64  `json:"pageSize,default=20"`
	Name     string `json:"name,optional"`
	NickName string `json:"nickName,optional"`
	Mobile   string `json:"mobile,optional"`
	Email    string `json:"email,optional"`
	Status   int64  `json:"status,optional"`
	DeptId   int64  `json:"deptId,optional"`
	JobId    int64  `json:"deptId,optional"`
}

type ListRoleReq struct {
	Current  int64  `json:"current,default=1"`
	PageSize int64  `json:"pageSize,default=20"`
	Name     string `json:"name,optional "`
	Status   int64  `json:"delFlag,optional "`
}