package ticket

import (
	"example/backend/model"
	INDENT "example/backend/v1/api/struct"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTicketList(c *gin.Context) {
	if c.Request.Method != "GET" {
		c.IndentedJSON(http.StatusNotFound, gin.H{"success": false, "message": "This api request is valid."})
		return
	}
	var data []INDENT.TicketList

	model.DB.Model(&model.Ticket{}).Order("timestamp DESC").Find(&data)
	c.IndentedJSON(http.StatusOK, gin.H{"success": true, "ticketList": data})
}
