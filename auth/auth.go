package auth

import "github.com/gin-gonic/gin"

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		// 这里可以添加具体的 token 验证逻辑
		if token != "valid_token" {
			c.AbortWithStatusJSON(403, gin.H{"error": "Forbidden"})
			return
		}

		// 如果验证通过，继续处理请求
		c.Next()
	}
}
