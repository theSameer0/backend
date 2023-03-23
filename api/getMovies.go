package api

import (
	DB "example/sameer_mbs/src/server/database"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetMovies(c *gin.Context) {
	db := DB.Connect()
	fmt.Printf("Great !!! we are connected to a browser\n")
	if c.Request.Method != "GET" {
		c.IndentedJSON(http.StatusMethodNotAllowed, gin.H{"success": false, "message": "This api is not valid. "})
		return
	}
	row, err := db.Query("select Id,Name,Language,Image,HeadImage,Tags,Comment from movies")
	if err != nil {
		c.IndentedJSON(http.StatusMethodNotAllowed, gin.H{"success": false, "message": "Some issue while fetching querying SQL."})
		return
	}
	defer func() {
		row.Close()
		db.Close()
	}()
	var movieList []Movie
	for row.Next() {
		var tmpMovie Movie
		var tags string
		err := row.Scan(&tmpMovie.Id, &tmpMovie.Name, &tmpMovie.Language, &tmpMovie.Image, &tmpMovie.HeadImage, &tags, &tmpMovie.Comment)
		if err != nil {
			log.Println(err)
			c.IndentedJSON(http.StatusMethodNotAllowed, gin.H{"success": false, "message": "Error in scanning the row."})
			return
		}
		tmpMovie.Tags = strings.Split(tags, ":")
		movieList = append(movieList, tmpMovie)
	}
	c.IndentedJSON(http.StatusOK, gin.H{"sucess": true, "movieList": movieList})
	db.Close()
}
