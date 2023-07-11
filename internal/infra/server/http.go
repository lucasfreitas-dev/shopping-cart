package server

import (
	"shopping-cart/internal/config"
	"shopping-cart/internal/handlers/carthandler"

	"fmt"

	"github.com/gin-gonic/gin"
)

type server struct {
	ginSrv *gin.Engine
	port   string
}

func NewServer(cfg config.Server) *server {

	return &server{
		ginSrv: gin.New(),
		port:   cfg.Port,
	}

}

func (srv *server) NewRouter(carthandler *carthandler.HTTPHandler) {

	srv.ginSrv.Use(gin.Logger())
	srv.ginSrv.Use(gin.Recovery())
	srv.ginSrv.GET("/shopping-carts", carthandler.Get)
	srv.ginSrv.POST("/shopping-carts/items", carthandler.AddItem)
	srv.ginSrv.DELETE("/shopping-carts/items/:itemID", carthandler.RemoveItem)

}

func (srv *server) Run() error {
	return srv.ginSrv.Run(fmt.Sprintf(":%v", srv.port))
}
