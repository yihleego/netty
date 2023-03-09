package netty

import "fmt"

type Channel interface {
	Id() int64
	Pipeline() ChannelPipeline
}

type ChannelHandler interface {
	HandlerAdded(ctx ChannelHandlerContext)
	HandlerRemoved(ctx ChannelHandlerContext)
}

type ChannelInboundHandler interface {
	ChannelHandler
}

type ChannelOutboundHandler interface {
	ChannelHandler
}

type ChannelInitializer interface {
	ChannelInboundHandler
	InitChannel(channel Channel)
}

type ChannelHandlerContext interface {
	Channel() Channel
	Handler() ChannelHandler
	Pipeline() ChannelPipeline
	Name() string
	IsRemoved() bool
}

type ChannelPipeline interface {
	Channel() Channel
	AddFirst(handlers ...ChannelHandler) ChannelPipeline
	AddLast(handlers ...ChannelHandler) ChannelPipeline
	AddFirstWithName(name string, handler ChannelHandler) ChannelPipeline
	AddLastWithName(name string, handler ChannelHandler) ChannelPipeline
	AddBefore(baseName string, name string, handler ChannelHandler) ChannelPipeline
	AddAfter(baseName string, name string, handler ChannelHandler) ChannelPipeline
	Remove(handler ChannelHandler) ChannelPipeline
	RemoveByName(name string) ChannelHandler
	RemoveFirst() ChannelHandler
	RemoveLast() ChannelHandler
}

type ChannelFuture interface {
	Future[ChannelFuture]
	Channel() *Channel
}

type channelInitializer struct {
	initChannel func(channel Channel)
}

func (i *channelInitializer) InitChannel(channel Channel) {
	i.initChannel(channel)
}

func NewChannelInitializer(initChannel func(channel Channel)) ChannelInitializer {
	return &channelInitializer{initChannel}
}

type channelFuture struct {
	channel *Channel
}

func (f *channelFuture) Channel() *Channel {
	return f.channel
}

func (f *channelFuture) Sync() error {

	if nil != f.acceptor {
		return fmt.Errorf("duplicate call Listener:Sync")
	}

	var err error
	if f.options, err = transport.ParseOptions(f.bs.Context(), f.url, f.option...); nil != err {
		return err
	}

	if f.acceptor, err = f.bs.transportFactory.Listen(f.options); nil != err {
		return err
	}

	for {
		// accept the transport
		t, err := f.acceptor.Accept()
		if nil != err {
			return err
		}

		select {
		case <-f.bs.Context().Done():
			// bootstrap has been closed
			return t.Close()
		default:
			// serve child transport
			f.bs.serveTransport(t, nil, true)
		}
	}
}

func (f *channelFuture) Async(fn func(err error)) {
	go func() {
		fn(f.Sync())
	}()
}

func (f *channelFuture) Cancel() error {
	if f.acceptor != nil {
		f.bs.removeListener(f.url)
		return f.acceptor.Close()
	}
	return nil
}
