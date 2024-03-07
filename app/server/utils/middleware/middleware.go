package middleware

import (
	"app/server/utils/api_helper"
	jwtHelper "app/server/utils/jwt"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 管理员授权
func AuthAdminMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("Token")
		fmt.Printf("token: %v\n", token)
		if err != nil {
			api_helper.HandleError(c, err)
			c.Abort()
			return
		}
		if token != "" {
			decodedClaims := jwtHelper.VerifyToken(token, secretKey)
			if decodedClaims != nil && decodedClaims.IsAdmin {
				c.Next()
				return
			}
			c.JSON(http.StatusForbidden, gin.H{"error": "你没有权限访问!"})
			c.Abort()
			return
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "你没有授权!"})
			c.Abort()
		}
	}
}

// 用户授权
func AuthUserMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("Token")
		if err != nil {
			api_helper.HandleError(c, err)
			c.Abort()
			return
		}
		if token != "" {
			decodedClaims := jwtHelper.VerifyToken(token, secretKey)
			if decodedClaims != nil {
				c.Set("userId", decodedClaims.UserId)
				c.Set("username", decodedClaims.Username)
				c.Set("isadmin", decodedClaims.IsAdmin)
				c.Next()
				return
			}
			c.JSON(http.StatusForbidden, gin.H{"error": "你没有权限访问!"})
			c.Abort()
			return
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "你没授权!"})
			c.Abort()
		}
	}
}

// CORSMiddleware 处理 CORS 中间件
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 允许所有来源访问，你也可以根据需要设置特定的域名
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		// 设置其他 CORS 相关的头部
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		// 允许发送 cookies
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		// 处理预检请求（OPTIONS 请求）
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		// 继续处理请求
		c.Next()
	}
}
