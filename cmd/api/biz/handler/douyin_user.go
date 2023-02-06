// Code generated by hertz generator.

package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	api_user "mini-tiktok-backend/cmd/api/biz/model/api/user"
	"mini-tiktok-backend/cmd/api/biz/mw"
	"mini-tiktok-backend/cmd/api/biz/rpc"
	"mini-tiktok-backend/kitex_gen/user"
	pkg_consts "mini-tiktok-backend/pkg/consts"
	"mini-tiktok-backend/pkg/errno"
)

// DouyinUser .
// @router /douyin/user/ [GET]
func DouyinUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api_user.DouyinUserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		SendResponse(c, err, utils.H{})
		return
	}

	identityUser, _ := c.Get(pkg_consts.IdentityKey)

	resp, err := rpc.QueryUser(ctx, &user.QueryUserRequest{
		UserId:       identityUser.(*mw.User).UserId,
		TargetUserId: req.UserID,
	})
	if err != nil {
		SendResponse(c, err, utils.H{})
		return
	}

	SendResponse(c, errno.Success, utils.H{
		"user": resp,
	})
}
