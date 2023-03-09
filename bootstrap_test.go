package netty

import (
	"testing"
	"time"
)

func TestBootstrap(t *testing.T) {
	ci := func(channel Channel) {
		channel.Pipeline().
			AddLast(delimiterCodec{maxFrameLength: 1024, delimiter: []byte("$"), stripDelimiter: true}).
			AddLast(ReadIdleHandler(time.Second), WriteIdleHandler(time.Second)).
			AddLast(&textCodec{}).
			AddLast(&echoHandler{}).
			AddLast(&eventHandler{})
	}

	b := NewBootstrap().
		ChildHandler(NewChannelInitializer(ci))
	f := b.Bind(10000).Sync()
	b.Bind(10001).Sync()
}
