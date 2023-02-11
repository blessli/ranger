package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/blessli/ranger/dao"
	"github.com/dgrijalva/jwt-go"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type loginService struct {
	UserModel *dao.SysUserModel
	MenuModel *dao.SysMenuModel
	RoleModel *dao.SysRoleModel
}

func NewLoginService() *loginService {
	sqlConn := sqlx.NewMysql("root:rootroot@tcp(127.0.0.1:3306)/gozero?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai")
	return &loginService{
		UserModel: dao.NewSysUserModel(sqlConn),
		MenuModel: dao.NewSysMenuModel(sqlConn),
		RoleModel: dao.NewSysRoleModel(sqlConn),
	}
}

func (s *loginService) Login(username, password string) (*LoginResp, error) {
	userInfo, err := s.UserModel.FindOneByName(username)
	switch err {
	case nil:
	case sqlc.ErrNotFound:
		logx.Errorf("用户不存在,参数:%s,异常:%s", username, err.Error())
		return nil, errors.New("用户不存在")
	default:
		logx.Errorf("用户登录失败,参数:%s,异常:%s", username, err.Error())
		return nil, err
	}

	if userInfo.Password != password {
		logx.Errorf("用户密码不正确,参数:%s", password)
		return nil, errors.New("用户密码不正确")
	}

	now := time.Now().Unix()
	accessExpire := int64(86400)
	jwtToken, err := getJwtToken("ad879037-c7a4-4063-9236-6bfc35d54b7d", now, accessExpire, userInfo.Id)

	if err != nil {
		logx.Errorf("生成token失败,参数:%s,%s,异常:%s", username, password, err.Error())
		return nil, err
	}

	resp := &LoginResp{
		Code:             "000000",
		Status:           "ok",
		CurrentAuthority: "admin",
		Id:               userInfo.Id,
		UserName:         userInfo.Name,
		AccessToken:      jwtToken,
		AccessExpire:     now + accessExpire,
		RefreshAfter:     now + accessExpire/2,
	}

	listStr, _ := json.Marshal(resp)
	logx.Infof("登录成功,参数:%s,%s,响应:%s", username, password, listStr)
	return resp, nil
}

func (s *loginService) UserInfo(userId int64) (*UserInfoResp, error) {
	userInfo, err := s.UserModel.FindOne(userId)

	switch err {
	case nil:
	case sqlc.ErrNotFound:
		logx.Infof("用户不存在userId: %s", userId)
		return nil, fmt.Errorf("用户不存在userId: %s", userId)
	default:
		return nil, err
	}

	var list []*MenuListTree
	var listUrls []string

	if userId == 1 {
		menus, _ := s.MenuModel.FindAll(1, 1000)
		list, listUrls = listTrees(menus, list, listUrls)
		logx.Infof("超级管理员: %s登录,菜单: %+v", userInfo.Name, list)
	} else {
		menus, _ := s.MenuModel.FindAllByUserId(userId)
		list, listUrls = listTrees(menus, list, listUrls)
		logx.Infof("普通管理员: %s登录,菜单: %+v", userInfo.Name, list)
	}

	resp := &InfoResp{
		Avatar:         "https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png",
		Name:           userInfo.Name,
		MenuListTree:   list,
		BackgroundUrls: listUrls,
	}
	var MenuTree []*ListMenuTree

	//组装ant ui中的菜单
	for _, item := range resp.MenuListTree {
		MenuTree = append(MenuTree, &ListMenuTree{
			Id:       item.Id,
			Path:     item.Path,
			Name:     item.Name,
			ParentId: item.ParentId,
			Icon:     item.Icon,
		})
	}
	//组装element ui中的菜单
	var MenuTreeVue []*ListMenuTreeVue

	for _, item := range resp.MenuListTree {

		if len(strings.TrimSpace(item.VuePath)) != 0 {
			MenuTreeVue = append(MenuTreeVue, &ListMenuTreeVue{
				Id:           item.Id,
				ParentId:     item.ParentId,
				Title:        item.Name,
				Path:         item.VuePath,
				Name:         item.Name,
				Icon:         item.VueIcon,
				VueRedirect:  item.VueRedirect,
				VueComponent: item.VueComponent,
				Meta: MenuTreeMeta{
					Title: item.Name,
					Icon:  item.VueIcon,
				},
			})
		}

	}
	rsp := &UserInfoResp{
		Code:        "000000",
		Message:     "获取个人信息成功",
		Avatar:      resp.Avatar,
		Name:        resp.Name,
		MenuTree:    MenuTree,
		MenuTreeVue: MenuTreeVue,
	}
	return rsp, nil
}

func listTrees(menus *[]dao.SysMenu, list []*MenuListTree, listUrls []string) ([]*MenuListTree, []string) {
	for _, menu := range *menus {
		list = append(list, &MenuListTree{
			Id:           menu.Id,
			Name:         menu.Name,
			Icon:         menu.Icon,
			ParentId:     menu.ParentId,
			Path:         menu.Url,
			VuePath:      menu.VuePath,
			VueComponent: menu.VueComponent,
			VueIcon:      menu.VueIcon,
			VueRedirect:  menu.VueRedirect,
		})

		if len(strings.TrimSpace(menu.BackgroundUrl)) != 0 {
			listUrls = append(listUrls, menu.BackgroundUrl)
		}

	}
	return list, listUrls
}

func getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
