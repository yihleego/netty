package netty

import (
	"context"
)

type Bootstrap interface {
	Context() context.Context
	Handler(handler ChannelHandler) Bootstrap
	ChildHandler(childHandler ChannelHandler) Bootstrap
	Connect(host string, port int) ChannelFuture
	Bind(port int) ChannelFuture
	Shutdown()
}

type bootstrap struct {
}

func (b *bootstrap) Context() context.Context {
	return nil
}
func (b *bootstrap) Handler(handler ChannelHandler) Bootstrap {
	return nil
}
func (b *bootstrap) ChildHandler(childHandler ChannelHandler) Bootstrap {
	return nil
}
func (b *bootstrap) Connect(host string, port int) ChannelFuture {
	return nil
}
func (b *bootstrap) Bind(port int) ChannelFuture {
	return nil
}
func (b *bootstrap) Shutdown() {
	return
}
func NewBootstrap() Bootstrap {
	return &bootstrap{}
}
