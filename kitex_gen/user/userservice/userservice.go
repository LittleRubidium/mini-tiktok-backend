// Code generated by Kitex v0.4.4. DO NOT EDIT.

package userservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	user "mini-tiktok-backend/kitex_gen/user"
)

func serviceInfo() *kitex.ServiceInfo {
	return userServiceServiceInfo
}

var userServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "UserService"
	handlerType := (*user.UserService)(nil)
	methods := map[string]kitex.MethodInfo{
		"CheckUser": kitex.NewMethodInfo(checkUserHandler, newUserServiceCheckUserArgs, newUserServiceCheckUserResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "user",
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

func checkUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceCheckUserArgs)
	realResult := result.(*user.UserServiceCheckUserResult)
	success, err := handler.(user.UserService).CheckUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceCheckUserArgs() interface{} {
	return user.NewUserServiceCheckUserArgs()
}

func newUserServiceCheckUserResult() interface{} {
	return user.NewUserServiceCheckUserResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) CheckUser(ctx context.Context, req *user.CheckUserRequest) (r *user.CheckUserResponse, err error) {
	var _args user.UserServiceCheckUserArgs
	_args.Req = req
	var _result user.UserServiceCheckUserResult
	if err = p.c.Call(ctx, "CheckUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}