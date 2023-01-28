// Code generated by hertz generator.

package api

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	api "mini-tiktok-backend/cmd/api/biz/model/api"
	"mini-tiktok-backend/cmd/api/biz/mw"
	"mini-tiktok-backend/cmd/api/biz/rpc"
	"mini-tiktok-backend/kitex_gen/user"
	"mini-tiktok-backend/pkg/errno"
	"strconv"
)

// CheckUser .
// @router /douyin/user/login/ [POST]
func CheckUser(ctx context.Context, c *app.RequestContext) {
	mw.JwtMiddleware.LoginHandler(ctx, c)
}

// CreateUser .
// @router /douyin/user/register/ [POST]
func CreateUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.CreateUserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.CreateUserResponse)
	err = rpc.CreateUser(context.Background(), &user.CreateUserRequest{
		Username: req.Username,
		Password: req.Password,
	})

	if err != nil {
		Err := errno.ConvertErr(err)
		resp.StatusCode = Err.ErrCode
		resp.StatusMsg = Err.ErrMsg
		c.JSON(consts.StatusOK, resp)
		return
	}

	// login after register and return login response message
	mw.JwtMiddleware.LoginHandler(ctx, c)
}

// QueryUser .
// @router /douyin/user/ [GET]
func QueryUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.QueryUserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.QueryUserResponse)
	resp.User = new(api.User)

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
