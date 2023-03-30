package ticket

import (
	INDENT "example/backend/api/struct"
	DB "example/backend/database"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetTicket(c *gin.Context) {
	var ticket INDENT.Ticket
	if err := c.BindJSON(&ticket); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"success": false, "message": "Some Error Occurred.", "Error": err})
		return
	}
	db := DB.Connect()
	db.QueryRow("delete from ticket where 1 = 1;")

	fmt.Printf("Great !!! we are connected to a browser\n")
	if c.Request.Method != "POST" {
		fmt.Print("This api is not valid. ")
		c.IndentedJSON(http.StatusMethodNotAllowed, gin.H{"success": false, "message": "This api is not valid. "})
		return
	}
	row, err := db.Query("insert into ticket (Id,Date,Time,Seats,SeatCount,Screen,MovieId,TheatreId,ShowId) values ($1,$2,$3,$4,$5,$6,$7,$8,$9);",
		ticket.Id, ticket.Date, ticket.Time, ticket.Seats, ticket.SeatCount, ticket.Screen, ticket.MovieId, ticket.TheatreId, ticket.ShowId)
	if err != nil {
		fmt.Print("The query is not sent properly.")
		c.IndentedJSON(http.StatusNotFound, gin.H{"success": false, "message": "The query is not sent properly.", "Error": err})
		return
	}
	defer func() {
		row.Close()
		db.Close()
	}()

	fmt.Printf("%v", ticket)
	c.IndentedJSON(http.StatusOK, gin.H{"success": true, "ticket": ticket})
}
