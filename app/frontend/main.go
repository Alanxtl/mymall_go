// Code generated by hertz generator.

package main

import (
	"context"
	"github.com/Alanxtl/mymall_go/app/frontend/middleware"
	"github.com/hertz-contrib/sessions"
	"github.com/hertz-contrib/sessions/redis"
	"github.com/joho/godotenv"
	"os"
	"time"

	"github.com/Alanxtl/mymall_go/app/frontend/biz/router"
	myutils "github.com/Alanxtl/mymall_go/app/frontend/biz/utils"
	"github.com/Alanxtl/mymall_go/app/frontend/conf"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/cors"
	"github.com/hertz-contrib/gzip"
	"github.com/hertz-contrib/logger/accesslog"
	hertzlogrus "github.com/hertz-contrib/logger/logrus"
	"github.com/hertz-contrib/pprof"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	_ = godotenv.Load()
	// init dal
	// dal.Init()
	address := conf.GetConf().Hertz.Address
	h := server.New(server.WithHostPorts(address))

	registerMiddleware(h)

	// add a ping route to test
	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, utils.H{"ping": "pong"})
	})

	h.GET("/about", middleware.Auth(), func(c context.Context, ctx *app.RequestContext) {
		ctx.HTML(consts.StatusOK, "about", utils.H{
			"title": "About",
		})
	})

	h.GET("/sign-in", middleware.AuthLogin(), func(c context.Context, ctx *app.RequestContext) {
		resp := myutils.WarpResponse(c, ctx, utils.H{
			"title": "Sign In",
			"from":  ctx.Request.Header.Get("Referer"),
		})

		if ctx.Query("from") != "" {
			resp["from"] = ctx.Query("from")
		}

		ctx.HTML(consts.StatusOK, "sign-in", resp)
	})

	h.GET("/sign-up", middleware.AuthLogin(), func(c context.Context, ctx *app.RequestContext) {
		resp := myutils.WarpResponse(c, ctx, utils.H{
			"title": "Sign Up",
			"from":  ctx.Query("from"),
		})
		ctx.HTML(consts.StatusOK, "sign-up", resp)
	})

	router.GeneratedRegister(h)
	h.LoadHTMLGlob("template/*")
	h.Static("/static", "./")

	h.Spin()
}

func registerMiddleware(h *server.Hertz) {
	store, _ := redis.NewStore(10, "tcp",
		conf.GetConf().Redis.Address, "",
		[]byte(os.Getenv("SESSION_SECRET")))
	h.Use(sessions.New("mymall-go-shop", store))

	// log
	logger := hertzlogrus.NewLogger()
	hlog.SetLogger(logger)
	hlog.SetLevel(conf.LogLevel())
	asyncWriter := &zapcore.BufferedWriteSyncer{
		WS: zapcore.AddSync(&lumberjack.Logger{
			Filename:   conf.GetConf().Hertz.LogFileName,
			MaxSize:    conf.GetConf().Hertz.LogMaxSize,
			MaxBackups: conf.GetConf().Hertz.LogMaxBackups,
			MaxAge:     conf.GetConf().Hertz.LogMaxAge,
		}),
		FlushInterval: time.Minute,
	}
	hlog.SetOutput(asyncWriter)
	h.OnShutdown = append(h.OnShutdown, func(ctx context.Context) {
		asyncWriter.Sync()
	})

	// pprof
	if conf.GetConf().Hertz.EnablePprof {
		pprof.Register(h)
	}

	// gzip
	if conf.GetConf().Hertz.EnableGzip {
		h.Use(gzip.Gzip(gzip.DefaultCompression))
	}

	// access log
	if conf.GetConf().Hertz.EnableAccessLog {
		h.Use(accesslog.New())
	}

	// recovery
	h.Use(recovery.Recovery())

	// cores
	h.Use(cors.Default())

	middleware.Register(h)
}
