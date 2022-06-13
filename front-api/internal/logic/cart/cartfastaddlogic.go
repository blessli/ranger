package cart

import (
	"context"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CartFastAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCartFastAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) CartFastAddLogic {
	return CartFastAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CartFastAddLogic) CartFastAdd(req types.CartFastAddReq) (resp *types.CartFastAddResp, err error) {
	// todo: add your logic here and delete this line

	return
}
