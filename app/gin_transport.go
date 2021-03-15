package app

import (
	ginmw "github.com/DoNewsCode/core-gin/mw"
	"github.com/DoNewsCode/core-starter/app/handlers"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GinTransport struct {
	http.Handler
}

func NewGinTransport(handlers handlers.Handlers) GinTransport {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(ginmw.Context())
	r.Use(gin.Recovery())
	routes(r, handlers)
	return GinTransport{Handler: r}
}

func routes(gin *gin.Engine, handlers handlers.Handlers) {
	gin.GET("/users", handlers.Users.GetOne)
}
