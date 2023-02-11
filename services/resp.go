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
