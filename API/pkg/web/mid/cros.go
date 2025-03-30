package mid

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetHeaderAllowCors(w *gin.ResponseWriter, origin string) {
	(*w).Header().Set("Access-Control-Allow-Origin", origin)
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func Cors(origin string) gin.HandlerFunc {
	return func(c *gin.Context) {
		SetHeaderAllowCors(&c.Writer, origin)

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
