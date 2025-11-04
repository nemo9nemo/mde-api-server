package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 로그인 요청 파라미터
type LoginRequest struct {
	ID string `json:"id" binding:"required"`
	PW string `json:"pw" binding:"required"`
}

// 로그인 핸들러
func LoginHandler(c *gin.Context) {
	var req LoginRequest

	// 검증
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errorMessage": "형식이 잘못되었습니다.", "isSuccess": false})
	}

	tokenPair, err := LoginService(req.ID, req.PW)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"errorMessage": err.Error(), "isSuccess": false})
	}

	c.JSON(http.StatusOK, gin.H{
		"isSuccess":     true,
		"message":       "로그인 성공",
		"access_token":  tokenPair.AccessToken,
		"refresh_token": tokenPair.RefreshToken,
	})
}
