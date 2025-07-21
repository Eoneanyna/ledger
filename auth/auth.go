package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"ledger/conf"
	"net/http"
	"time"
)

const AuthUserIDKey = "user_id"

type CustomClaims struct {
	UserID int64 `json:"user_id"`
	jwt.StandardClaims
}

// GenerateToken 生成JWT token
func GenerateToken(userID int64) (string, error) {
	claims := CustomClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(), //设置过期时间
			IssuedAt:  time.Now().Unix(),
		},
	}
	//此处是对称加密，可以修改为es256非对称加密。生成公私钥，此处加载私钥文件进行生成token，验证时使用公钥。
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(conf.Conf.Server.JwtSecret))
}

// 解析token
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		// 解析token并验证
		claims := &CustomClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			//非对称加密return公钥
			return []byte(conf.Conf.Server.JwtSecret), nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token"})
			return
		}

		// 将用户ID存入上下文，供后续处理函数使用
		c.Set(AuthUserIDKey, claims.UserID)
		c.Next()
	}
}
