package movie

import (
	"example/backend/model"
	INDENT "example/backend/v1/api/struct"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func InsertMovies(c *gin.Context) {
	// db := database.Connect()
	if c.Request.Method != "POST" {
		c.IndentedJSON(http.StatusMethodNotAllowed, gin.H{"success": false, "message": "This type of api is not allowed"})
		return
	}
	var data []INDENT.Movie
	if err := c.BindJSON(&data); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"success": false, "message": "cant bind the input data."})
		return
	}

	var tags []string
	var messages []INDENT.Message
	var ids []int
	var movieList []model.Movie
	// var insert string
	for i, t := range data {
		ids = append(ids, -1)
		tags = append(tags, "")
		for ii, tt := range t.Tags {
			tags[i] += tt
			if ii != len(t.Tags)-1 {
				tags[i] += ":"
			}
		}
		if check(t, tags[i]) {
			messages = append(messages, INDENT.Message{Message: "Unsuccessful Some values are Empty. "})
		} else {
			ids[i] = 0
			messages = append(messages, INDENT.Message{Message: ""})
			movieList = append(movieList, model.Movie{
				Name:      t.Name,
				Language:  t.Language,
				Image:     t.Image,
				Headimage: t.HeadImage,
				Tags:      tags[i],
				Comment:   t.Comment,
			})
		}
	}

	model.DB.Select("Name", "Language", "Image", "Headimage", "Tags", "Comment").Create(&movieList)
	var cnt int = 0
	for _, movie := range movieList {
		for {
			if ids[cnt] == 0 {
				break
			}
			cnt++
		}
		ids[cnt] = movie.Id
		messages[cnt].Message = "Data is Successfully updated."
	}
	c.IndentedJSON(http.StatusOK, gin.H{"success": true, "Ids": ids, "Messages": messages})

	// if len(insert) != 0 {
	// 	insert = insert[:len(insert)-1]
	// }
	// row, err := db.Query("Insert into movies (Name,Language,Image,HeadImage,Tags,Comment) values " + insert + " RETURNING Id;")

	// if err != nil {
	// 	c.IndentedJSON(http.StatusNotFound, gin.H{"successs": false, "message": "Error in fetching the sql query.", "error": err})
	// 	return
	// }
	// defer func() {
	// 	row.Close()
	// 	db.Close()
	// }()
	// var count int = 0
	// for row.Next() {
	// 	var tmp INDENT.Movie
	// 	err := row.Scan(&tmp.Id)
	// 	for {
	// 		if ids[count] == 0 {
	// 			break
	// 		}
	// 		count++
	// 	}
	// 	if err != nil {
	// 		messages[count].Message += "Unsuccessfully Inserting this data."
	// 	} else {
	// 		messages[count].Message = "Successfully appended this Id."
	// 		ids[count] = tmp.Id
	// 	}
	// 	count++
	// }
	// c.IndentedJSON(http.StatusOK, gin.H{"success": true, "Ids": ids, "Messages": messages})
}

func check(t INDENT.Movie, tags string) bool {
	if strings.Trim(t.Name, " ") == "" || strings.Trim(t.Language, " ") == "" || strings.Trim(t.Image, " ") == "" ||
		strings.Trim(t.Image, " ") == "" || strings.Trim(t.HeadImage, " ") == "" || strings.Trim(tags, " ") == "" || strings.Trim(t.Comment, " ") == "" {
		return true
	}
	return false
}
