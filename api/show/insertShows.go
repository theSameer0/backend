package show

import (
	INDENT "example/backend/api/struct"
	"example/backend/database"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func InsertShows(c *gin.Context) {
	db := database.Connect()
	if c.Request.Method != "POST" {
		c.IndentedJSON(http.StatusMethodNotAllowed, gin.H{"success": false, "message": "This type of api is not allowed"})
		return
	}
	var showList []INDENT.Show
	if err := c.BindJSON(&showList); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"success": false, "message": "cant bind the input data."})
		return
	}
	var ids []int
	var messages []INDENT.Message
	var insert string
	for i, t := range showList {
		ids = append(ids, -1)

		if checkShow(t) {
			messages = append(messages, INDENT.Message{Message: "Unsuccessful Some values are Empty."})
		} else {
			messages = append(messages, INDENT.Message{Message: ""})
			ids[i] = 0
			insert += "('" + t.Time + "','" + t.Seats + "','" + t.Date + "'," + strconv.Itoa(t.Screen) + "," + strconv.Itoa(t.MovieId) + "," + strconv.Itoa(t.TheatreId) + "),"
		}
	}
	if len(insert) != 0 {
		insert = insert[:len(insert)-1]
	}
	fmt.Print("Insert into shows (Time,Seats,Date,Screen,MovieId,TheatreId) values " + insert + " RETURNING Id;")
	row, err := db.Query("Insert into shows (Time,Seats,Date,Screen,MovieId,TheatreId) values " + insert + " RETURNING Id;")
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"successs": false, "message": "Error in fetching the sql query."})
		return
	}
	defer func() {
		row.Close()
		db.Close()
	}()
	var count int = 0
	for row.Next() {
		var tmp int
		err := row.Scan(&tmp)
		for {
			if ids[count] == 0 {
				break
			}
			count++
		}
		if err != nil {
			messages[count].Message = "Data Insert Was UnsuccessFul for this record due to some error."
		} else {
			messages[count].Message = "Data Insert Was Successful for this record."
			ids[count] = tmp
		}
	}
	c.IndentedJSON(http.StatusOK, gin.H{"success": true, "Ids": ids, "Message": messages})
}

func checkShow(t INDENT.Show) bool {
	if strings.Trim(t.Time, " ") == "" || strings.Trim(t.Date, " ") == "" || t.Screen <= 0 || t.MovieId < 0 || t.TheatreId < 0 {
		return true
	}
	return false
}
