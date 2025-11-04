package server

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nemo9nemo/mds-api-server/internal/auth"
	"github.com/nemo9nemo/mds-api-server/internal/common"

	_ "github.com/nemo9nemo/mds-api-server/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Start() {
	r := gin.Default()

	// CORS 설정 추가
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // React 개발 서버 주소
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to NDE API Server",
		})
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/login", auth.LoginHandler)

	addr := fmt.Sprintf(":%d", common.Cfg.Server.Port)
	fmt.Printf("[OK] Server running on %s\n", addr)
	// r := gin.Default()

	fmt.Println("[OK] Server running on :8080")
	r.Run(":8080")
	// r.Run(addr)
}
