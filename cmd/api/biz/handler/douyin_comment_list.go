// Code generated by hertz generator.

package handler

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	comment "mini-tiktok-backend/cmd/api/biz/model/api/comment"
)

// DouyinCommentList .
// @router /douyin/comment/list/ [GET]
func DouyinCommentList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req comment.DouyinCommentListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(comment.DouyinCommentListResponse)

	c.JSON(consts.StatusOK, resp)
}