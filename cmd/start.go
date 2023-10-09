package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/krau/bililink-converter/config"
	_ "github.com/krau/bililink-converter/docs"
	"github.com/krau/bililink-converter/handlers"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Start starts the server
func Start() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1", "::1"})
	r.GET("/bv2av/:bv", handlers.Bv2avEndpoint)
	r.GET("/av2bv/:av", handlers.Av2bvEndpoint)
	r.GET("/b23/:b23code", handlers.B23ConvertEndpoint)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(config.C.Address)
}
