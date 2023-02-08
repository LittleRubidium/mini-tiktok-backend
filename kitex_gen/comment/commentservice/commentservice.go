// Code generated by Kitex v0.4.4. DO NOT EDIT.

package commentservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	comment "mini-tiktok-backend/kitex_gen/comment"
)

func serviceInfo() *kitex.ServiceInfo {
	return commentServiceServiceInfo
}

var commentServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "CommentService"
	handlerType := (*comment.CommentService)(nil)
	methods := map[string]kitex.MethodInfo{
		"CreateComment":   kitex.NewMethodInfo(createCommentHandler, newCommentServiceCreateCommentArgs, newCommentServiceCreateCommentResult, false),
		"DeleteComment":   kitex.NewMethodInfo(deleteCommentHandler, newCommentServiceDeleteCommentArgs, newCommentServiceDeleteCommentResult, false),
		"GetCommentCount": kitex.NewMethodInfo(getCommentCountHandler, newCommentServiceGetCommentCountArgs, newCommentServiceGetCommentCountResult, false),
		"GetCommentList":  kitex.NewMethodInfo(getCommentListHandler, newCommentServiceGetCommentListArgs, newCommentServiceGetCommentListResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "comment",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func createCommentHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*comment.CommentServiceCreateCommentArgs)
	realResult := result.(*comment.CommentServiceCreateCommentResult)
	success, err := handler.(comment.CommentService).CreateComment(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCommentServiceCreateCommentArgs() interface{} {
	return comment.NewCommentServiceCreateCommentArgs()
}

func newCommentServiceCreateCommentResult() interface{} {
	return comment.NewCommentServiceCreateCommentResult()
}

func deleteCommentHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*comment.CommentServiceDeleteCommentArgs)
	realResult := result.(*comment.CommentServiceDeleteCommentResult)
	success, err := handler.(comment.CommentService).DeleteComment(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCommentServiceDeleteCommentArgs() interface{} {
	return comment.NewCommentServiceDeleteCommentArgs()
}

func newCommentServiceDeleteCommentResult() interface{} {
	return comment.NewCommentServiceDeleteCommentResult()
}

func getCommentCountHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*comment.CommentServiceGetCommentCountArgs)
	realResult := result.(*comment.CommentServiceGetCommentCountResult)
	success, err := handler.(comment.CommentService).GetCommentCount(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCommentServiceGetCommentCountArgs() interface{} {
	return comment.NewCommentServiceGetCommentCountArgs()
}

func newCommentServiceGetCommentCountResult() interface{} {
	return comment.NewCommentServiceGetCommentCountResult()
}

func getCommentListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*comment.CommentServiceGetCommentListArgs)
	realResult := result.(*comment.CommentServiceGetCommentListResult)
	success, err := handler.(comment.CommentService).GetCommentList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCommentServiceGetCommentListArgs() interface{} {
	return comment.NewCommentServiceGetCommentListArgs()
}

func newCommentServiceGetCommentListResult() interface{} {
	return comment.NewCommentServiceGetCommentListResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) CreateComment(ctx context.Context, req *comment.CreateCommentRequest) (r *comment.CreateCommentResponse, err error) {
	var _args comment.CommentServiceCreateCommentArgs
	_args.Req = req
	var _result comment.CommentServiceCreateCommentResult
	if err = p.c.Call(ctx, "CreateComment", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteComment(ctx context.Context, req *comment.DeleteCommentRequest) (r *comment.DeleteCommentResponse, err error) {
	var _args comment.CommentServiceDeleteCommentArgs
	_args.Req = req
	var _result comment.CommentServiceDeleteCommentResult
	if err = p.c.Call(ctx, "DeleteComment", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetCommentCount(ctx context.Context, req *comment.GetCommentCountRequest) (r *comment.GetCommentCountResponse, err error) {
	var _args comment.CommentServiceGetCommentCountArgs
	_args.Req = req
	var _result comment.CommentServiceGetCommentCountResult
	if err = p.c.Call(ctx, "GetCommentCount", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetCommentList(ctx context.Context, req *comment.GetCommentListRequest) (r *comment.GetCommentListResponse, err error) {
	var _args comment.CommentServiceGetCommentListArgs
	_args.Req = req
	var _result comment.CommentServiceGetCommentListResult
	if err = p.c.Call(ctx, "GetCommentList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
