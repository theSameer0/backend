package api

import (
	DB "example/sameer_mbs/src/server/database"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetTicket(c *gin.Context) {
	var ticket Ticket
	if err := c.BindJSON(&ticket); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"success": false, "message": "Some Error Occurred."})
		return
	}
	db := DB.Connect()
	fmt.Printf("Great !!! we are connected to a browser\n")
	if c.Request.Method != "POST" {
		fmt.Print("This api is not valid. ")
		c.IndentedJSON(http.StatusMethodNotAllowed, gin.H{"success": false, "message": "This api is not valid. "})
		return
	}
	row, err := db.Query("Update ticket set Id=$1,Date=$2,Time=$3,Seats=$4,SeatCount=$5,Screen=$6,MovieId=$7,TheatreId=$8,ShowId=$9 where Screen=4",
		ticket.Id, ticket.Date, ticket.Time, ticket.Seats, ticket.SeatCount, ticket.Screen, ticket.MovieId, ticket.TheatreId, ticket.ShowId)
	if err != nil {
		fmt.Print("The query is not sent properly.")
		c.IndentedJSON(http.StatusNotFound, gin.H{"success": false, "message": "The query is not sent properly."})
		return
	}
	defer func() {
		row.Close()
		db.Close()
	}()

	fmt.Printf("%v", ticket)
	c.IndentedJSON(http.StatusOK, gin.H{"success": true, "ticket": ticket})
}
