package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"web_app/controller"
	"web_app/logger"
	"web_app/middlewares"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true), middlewares.RateLimitMiddleware(2*time.Second, 1))

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	v1 := r.Group("api/v1")

	// register route service
	v1.POST("/signup", controller.SignUpHandler)
	v1.POST("/login", controller.LoginHandler)

	v1.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	v1.Use(middlewares.JWTAuthMiddleware())

	//v1.GET("/ping", middlewares.JWTAuthMiddleware(), func(c *gin.Context) {
	//	c.String(http.StatusOK, "pong")
	//})

	{
		v1.GET("/community", controller.CommunityHandler)
		v1.GET("/community/:id", controller.CommunityDetailHandler)

		v1.POST("/post", controller.CreatePostHandler)
		v1.GET("/post/:id", controller.GetPostDetailHandler)
		v1.GET("/posts", controller.GetPostListHandler)
		v1.GET("/post2/", controller.GetPostListHandler2)

		v1.POST("/vote", controller.PostVoteHandler)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
