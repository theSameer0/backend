package movie

import (
	"example/backend/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTotalMovie(c *gin.Context) {
	if c.Request.Method != "GET" {
		c.IndentedJSON(http.StatusMethodNotAllowed, gin.H{"success": false, "message": "This api is not allowed."})
		return
	}
	var num int
	model.DB.Model(&model.Movie{}).Select("Count(Id) as num").Find(&num)
	c.IndentedJSON(http.StatusOK, gin.H{"success": true, "totalMovie": num})
}
