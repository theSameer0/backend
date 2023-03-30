package movie

import (
	INDENT "example/backend/api/struct"
	DB "example/backend/database"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetMovieDetail(c *gin.Context) {
	var mId = c.Param("mId")

	db := DB.Connect()
	fmt.Printf("Great !!! we are connected to a browser\n")
	if c.Request.Method != "GET" {
		c.IndentedJSON(http.StatusMethodNotAllowed, gin.H{"success": false, "message": "This api is not valid. "})
		return
	}
	row1 := db.QueryRow("select Id,Name,HeadImage,Tags,Comment from movies where Id = $1", mId)
	var movieData INDENT.MovieData
	var tags string
	err1 := row1.Scan(&movieData.Id, &movieData.Name, &movieData.HeadImage, &tags, &movieData.Comment)
	if err1 != nil {
		c.IndentedJSON(http.StatusMethodNotAllowed, gin.H{"success": false, "message": "Some issue while fetching querying SQL.", "Error": err1.Error()})
		return
	}
	movieData.Tags = strings.Split(tags, ":")

	c.IndentedJSON(http.StatusOK, gin.H{"success": true, "movieData": movieData})
}
