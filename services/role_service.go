package services

import (
	"encoding/json"
	"strconv"

	"github.com/blessli/ranger/services/sysclient"
	"github.com/zeromicro/go-zero/core/logx"
)

func (s *loginService) RoleList(in *ListRoleReq) (*ListRoleResp, error) {
	all, err := s.RoleModel.FindAll(in.Current, in.PageSize, in.Name)
	count, _ := s.RoleModel.Count()

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.Errorf("查询角色列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	var list []*sysclient.RoleListData
	for _, role := range *all {
		list = append(list, &sysclient.RoleListData{
			Id:             role.Id,
			Name:           role.Name,
			Remark:         role.Remark,
			CreateBy:       role.CreateBy,
			CreateTime:     role.CreateTime.Format("2006-01-02 15:04:05"),
			LastUpdateBy:   role.LastUpdateBy,
			LastUpdateTime: role.LastUpdateTime.Format("2006-01-02 15:04:05"),
			DelFlag:        role.DelFlag,
			Status:         role.Status,
		})
	}

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(list)
	logx.Infof("查询角色列表信息,参数：%s,响应：%s", reqStr, listStr)
	resp := &sysclient.RoleListResp{
		Total: count,
		List:  list,
	}
	var list1 []*ListRoleData

	for _, role := range resp.List {
		list1 = append(list1, &ListRoleData{
			Id:             role.Id,
			Name:           role.Name,
			Remark:         role.Remark,
			CreateBy:       role.CreateBy,
			CreateTime:     role.CreateTime,
			LastUpdateBy:   role.LastUpdateBy,
			LastUpdateTime: role.LastUpdateTime,
			DelFlag:        role.DelFlag,
			Label:          role.Name,
			Value:          strconv.FormatInt(role.Id, 10),
			Status:         role.Status,
		})
	}

	return &ListRoleResp{
		Code:     "000000",
		Message:  "查询角色成功",
		Current:  in.Current,
		Data:     list1,
		PageSize: in.PageSize,
		Success:  true,
		Total:    resp.Total,
	}, nil
}
