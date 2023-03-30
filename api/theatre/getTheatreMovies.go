package theatre

import (
	INDENT "example/backend/api/struct"
	DB "example/backend/database"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetTheatreMovies(c *gin.Context) {
	var tId = c.Param("tId")
	var date = c.Param("date")

	db := DB.Connect()
	fmt.Printf("Great !!! we are connected to a browser\n")
	if c.Request.Method != "GET" {
		c.IndentedJSON(http.StatusMethodNotAllowed, gin.H{"success": false, "message": "This api is not valid. "})
		return
	}

	row2, err2 := db.Query("select s.Id,s.Time,s.Seats,s.Date,s.MovieId,s.TheatreId,s.Screen,m.Name,m.Language,m.Image,m.Tags from shows s,movies m where s.MovieId = m.Id and s.TheatreId = $1 and s.Date = $2 order by Time,TheatreId", tId, date)
	if err2 != nil {
		c.IndentedJSON(http.StatusMethodNotAllowed, gin.H{"success": false, "message": "Some issue while fetching querying SQL.", "Error": err2.Error()})
		return
	}
	defer func() {
		row2.Close()
		db.Close()
	}()

	var movieList []INDENT.TheatreShow
	for row2.Next() {
		var tmpShow INDENT.Show
		var name, language, image, tags string
		err := row2.Scan(&tmpShow.Id, &tmpShow.Time, &tmpShow.Seats, &tmpShow.Date, &tmpShow.MovieId, &tmpShow.TheatreId, &tmpShow.Screen, &name, &language, &image, &tags)
		if err != nil {
			log.Println(err)
			c.IndentedJSON(http.StatusMethodNotAllowed, gin.H{"success": false, "message": "Error in scanning the row.", "Error": err.Error()})
			return
		}
		var isFound bool = false
		for i, t := range movieList {
			if t.MovieId == tmpShow.MovieId {
				isFound = true
				movieList[i].ShowTime = append(movieList[i].ShowTime, INDENT.ShowTime{
					ShowId: tmpShow.Id,
					Time:   tmpShow.Time,
					Seat:   tmpShow.Seats,
				})
			}
		}
		if !isFound {
			movieList = append(movieList, INDENT.TheatreShow{
				MovieId:       tmpShow.MovieId,
				MovieName:     name,
				MovieLanguage: language,
				MovieImage:    image,
				MovieTags:     strings.Split(tags, ":"),
				ShowDate:      tmpShow.Date,
				ShowScreen:    tmpShow.Screen,
				ShowTime: []INDENT.ShowTime{
					{
						ShowId: tmpShow.Id,
						Time:   tmpShow.Time,
						Seat:   tmpShow.Seats,
					},
				},
			})
		}
	}

	if len(movieList) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"success": false, "message": "Hey not found the data!!"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"success": true, "movieList": movieList})
}
