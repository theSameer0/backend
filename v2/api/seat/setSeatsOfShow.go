package seat

import (
	"example/backend/model"
	INDENT "example/backend/v2/api/struct"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetSeatsOfShow(c *gin.Context) {
	var mId = c.Param("mId")
	var tId = c.Param("tId")
	var time = c.Param("time")
	var date = c.Param("date")
	var showId int
	var seat INDENT.GetSeat

	if err := c.BindJSON(&seat); err != nil {
		fmt.Printf("error: %T", err)
		return
	}

	// db := DB.Connect()
	fmt.Printf("Great !!! we are connected to a browser Guiz\n")
	if c.Request.Method != "POST" {
		c.IndentedJSON(http.StatusMethodNotAllowed, gin.H{"success": false, "message": "This api is not valid. "})
		return
	}
	model.DB.Model(&model.Show{}).Where("MovieId = ? and TheatreId = ? and Time = ? and Date = ?", mId, tId, time, date).Update("Seats", seat.Seats).Select("Id").Find(&showId)
	// row := db.QueryRow("update shows set Seats = $1 where Date = $2 and Time = $3 and MovieId = $4 and TheatreId = $5 RETURNING Id;", seat.Seats, date, time, mId, tId)

	// defer func() {
	// 	db.Close()
	// }()
	// var id int
	// err1 := row.Scan(&id)
	// if err1 != nil {
	// 	c.IndentedJSON(http.StatusNotFound, gin.H{"success": false, "message": "Scanning the Id.", "error": err1.Error()})
	// 	return
	// }

	// fmt.Printf("It Got Hit the End Point.%v", seat)
	c.IndentedJSON(http.StatusOK, gin.H{"success": true, "seat": seat, "showId": showId})
}
