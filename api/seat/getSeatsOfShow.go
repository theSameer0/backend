package seat

import (
	DB "example/backend/database"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSeatsOfShow(c *gin.Context) {
	var mId = c.Param("mId")
	var tId = c.Param("tId")
	var time = c.Param("time")
	var date = c.Param("date")
	var seat string
	db := DB.Connect()
	fmt.Printf("Great !!! we are connected to a browser\n")
	if c.Request.Method != "GET" {
		c.IndentedJSON(http.StatusMethodNotAllowed, gin.H{"success": false, "message": "This api is not valid. "})
		return
	}
	row := db.QueryRow("select Seats from shows where Date = $1 and Time = $2 and MovieId = $3 and TheatreId = $4", date, time, mId, tId)
	defer func() {
		db.Close()
	}()
	err := row.Scan(&seat)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"success": false, "message": "Error in scanning the rows."})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"success": true, "seat": seat})
}
