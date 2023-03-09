package netty

import (
	"testing"
	"time"
)

func TestBootstrap(t *testing.T) {
	ci := NewChannelInitializer(func(channel Channel) {
		channel.Pipeline().
			AddLast(delimiterCodec{maxFrameLength: 1024, delimiter: []byte("$"), stripDelimiter: true}).
			AddLast(ReadIdleHandler(time.Second), WriteIdleHandler(time.Second)).
			AddLast(&textCodec{}).
			AddLast(&echoHandler{}).
			AddLast(&eventHandler{})
	})
	b := NewBootstrap().ChildHandler(ci)
	b.Bind(10000).Channel().CloseFuture().Sync()
}
