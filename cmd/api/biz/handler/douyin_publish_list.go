// Code generated by hertz generator.

package handler

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	publish "mini-tiktok-backend/cmd/api/biz/model/api/publish"
)

// DouyinPublishList .
// @router /douyin/publish/list/ [GET]
func DouyinPublishList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req publish.DouyinPublishListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(publish.DouyinPublishListResponse)

	c.JSON(consts.StatusOK, resp)
}