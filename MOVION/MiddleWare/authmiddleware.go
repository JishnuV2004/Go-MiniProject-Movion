package middleware

import (
	"fmt"
	utils "movion/Utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func RBACmiddleware(Roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "missing token",
			})
			c.Abort()
			return
		}
		tokenStr := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))
		claims, err := utils.VerifyToken(tokenStr)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid or expired token",
			})
			c.Abort()
			return
		}

		if claims.Block {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":"user is blocked. access denied",
			})
			c.SetCookie("refresh_token", "", -1, "/", "", false, true)
			c.Abort()
			return 
		}
		
		fmt.Println("RBAC DEBUG → Email:", claims.Email, "| Role in token:", claims.Role, "| Allowed roles:", Roles)

		authorized := false
		authorized = utils.CheckRole(Roles, claims.Role)
		// for _, role := range Roles {
		// 	if claims.Role == role {
		// 		authorized = true
		// 		break
		// 	}
		// }
		if !authorized {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "access denied",
			})
			c.Abort()
			return
		}
		c.Set("userID", claims.UserID)
		c.Next()
	}
}
