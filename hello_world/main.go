package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func main() {
	h := server.Default()

	h.GET("/hello", func(ctx context.Context, c *app.RequestContext) {
		// c.JSON(consts.StatusOK, utils.H{"message": "pong"})
		c.Data(consts.StatusOK, consts.MIMETextPlain, []byte("hell hertz"))
	})

	h.Spin()
}
