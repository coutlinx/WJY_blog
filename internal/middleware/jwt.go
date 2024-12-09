package middleware

import (
	"blog/pkg"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		var code uint
		var claims *pkg.MemberClaims
		var result any
		var err error
		if result, err = pkg.NewToken().ParseToken(token); err != nil {
			zap.S().Info("token鉴权失败.", zap.Error(err))
			code = http.StatusMethodNotAllowed
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  "token鉴权失败",
			})
		}
		claims = result.(*pkg.MemberClaims)
		c.Set("user_id", claims.UserId)
		c.Next()
	}
}
