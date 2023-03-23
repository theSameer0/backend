package api

import (
	"example/sameer_mbs/src/server/database"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func InsertMovies(c *gin.Context) {
	db := database.Connect()
	if c.Request.Method != "POST" {
		c.IndentedJSON(http.StatusMethodNotAllowed, gin.H{"success": false, "message": "This type of api is not allowed"})
		return
	}
	var movieList []Movie
	if err := c.BindJSON(&movieList); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"success": false, "message": "cant bind the input data."})
		return
	}
	var tags []string
	var ids []string
	var insert string
	for i, t := range movieList {
		ids = append(ids, t.Id)
		insert += "("
		insert += "'" + t.Id + "','" + t.Name + "','" + t.Language + "','" + t.Image + "','" +
			t.HeadImage + "','"
		tags = append(tags, "")
		for ii, tt := range movieList[i].Tags {
			tags[i] += tt
			if ii != len(movieList[i].Tags)-1 {
				tags[i] += ":"
			}
		}
		insert += tags[i] + "','" + t.Comment + "')"
		if i != len(movieList)-1 {
			insert += ","
		}

		if check(t, tags[i]) {
			c.IndentedJSON(http.StatusMethodNotAllowed, gin.H{"successs": false, "message": "Some fields are empty."})
			return
		}
	}
	row, err := db.Query("Insert into movies (Id,Name,Language,Image,HeadImage,Tags,Comment) values " + insert + ";")
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"successs": false, "message": "Error in fetching the sql query."})
		return
	}
	defer func() {
		row.Close()
		db.Close()
	}()
	c.IndentedJSON(http.StatusOK, gin.H{"success": true, "Ids": ids})
}

func check(t Movie, tags string) bool {
	if strings.Trim(t.Id, " ") == "" || strings.Trim(t.Name, " ") == "" || strings.Trim(t.Language, " ") == "" || strings.Trim(t.Image, " ") == "" ||
		strings.Trim(t.Image, " ") == "" || strings.Trim(t.HeadImage, " ") == "" || strings.Trim(tags, " ") == "" || strings.Trim(t.Comment, " ") == "" {
		return true
	}
	return false
}
