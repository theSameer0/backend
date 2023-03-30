package movie

import (
	INDENT "example/backend/api/struct"
	DB "example/backend/database"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMovieTheatres(c *gin.Context) {
	var mId = c.Param("mId")
	var date = c.Param("date")

	db := DB.Connect()
	fmt.Printf("Great !!! we are connected to a browser\n")
	if c.Request.Method != "GET" {
		c.IndentedJSON(http.StatusMethodNotAllowed, gin.H{"success": false, "message": "This api is not valid. "})
		return
	}

	row2, err2 := db.Query("select s.Id,s.Time,s.Seats,s.Date,s.MovieId,s.TheatreId,s.Screen,t.Name,t.Location from shows s,theatre t where s.TheatreId = t.Id and s.MovieId = $1 and s.Date = $2 order by Time,TheatreId", mId, date)
	if err2 != nil {
		c.IndentedJSON(http.StatusMethodNotAllowed, gin.H{"success": false, "message": "Some issue while fetching querying SQL.", "Error": err2.Error()})
		return
	}
	defer func() {
		row2.Close()
		db.Close()
	}()

	var theatreList []INDENT.MovieShow
	for row2.Next() {
		var tmpShow INDENT.Show
		var name, location string
		err := row2.Scan(&tmpShow.Id, &tmpShow.Time, &tmpShow.Seats, &tmpShow.Date, &tmpShow.MovieId, &tmpShow.TheatreId, &tmpShow.Screen, &name, &location)
		if err != nil {
			log.Println(err)
			c.IndentedJSON(http.StatusMethodNotAllowed, gin.H{"success": false, "message": "Error in scanning the row.", "Error": err.Error()})
			return
		}
		var isFound bool = false
		for i, t := range theatreList {
			if t.TheatreId == tmpShow.TheatreId {
				isFound = true
				theatreList[i].ShowTime = append(theatreList[i].ShowTime, INDENT.ShowTime{
					ShowId: tmpShow.Id,
					Time:   tmpShow.Time,
					Seat:   tmpShow.Seats,
				})
			}
		}
		if !isFound {
			theatreList = append(theatreList, INDENT.MovieShow{
				TheatreId:       tmpShow.TheatreId,
				TheatreName:     name,
				TheatreLocation: location,
				ShowTime: []INDENT.ShowTime{
					{
						ShowId: tmpShow.Id,
						Time:   tmpShow.Time,
						Seat:   tmpShow.Seats,
					},
				},
				ShowDate:   tmpShow.Date,
				ShowScreen: tmpShow.Screen,
			})
		}
	}

	if len(theatreList) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"success": false, "message": "Hey not found the data!!"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"success": true, "theatreList": theatreList})
}
