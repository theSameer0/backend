package movie

import (
	INDENT "example/backend/api/struct"
	"example/backend/database"
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
	var movieList []INDENT.Movie
	if err := c.BindJSON(&movieList); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"success": false, "message": "cant bind the input data."})
		return
	}
	var tags []string
	var messages []INDENT.Message
	var ids []int
	var insert string
	for i, t := range movieList {
		ids = append(ids, -1)
		tags = append(tags, "")
		for ii, tt := range movieList[i].Tags {
			tags[i] += tt
			if ii != len(movieList[i].Tags)-1 {
				tags[i] += ":"
			}
		}

		if check(t, tags[i]) {
			messages = append(messages, INDENT.Message{Message: "Unsuccessful Some values are Empty. "})
		} else {
			ids[i] = 0
			insert += "("
			insert += "'" + t.Name + "','" + t.Language + "','" + t.Image + "','" + t.HeadImage + "','" + tags[i] + "','" + t.Comment + "'),"
			messages = append(messages, INDENT.Message{Message: ""})
		}
	}
	if len(insert) != 0 {
		insert = insert[:len(insert)-1]
	}
	row, err := db.Query("Insert into movies (Name,Language,Image,HeadImage,Tags,Comment) values " + insert + " RETURNING Id;")

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"successs": false, "message": "Error in fetching the sql query.", "error": err})
		return
	}
	defer func() {
		row.Close()
		db.Close()
	}()
	var count int = 0
	for row.Next() {
		var tmp INDENT.Movie
		err := row.Scan(&tmp.Id)
		for {
			if ids[count] == 0 {
				break
			}
			count++
		}
		if err != nil {
			messages[count].Message += "Unsuccessfully Inserting this data."
		} else {
			messages[count].Message = "Successfully appended this Id."
			ids[count] = tmp.Id
		}
		count++
	}
	c.IndentedJSON(http.StatusOK, gin.H{"success": true, "Ids": ids, "Messages": messages})
}

func check(t INDENT.Movie, tags string) bool {
	if strings.Trim(t.Name, " ") == "" || strings.Trim(t.Language, " ") == "" || strings.Trim(t.Image, " ") == "" ||
		strings.Trim(t.Image, " ") == "" || strings.Trim(t.HeadImage, " ") == "" || strings.Trim(tags, " ") == "" || strings.Trim(t.Comment, " ") == "" {
		return true
	}
	return false
}
