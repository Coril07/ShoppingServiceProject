package main

import (
	controllers "app/server/controllers"
	"fmt"

	"app/server/utils/middleware"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// main アプリケーションの入口
func main() {
	r := gin.Default()
	registerMiddlewares(r)
	controllers.RegisterHandlers(r)
	r.Run(":8080")

}

// ミドルウェアの登録
func registerMiddlewares(r *gin.Engine) {
	r.Use(
		gin.LoggerWithFormatter(
			func(param gin.LogFormatterParams) string {

				return fmt.Sprintf(
					"%s - [%s] \"%s %s %s %d %s %s\"\n",
					param.ClientIP,
					param.TimeStamp.Format(time.RFC3339),
					param.Method,
					param.Path,
					param.Request.Proto,
					param.StatusCode,
					param.Latency,
					param.ErrorMessage,
				)
			}))
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))
	r.Use(middleware.CORSMiddleware())
	r.Use(gin.Recovery())
}
