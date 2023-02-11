package services

type UserInfoResp struct {
	Code        string             `json:"code"`
	Message     string             `json:"message"`
	Avatar      string             `json:"avatar"`
	Name        string             `json:"name"`
	MenuTree    []*ListMenuTree    `json:"menuTree"`
	MenuTreeVue []*ListMenuTreeVue `json:"menuTreeVue"`
}

type LoginResp struct {
	Code             string `json:"code"`
	Message          string `json:"message"`
	Status           string `json:"status"`
	CurrentAuthority string `json:"currentAuthority"`
	Id               int64  `json:"id"`
	UserName         string `json:"userName"`
	AccessToken      string `json:"token"`
	AccessExpire     int64  `json:"accessExpire"`
	RefreshAfter     int64  `json:"refreshAfter"`
}

type ListMenuTree struct {
	Id       int64  `json:"id"`       // 编号
	Path     string `json:"path"`     // 菜单路径
	Name     string `json:"name"`     // 菜单名称
	ParentId int64  `json:"parentId"` // 父菜单ID，一级菜单为0
	Icon     string `json:"icon"`     // 菜单图标
}

type ListMenuTreeVue struct {
	Id           int64        `json:"id"`
	ParentId     int64        `json:"parentId"`
	Title        string       `json:"title"`
	Path         string       `json:"path"`
	Name         string       `json:"name"`
	Icon         string       `json:"icon"`
	VueRedirect  string       `json:"redirect"`
	VueComponent string       `json:"component"`
	Meta         MenuTreeMeta `json:"meta"`
}

type MenuTreeMeta struct {
	Title string `json:"title"`
	Icon  string `json:"icon"`
}

type InfoResp struct {
	Avatar         string          `json:"avatar"`
	Name           string          `json:"name"`
	MenuListTree   []*MenuListTree `json:"menuListTree"`
	BackgroundUrls []string        `json:"BackgroundUrls"`
}

type MenuListTree struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	Icon         string `json:"icon"`
	ParentId     int64  `json:"parent_id"`
	Path         string `json:"path"`
	VuePath      string `json:"vue_path"`
	VueComponent string `json:"vue_component"`
	VueIcon      string `json:"vue_icon"`
	VueRedirect  string `json:"vue_redirect"`
}

type ListUserResp struct {
	Code     string          `json:"code"`
	Message  string          `json:"message"`
	Current  int64           `json:"current,default=1"`
	Data     []*ListUserData `json:"data"`
	PageSize int64           `json:"pageSize,default=20"`
	Success  bool            `json:"success"`
	Total    int64           `json:"total"`
}
type ListUserData struct {
	Id             int64  `json:"id"`             // 编号
	Name           string `json:"name"`           // 用户名
	NickName       string `json:"nickName"`       // 昵称
	Avatar         string `json:"avatar"`         // 头像
	Password       string `json:"password"`       // 密码
	Salt           string `json:"salt"`           // 加密盐
	Email          string `json:"email"`          // 邮箱
	Mobile         string `json:"mobile"`         // 手机号
	Status         int64  `json:"status"`         // 状态  0：禁用   1：正常
	DeptId         int64  `json:"deptId"`         // 机构ID
	CreateBy       string `json:"createBy"`       // 创建人
	CreateTime     string `json:"createTime"`     // 创建时间
	LastUpdateBy   string `json:"lastUpdateBy"`   // 更新人
	LastUpdateTime string `json:"lastUpdateTime"` // 更新时间
	DelFlag        int64  `json:"delFlag"`        // 是否删除  -1：已删除  0：正常
	JobId          int64  `json:"jobId"`
	RoleId         int64  `json:"roleId"`
	RoleName       string `json:"roleName"`
	JobName        string `json:"jobName"`
	DeptName       string `json:"deptName"`
}

type ListRoleData struct {
	Id             int64  `json:"id"`             // 编号
	Name           string `json:"name"`           // 角色名称
	Remark         string `json:"remark"`         // 备注
	CreateBy       string `json:"createBy"`       // 创建人
	CreateTime     string `json:"createTime"`     // 创建时间
	LastUpdateBy   string `json:"lastUpdateBy"`   // 更新人
	LastUpdateTime string `json:"lastUpdateTime"` // 更新时间
	DelFlag        int64  `json:"delFlag"`        // 是否删除  -1：已删除  0：正常
	Label          string `json:"label"`          // 编号
	Value          string `json:"value"`          // 角色名称
	Status         int64  `json:"status"`         // 角色名称
}
type ListRoleResp struct {
	Code     string          `json:"code"`
	Message  string          `json:"message"`
	Current  int64           `json:"current,default=1"`
	Data     []*ListRoleData `json:"data"`
	PageSize int64           `json:"pageSize,default=20"`
	Success  bool            `json:"success"`
	Total    int64           `json:"total"`
}