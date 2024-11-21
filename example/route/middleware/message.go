package middleware

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

func PrintMessage(message string) gin.HandlerFunc {
	return func(c *gin.Context) {
		slog.Info(message)
		c.Next()
	}
}
