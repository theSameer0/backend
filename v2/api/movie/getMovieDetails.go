package movie

import (
	"encoding/json"
	"example/backend/model"
	INDENT "example/backend/v2/api/struct"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetMovieDetail(c *gin.Context) {
	var mId = c.Param("mId")
	fmt.Printf("Great !!! we are connected to a browser\n")
	if c.Request.Method != "GET" {
		c.IndentedJSON(http.StatusMethodNotAllowed, gin.H{"success": false, "message": "This api is not valid. "})
		return
	}
	var movies model.Movie
	var data INDENT.MovieData
	model.DB.Model(&movies).Where("Id = ?", mId).Find(&movies)
	data.Id = movies.Id
	// data.Name = movies.Name
	data.HeadImage = movies.Headimage
	data.Tags = strings.Split(movies.Tags, ":")
	data.Comment = movies.Comment
	data.Name = getName(movies.Name)
	c.JSON(http.StatusOK, gin.H{"success": true, "movieData": data})
}

func getName(name string) string {
	url := `http://localhost:50000/name/` + name
	req, err := http.Get(url)
	if err != nil {
		fmt.Printf("Some error occured %v\n", err)
		return ""
	}
	defer req.Body.Close()
	var res map[string]string
	json.NewDecoder(req.Body).Decode(&res)
	return res["name"]
}
