package api

import (
	"example/sameer_mbs/src/server/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetLastIds(c *gin.Context) {
	var showId, theatreId, movieId string
	db := database.Connect()
	if c.Request.Method != "GET" {
		c.IndentedJSON(http.StatusMethodNotAllowed, gin.H{"success": false, "message": "Wrong Api Type."})
		return
	}
	row1 := db.QueryRow("select COUNT(Id) from shows;")
	defer func() {
		db.Close()
	}()
	row2 := db.QueryRow("select COUNT(Id) from theatre;")
	defer func() {
		db.Close()
	}()
	row3 := db.QueryRow("select COUNT(Id) from movies;")
	defer func() {
		db.Close()
	}()

	err1 := row1.Scan(&showId)
	if err1 != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"success": false, "message": "Some error while processing the query"})
		return
	}
	err2 := row2.Scan(&theatreId)
	if err2 != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"success": false, "message": "Some error while processing the query"})
		return
	}
	err3 := row3.Scan(&movieId)
	if err3 != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"success": false, "message": "Some error while processing the query"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"success": true, "ShowId": showId, "TheatreId": theatreId, "MovieId": movieId})
}
