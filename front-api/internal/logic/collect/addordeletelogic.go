package collect

import (
	"context"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddordeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddordeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddordeleteLogic {
	return AddordeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddordeleteLogic) Addordelete(req types.AddOrDeleteReq) (resp *types.AddOrDeleteResp, err error) {
	// todo: add your logic here and delete this line

	return
}
