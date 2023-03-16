package netty

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
	Write(message any)
	WriteAndFlush(message any)
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
	return nil
}

type channelFuture struct {
	channel *Channel
}

func (f *channelFuture) Channel() *Channel {
	return f.channel
}

func (f *channelFuture) Sync() *channelFuture {
	return nil
}

func (f *channelFuture) Async(fn func(f *channelFuture)) {
	go func() {
		fn(f.Sync())
	}()
}

func (f *channelFuture) Cancel() error {
	return nil
}
