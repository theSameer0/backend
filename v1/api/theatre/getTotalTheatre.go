package theatre

import (
	"example/backend/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTotalTheatre(c *gin.Context) {
	if c.Request.Method != "GET" {
		c.IndentedJSON(http.StatusMethodNotAllowed, gin.H{"success": false, "message": "This API is not allowed."})
		return
	}
	var num int
	model.DB.Model(&model.Theatre{}).Select("Distinct(Id)").Find(&num)
	c.IndentedJSON(http.StatusOK, gin.H{"success": true, "totalTheatre": num})
}
