package pkg

import (
	c "context"
	. "github.com/smartystreets/goconvey/convey"
	"go-gin-api/boot"
	"go-gin-api/pkg/g"
	"testing"
)

func TestRedis(t *testing.T) {
	boot.Init()
	ctx := c.Background()
	Convey("redis ping", t, func() {
		So(g.Redis().Ping(ctx).Err(), ShouldBeNil)
	})
}
