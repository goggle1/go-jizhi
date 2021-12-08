package client

import (
	"context"

	"git.tvblack.com/video/frame/proto/p_rpc"
	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
)

type adminKey struct{}

func AdminFromContext(ctx context.Context) (p_rpc.MessageService, bool) {
	c, ok := ctx.Value(adminKey{}).(p_rpc.MessageService)
	return c, ok
}

//AdminWrapper returns a wrapper for the Wrapper
func AdminWrapper(service micro.Service) server.HandlerWrapper {
	client := p_rpc.NewMessageService("live.admin.srv", service.Client())

	return func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			ctx = context.WithValue(ctx, adminKey{}, client)
			return fn(ctx, req, rsp)
		}
	}
}
