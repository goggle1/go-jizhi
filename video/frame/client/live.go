package client

import (
	"context"

	"git.tvblack.com/video/frame/proto/p_rpc"
	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
)

type liveSrvKey struct{}

func LiveSrvFromContext(ctx context.Context) (p_rpc.MessageService, bool) {
	c, ok := ctx.Value(liveSrvKey{}).(p_rpc.MessageService)
	return c, ok
}

//LiveSrvWrapper returns a wrapper for the Wrapper
func LiveSrvWrapper(service micro.Service) server.HandlerWrapper {
	client := p_rpc.NewMessageService("live.srv", service.Client())

	return func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			ctx = context.WithValue(ctx, liveSrvKey{}, client)
			return fn(ctx, req, rsp)
		}
	}
}
