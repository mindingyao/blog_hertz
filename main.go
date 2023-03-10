// Code generated by hertz generator.

package main

import (
	"blog_hertz/biz/dal"
	"blog_hertz/biz/mw"
	"context"
	"os"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/sessions"
	"github.com/hertz-contrib/sessions/cookie"
)

func main() {
	f, err := os.OpenFile("./biz/logs/tag.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	hlog.SetOutput(f)

	dal.Init()
	h := server.Default()
	store := cookie.NewStore([]byte("secret"))
	h.Use(sessions.New("mysession", store))
	h.Use(mw.AccessLog())
	register(h)
	h.GET("/hello", func(ctx context.Context, c *app.RequestContext) {
		session := sessions.Default(c)

		if session.Get("hello") != "world" {
			session.Set("hello", "world")
			session.Save()
		}

		c.JSON(200, utils.H{"hello": session.Get("hello")})
	})
	h.Spin()
}
