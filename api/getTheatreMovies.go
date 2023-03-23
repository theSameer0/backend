package api

import (
	DB "example/backend/database"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetTheatreMovies(c *gin.Context) {
	var tId string = c.Param("tId")
	var date string = c.Param("date")

	db := DB.Connect()
	fmt.Printf("We have landed to the Application!!")
	if c.Request.Method != "GET" {
		c.IndentedJSON(http.StatusMethodNotAllowed, gin.H{"success": false, "message": "This api is not valid. "})
		return
	}

	row, err := db.Query("select m.Id, m.Name, m.Tags, m.Image, s.TheatreId, s.Id, s.Date , s.Time , s.Seats from movies m,shows s where m.Id=s.MovieId and s.TheatreId = $1 and s.Date = $2;", tId, date)
	if err != nil {
		c.IndentedJSON(http.StatusMethodNotAllowed, gin.H{"success": false, "message": "Some issue while fetching querying SQL."})
		return
	}
	defer func() {
		row.Close()
		db.Close()
	}()

	var rawMovieList []rawMovieTime
	for row.Next() {
		var tmpMovie rawMovieTime
		var tags string

		row.Scan(&tmpMovie.Id, &tmpMovie.Name, &tags, &tmpMovie.Image, &tmpMovie.TheatreId, &tmpMovie.ShowId, &tmpMovie.Date, &tmpMovie.Time, &tmpMovie.Seats) //, &tmpShow.Screen)
		if err != nil {
			log.Println(err)
			c.IndentedJSON(http.StatusMethodNotAllowed, gin.H{"success": false, "message": "Error in scanning the row."})
			return
		}
		tmpMovie.Tags = strings.Split(tags, ":")
		rawMovieList = append(rawMovieList, tmpMovie)
	}
	var theatreMovieList []MovieTime
	for _, t := range rawMovieList {
		var found bool = false
		for i, tt := range theatreMovieList {
			if tt.Id == t.Id {
				found = true
				var tmpSeat = Seat{
					Timing: t.Time,
					Seats:  t.Seats,
				}
				theatreMovieList[i].Time = append(theatreMovieList[i].Time, tmpSeat)
			}
		}
		if !found {
			var tmpSeat = Seat{
				Timing: t.Time,
				Seats:  t.Seats,
			}
			var tmpList = MovieTime{
				Id:        t.Id,
				Name:      t.Name,
				Tags:      t.Tags,
				Image:     t.Image,
				TheatreId: t.TheatreId,
				ShowId:    t.ShowId,
				Date:      t.Date,
				Time:      []Seat{tmpSeat},
			}
			theatreMovieList = append(theatreMovieList, tmpList)
		}
	}
	c.IndentedJSON(http.StatusOK, gin.H{"success": true, "theatreMovieList": theatreMovieList})
}
