// Code generated by Kitex v0.4.4. DO NOT EDIT.

package commentservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	comment "mini-tiktok-backend/kitex_gen/comment"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	CreateComment(ctx context.Context, req *comment.CreateCommentRequest, callOptions ...callopt.Option) (r *comment.CreateCommentResponse, err error)
	DeleteComment(ctx context.Context, req *comment.DeleteCommentRequest, callOptions ...callopt.Option) (r *comment.DeleteCommentResponse, err error)
	GetCommentList(ctx context.Context, req *comment.GetCommentListRequest, callOptions ...callopt.Option) (r *comment.GetCommentListResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kCommentServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kCommentServiceClient struct {
	*kClient
}

func (p *kCommentServiceClient) CreateComment(ctx context.Context, req *comment.CreateCommentRequest, callOptions ...callopt.Option) (r *comment.CreateCommentResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateComment(ctx, req)
}

func (p *kCommentServiceClient) DeleteComment(ctx context.Context, req *comment.DeleteCommentRequest, callOptions ...callopt.Option) (r *comment.DeleteCommentResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteComment(ctx, req)
}

func (p *kCommentServiceClient) GetCommentList(ctx context.Context, req *comment.GetCommentListRequest, callOptions ...callopt.Option) (r *comment.GetCommentListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetCommentList(ctx, req)
}
