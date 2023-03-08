package logic

import (
	"context"
	"strconv"
	"time"

	"zero_demo/api_user/internal/svc"
	"zero_demo/api_user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"zero_demo/rpc_user/userclient"
)

type Api_userLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewApi_userLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Api_userLogic {
	return &Api_userLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Api_userLogic) Api_user(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	var (
		rsp *userclient.UserInfoReply
	)
	rsp, err = l.svcCtx.User.GetUser(l.ctx, &userclient.IdReq{Id: time.Now().Unix()})
	if err != nil {
		return nil, err
	}
	resp = &types.Response{}
	resp.Message = strconv.FormatInt(rsp.Id, 10)
	return
}
