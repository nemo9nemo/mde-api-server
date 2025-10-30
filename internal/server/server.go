package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/nemo9nemo/mds-api-server/internal/common"

	_ "github.com/nemo9nemo/mds-api-server/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Start() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to NDE API Server",
		})
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	addr := fmt.Sprintf(":%d", common.Cfg.Server.Port)
	fmt.Printf("[OK] Server running on %s\n", addr)
	r.Run(addr)
}
