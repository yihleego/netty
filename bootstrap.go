package netty

import (
	"context"
)

type Bootstrap interface {
	Context() context.Context
	Handler(handler ChannelHandler) Bootstrap
	Connect(host string, port int) ChannelFuture
	Shutdown()
}

type ServerBootstrap interface {
	Bootstrap
	ChildHandler(childHandler ChannelHandler) ServerBootstrap
	Bind(port int) ChannelFuture
}

type bootstrap struct {
	ctx    context.Context
	cancel context.CancelFunc
}

func (b *bootstrap) Context() context.Context {
	return nil
}

func (b *bootstrap) Handler(handler ChannelHandler) Bootstrap {
	return nil
}

func (b *bootstrap) Connect(host string, port int) ChannelFuture {
	return nil
}

func (b *bootstrap) Shutdown() {
	return
}

type serverBootstrap struct {
	bootstrap
}

func (b *serverBootstrap) ChildHandler(childHandler ChannelHandler) ServerBootstrap {
	return nil
}

func (b *serverBootstrap) Bind(port int) ChannelFuture {
	return nil
}

func NewBootstrap() Bootstrap {
	ctx, cancel := context.WithCancel(context.Background())
	return &bootstrap{ctx, cancel}
}

func NewServerBootstrap() ServerBootstrap {
	ctx, cancel := context.WithCancel(context.Background())
	return &serverBootstrap{bootstrap{ctx, cancel}}
}
