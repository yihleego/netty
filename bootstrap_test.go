package netty

import (
	"fmt"
	"testing"
)

func TestBootstrap(t *testing.T) {
	host, port := "0.0.0.0", 10000
	ci := NewChannelInitializer(func(channel Channel) {
		channel.Pipeline().
			AddLast(&LineBasedFrameDecoder{4096}).
			AddLast(&StringDecoder{}).
			AddLast(&LineEncoder{}).
			AddLast(&EchoClientHandler{})
	})
	b := NewBootstrap().Handler(ci)
	f := b.Connect(host, port).Sync()
	if f.IsSuccess() {
		t.Errorf("Connect %s:%d failed", host, port)
	}
}

func TestServerBootstrap(t *testing.T) {
	_, port := "0.0.0.0", 10000
	ci := NewChannelInitializer(func(channel Channel) {
		channel.Pipeline().
			AddLast(&LineBasedFrameDecoder{4096}).
			AddLast(&StringDecoder{}).
			AddLast(&LineEncoder{}).
			AddLast(&EchoServerHandler{})
	})
	b := NewServerBootstrap().ChildHandler(ci)
	f := b.Bind(port).Sync()
	if f.IsSuccess() {
		t.Errorf("Initialized with port(s): %d", port)
	} else {
		t.Errorf("Bind port(s) %d failed", port)
	}
}

type EchoClientHandler struct {
	ChannelHandler
}

func (e *EchoClientHandler) ChannelActive(ctx ChannelHandlerContext) {
	message := "Hello, World!"
	fmt.Printf("client >> \"%s\"\n", message)
	ctx.WriteAndFlush(message)
}

func (e *EchoClientHandler) ChannelRead(ctx ChannelHandlerContext, message any) {
	fmt.Printf("client << \"%s\"\n", message)
}

type EchoServerHandler struct {
	ChannelHandler
}

func (e *EchoServerHandler) ChannelRead(ctx ChannelHandlerContext, message any) {
	fmt.Printf("server << \"%s\"\n", message)
	ctx.WriteAndFlush(message)
}
