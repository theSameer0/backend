package theatre

import (
	"example/backend/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTheatreDetail(c *gin.Context) {
	var tId = c.Param("tId")
	// db := DB.Connect()
	fmt.Printf("Great !!! we are connected to a browser\n")
	if c.Request.Method != "GET" {
		c.IndentedJSON(http.StatusMethodNotAllowed, gin.H{"success": false, "message": "This api is not valid. "})
		return
	}
	// row1 := db.QueryRow("select Id,Name,Image,Location,City,Screen from theatre where Id = $1", tId)
	// var theatreData INDENT.Theatre
	var theatre model.Theatre
	model.DB.Model(&model.Theatre{}).Where("Id = ?", tId).Find(&theatre)

	// err1 := row1.Scan(&theatreData.Id, &theatreData.Name, &theatreData.Image, &theatreData.Location, &theatreData.City, &theatreData.Screen)
	c.JSON(http.StatusOK, gin.H{"success": true, "theatreData": theatre})
}
