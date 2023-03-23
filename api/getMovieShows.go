package api

import (
	DB "example/backend/database"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMovieShows(c *gin.Context) {
	var mId = c.Param("mId")
	var date = c.Param("date")

	db := DB.Connect()
	fmt.Printf("Great !!! we are connected to a browser\n")
	if c.Request.Method != "GET" {
		c.IndentedJSON(http.StatusMethodNotAllowed, gin.H{"success": false, "message": "This api is not valid. "})
		return
	}
	row, err := db.Query("select Id,Time,Seats,Date,MovieId,TheatreId,Screen from shows where MovieId = $1 and Date = $2 order by Time,TheatreId", mId, date)
	if err != nil {
		c.IndentedJSON(http.StatusMethodNotAllowed, gin.H{"success": false, "message": "Some issue while fetching querying SQL."})
		return
	}
	defer func() {
		row.Close()
		db.Close()
	}()

	var showList []Show
	for row.Next() {
		var tmpShow Show
		row.Scan(&tmpShow.Id, &tmpShow.Time, &tmpShow.Seats, &tmpShow.Date, &tmpShow.MovieId, &tmpShow.TheatreId, &tmpShow.Screen) //, &tmpShow.Screen)
		if err != nil {
			log.Println(err)
			c.IndentedJSON(http.StatusMethodNotAllowed, gin.H{"success": false, "message": "Error in scanning the row."})
			return
		}
		showList = append(showList, tmpShow)
	}

	if len(showList) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"success": false, "message": "Hey not found the data!!"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"success": true, "showList": showList})
}
