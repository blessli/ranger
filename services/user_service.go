package services

import (
	"encoding/json"

	"github.com/blessli/ranger/services/sysclient"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

func (s *loginService) UserList(in *ListUserReq) (*ListUserResp, error) {
	all, err := s.UserModel.FindAll(in.Current, in.PageSize)

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.Errorf("查询用户列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	count, _ := s.UserModel.Count()

	var list []*sysclient.UserListData
	for _, user := range *all {
		list = append(list, &sysclient.UserListData{
			Id:             user.Id,
			Name:           user.Name,
			NickName:       user.NickName,
			Avatar:         user.Avatar,
			Password:       user.Password,
			Salt:           user.Salt,
			Email:          user.Email,
			Mobile:         user.Mobile,
			DeptId:         user.DeptId,
			Status:         user.Status,
			CreateBy:       user.CreateBy,
			CreateTime:     user.CreateTime.Format("2006-01-02 15:04:05"),
			LastUpdateBy:   user.LastUpdateBy,
			LastUpdateTime: user.LastUpdateTime.Format("2006-01-02 15:04:05"),
			DelFlag:        user.DelFlag,
			JobId:          user.JobId,
			RoleId:         user.RoleId,
			RoleName:       user.RoleName,
			JobName:        user.JobName,
			DeptName:       user.DeptName,
		})
	}

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(list)
	logx.Infof("查询用户列表信息,参数：%s,响应：%s", reqStr, listStr)
	resp := &sysclient.UserListResp{
		Total: count,
		List:  list,
	}
	var list1 []*ListUserData

	for _, item := range resp.List {
		listUserData := ListUserData{}
		_ = copier.Copy(&listUserData, &item)
		list1 = append(list1, &listUserData)
	}

	return &ListUserResp{
		Code:     "000000",
		Message:  "查询用户列表成功",
		Current:  in.Current,
		Data:     list1,
		PageSize: in.PageSize,
		Success:  true,
		Total:    resp.Total,
	}, nil
}
