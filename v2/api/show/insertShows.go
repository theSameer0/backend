package show

import (
	"example/backend/model"
	INDENT "example/backend/v2/api/struct"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func InsertShows(c *gin.Context) {
	// db := database.Connect()
	if c.Request.Method != "POST" {
		c.IndentedJSON(http.StatusMethodNotAllowed, gin.H{"success": false, "message": "This type of api is not allowed"})
		return
	}
	var data []INDENT.Show
	if err := c.BindJSON(&data); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"success": false, "message": "cant bind the input data."})
		return
	}
	var ids []int
	var messages []INDENT.Message
	var showList []model.Show
	for i, t := range data {
		ids = append(ids, -1)

		if checkShow(t) {
			messages = append(messages, INDENT.Message{Message: "Unsuccessful Some values are Empty."})
		} else {
			showList = append(showList, model.Show{
				Date:      t.Date,
				Time:      t.Time,
				Seats:     t.Seats,
				Screen:    t.Screen,
				Movieid:   t.MovieId,
				Theatreid: t.TheatreId,
			})
			messages = append(messages, INDENT.Message{Message: ""})
			ids[i] = 0
		}
	}

	model.DB.Select("Date", "Time", "Seats", "Screen", "MovieId", "TheatreId").Create(&showList)
	// row, err := db.Query("Insert into shows (Time,Seats,Date,Screen,MovieId,TheatreId) values " + insert + " RETURNING Id;")
	// if err != nil {
	// 	c.IndentedJSON(http.StatusNotFound, gin.H{"successs": false, "message": "Error in fetching the sql query."})
	// 	return
	// }
	// defer func() {
	// 	row.Close()
	// 	db.Close()
	// }()
	var count int = 0
	for _, t := range showList {
		for {
			if ids[count] == 0 {
				break
			}
			count++
		}
		messages[count].Message = "Data Insert Was Successful for this record."
		ids[count] = t.Id

	}
	c.IndentedJSON(http.StatusOK, gin.H{"success": true, "Ids": ids, "Message": messages})
}

func checkShow(t INDENT.Show) bool {
	if strings.Trim(t.Time, " ") == "" || strings.Trim(t.Date, " ") == "" || t.Screen <= 0 || t.MovieId < 0 || t.TheatreId < 0 {
		return true
	}
	return false
}
