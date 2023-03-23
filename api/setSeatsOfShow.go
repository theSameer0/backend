package api

import (
	DB "example/sameer_mbs/src/server/database"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetSeatsOfShow(c *gin.Context) {
	var mId = c.Param("mId")
	var tId = c.Param("tId")
	var time = c.Param("time")
	var date = c.Param("date")
	var seat getSeat

	if err := c.BindJSON(&seat); err != nil {
		fmt.Printf("error: %T", err)
		return
	}

	db := DB.Connect()
	fmt.Printf("Great !!! we are connected to a browser Guiz\n")
	if c.Request.Method != "POST" {
		c.IndentedJSON(http.StatusMethodNotAllowed, gin.H{"success": false, "message": "This api is not valid. "})
		return
	}
	row, err := db.Query("update shows set Seats = $1 where Date = $2 and Time = $3 and MovieId = $4 and TheatreId = $5 ", seat.Seats, date, time, mId, tId)
	if err != nil {
		c.IndentedJSON(http.StatusMethodNotAllowed, gin.H{"success": false, "message": "Some issue while posting query to sqlDB."})
		return
	}
	defer func() {
		row.Close()
		db.Close()
	}()

	fmt.Printf("It Got Hit the End Point.%v", seat)
	c.IndentedJSON(http.StatusOK, gin.H{"success": true, "seat": seat})
}
