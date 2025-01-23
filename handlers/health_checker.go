package handlers
import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HealtCheckHandler() func(c *gin.Context) {
	
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ok":true})
	}
}