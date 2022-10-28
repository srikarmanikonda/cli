package tailnet

import (
	"context"

	"tailscale.com/tsnet"
)

type tailnetServerKey struct{}

func WithTailnetServer(ctx context.Context, server *tsnet.Server) context.Context {
	return context.WithValue(ctx, tailnetServerKey{}, server)
}

func GetContextServer(ctx context.Context) *tsnet.Server {
	return ctx.Value(tailnetServerKey{}).(*tsnet.Server)
}
