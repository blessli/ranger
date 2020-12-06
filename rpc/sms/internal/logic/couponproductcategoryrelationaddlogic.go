package logic

import (
	"context"

	"go-zero-admin/rpc/sms/internal/svc"
	"go-zero-admin/rpc/sms/sms"

	"github.com/tal-tech/go-zero/core/logx"
)

type CouponProductCategoryRelationAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCouponProductCategoryRelationAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CouponProductCategoryRelationAddLogic {
	return &CouponProductCategoryRelationAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CouponProductCategoryRelationAddLogic) CouponProductCategoryRelationAdd(in *sms.CouponProductCategoryRelationAddReq) (*sms.CouponProductCategoryRelationAddResp, error) {
	// todo: add your logic here and delete this line

	return &sms.CouponProductCategoryRelationAddResp{}, nil
}