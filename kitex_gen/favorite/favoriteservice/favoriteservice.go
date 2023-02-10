// Code generated by Kitex v0.4.4. DO NOT EDIT.

package favoriteservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	favorite "mini-tiktok-backend/kitex_gen/favorite"
)

func serviceInfo() *kitex.ServiceInfo {
	return favoriteServiceServiceInfo
}

var favoriteServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "FavoriteService"
	handlerType := (*favorite.FavoriteService)(nil)
	methods := map[string]kitex.MethodInfo{
		"FavoriteAction":  kitex.NewMethodInfo(favoriteActionHandler, newFavoriteServiceFavoriteActionArgs, newFavoriteServiceFavoriteActionResult, false),
		"GetIsFavorite":   kitex.NewMethodInfo(getIsFavoriteHandler, newFavoriteServiceGetIsFavoriteArgs, newFavoriteServiceGetIsFavoriteResult, false),
		"GetFavoriteList": kitex.NewMethodInfo(getFavoriteListHandler, newFavoriteServiceGetFavoriteListArgs, newFavoriteServiceGetFavoriteListResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "favorite",
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

func favoriteActionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*favorite.FavoriteServiceFavoriteActionArgs)
	realResult := result.(*favorite.FavoriteServiceFavoriteActionResult)
	success, err := handler.(favorite.FavoriteService).FavoriteAction(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newFavoriteServiceFavoriteActionArgs() interface{} {
	return favorite.NewFavoriteServiceFavoriteActionArgs()
}

func newFavoriteServiceFavoriteActionResult() interface{} {
	return favorite.NewFavoriteServiceFavoriteActionResult()
}

func getIsFavoriteHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*favorite.FavoriteServiceGetIsFavoriteArgs)
	realResult := result.(*favorite.FavoriteServiceGetIsFavoriteResult)
	success, err := handler.(favorite.FavoriteService).GetIsFavorite(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newFavoriteServiceGetIsFavoriteArgs() interface{} {
	return favorite.NewFavoriteServiceGetIsFavoriteArgs()
}

func newFavoriteServiceGetIsFavoriteResult() interface{} {
	return favorite.NewFavoriteServiceGetIsFavoriteResult()
}

func getFavoriteListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*favorite.FavoriteServiceGetFavoriteListArgs)
	realResult := result.(*favorite.FavoriteServiceGetFavoriteListResult)
	success, err := handler.(favorite.FavoriteService).GetFavoriteList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newFavoriteServiceGetFavoriteListArgs() interface{} {
	return favorite.NewFavoriteServiceGetFavoriteListArgs()
}

func newFavoriteServiceGetFavoriteListResult() interface{} {
	return favorite.NewFavoriteServiceGetFavoriteListResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) FavoriteAction(ctx context.Context, req *favorite.FavoriteActionRequest) (r *favorite.FavoriteActionResponse, err error) {
	var _args favorite.FavoriteServiceFavoriteActionArgs
	_args.Req = req
	var _result favorite.FavoriteServiceFavoriteActionResult
	if err = p.c.Call(ctx, "FavoriteAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetIsFavorite(ctx context.Context, req *favorite.GetIsFavoriteRequest) (r *favorite.GetIsFavoriteResponse, err error) {
	var _args favorite.FavoriteServiceGetIsFavoriteArgs
	_args.Req = req
	var _result favorite.FavoriteServiceGetIsFavoriteResult
	if err = p.c.Call(ctx, "GetIsFavorite", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetFavoriteList(ctx context.Context, req *favorite.GetFavoriteListRequest) (r *favorite.GetFavoriteListResponse, err error) {
	var _args favorite.FavoriteServiceGetFavoriteListArgs
	_args.Req = req
	var _result favorite.FavoriteServiceGetFavoriteListResult
	if err = p.c.Call(ctx, "GetFavoriteList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
