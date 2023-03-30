package movie

import (
	st "example/backend/api/struct"
	DB "example/backend/database"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetMovieList(c *gin.Context) {
	var getLang []string = c.QueryArray("lang[]")
	var loc string = c.Query("loc")
	var lang string
	var isAll bool = false
	var len int = len(getLang)

	for i, t := range getLang {
		if t == "All" {
			isAll = true
		}
		if i != len-1 {
			lang += "'" + strings.Trim(t, "") + "',"
		} else {
			lang += "'" + strings.Trim(t, "") + "'"
		}
	}

	db := DB.Connect()
	fmt.Printf("Great !!! we are connected to a browser Guiz\n")
	if c.Request.Method != "GET" {
		c.IndentedJSON(http.StatusMethodNotAllowed, gin.H{"success": false, "message": "This api is not valid. "})
		return
	}
	row, err := db.Query("select Id, Name, Language, Image , HeadImage, Tags,Comment from movies where (Language in ("+lang+") or $1 = true) and Id in (select distinct(s.MovieId) from theatre t,shows s where s.TheatreId = t.Id and t.City = $2);", isAll, loc)
	if err != nil {
		c.IndentedJSON(http.StatusMethodNotAllowed, gin.H{"success": false, "message": "Some issue while fetching querying SQL.", "Error": err})
		return
	}
	defer func() {
		row.Close()
		db.Close()
	}()

	var movieList []st.Movie
	for row.Next() {
		var tmpMovie st.Movie
		var tags string
		row.Scan(&tmpMovie.Id, &tmpMovie.Name, &tmpMovie.Language, &tmpMovie.Image, &tmpMovie.HeadImage, tags, &tmpMovie.Comment)
		if err != nil {
			log.Println(err)
			c.IndentedJSON(http.StatusMethodNotAllowed, gin.H{"success": false, "message": "Error in scanning the row.", "Error": err})
			return
		}
		tmpMovie.Tags = strings.Split(tags, ":")
		movieList = append(movieList, tmpMovie)
	}
	c.IndentedJSON(http.StatusOK, gin.H{"success": true, "movieList": movieList})
}
