package ticket

import (
	"example/backend/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTicket(c *gin.Context) {
	// db := DB.Connect()
	id := c.Param("id")
	fmt.Printf("Great !!! we are connected to a browser\n")
	if c.Request.Method != "GET" {
		c.IndentedJSON(http.StatusMethodNotAllowed, gin.H{"success": false, "message": "This api is not valid. "})
		return
	}
	var ticket model.Ticket
	model.DB.Model(&model.Ticket{}).Where("Id = ?", id).Take(&ticket)
	// row := db.QueryRow("SELECT Id, Date, Time, Seats, SeatCount, Screen, MovieId, TheatreId, ShowId FROM ticket")
	// defer func() {
	// 	db.Close()
	// }()
	// var ticket INDENT.Ticket
	// err := row.Scan(&ticket.Id, &ticket.Date, &ticket.Time, &ticket.Seats, &ticket.SeatCount, &ticket.Screen, &ticket.MovieId, &ticket.TheatreId, &ticket.ShowId)
	// if err != nil {
	// 	c.IndentedJSON(http.StatusNotFound, gin.H{"success": false, "message": "Error while Scanning the List.", "Error": err.Error()})
	// 	return
	// }
	c.IndentedJSON(http.StatusOK, gin.H{"success": true, "ticket": ticket})
}
