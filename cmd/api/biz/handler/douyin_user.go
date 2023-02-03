// Code generated by hertz generator.

package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	api_user "mini-tiktok-backend/cmd/api/biz/model/api/user"
	"mini-tiktok-backend/cmd/api/biz/rpc"
	"mini-tiktok-backend/kitex_gen/user"
	"mini-tiktok-backend/pkg/errno"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
)

// DouyinUser .
// @router /douyin/user/ [GET]
func DouyinUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api_user.DouyinUserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api_user.DouyinUserResponse)
	resp.User = new(api_user.User)

	id, err := strconv.ParseInt(req.UserID, 10, 64)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	username, err := rpc.QueryUser(ctx, &user.QueryUserRequest{
		UserId: id,
	})

	if err != nil {
		Err := errno.ConvertErr(err)
		resp.StatusCode = Err.ErrCode
		resp.StatusMsg = Err.ErrMsg
		c.JSON(consts.StatusOK, resp)
		return
	}

	resp.User.ID = id
	resp.User.Name = username

	c.JSON(consts.StatusOK, resp)
}
