package app

import (
	"github.com/DoNewsCode/core/ginmw"
	"github.com/gin-gonic/gin"
	"github.com/nfangxu/core-skeleton/app/handlers"
	"net/http"
)

type GinTransport struct {
	http.Handler
}

func NewGinTransport(handlers handlers.Handlers) GinTransport {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(ginmw.WithContext())
	r.Use(gin.Recovery())
	routes(r, handlers)
	return GinTransport{Handler: r}
}

func routes(gin *gin.Engine, handlers handlers.Handlers) {
	gin.GET("/users", handlers.Users.GetOne)
}
