package theatre

import (
	INDENT "example/backend/api/struct"
	"example/backend/database"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func InsertTheatres(c *gin.Context) {
	db := database.Connect()
	if c.Request.Method != "POST" {
		c.IndentedJSON(http.StatusMethodNotAllowed, gin.H{"success": false, "message": "This type of api is not allowed"})
		return
	}
	var theatreList []INDENT.Theatre
	if err := c.BindJSON(&theatreList); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"success": false, "message": "cant bind the input data."})
		return
	}
	var ids []int
	var messages []INDENT.Message
	var insert string
	for i, t := range theatreList {
		ids = append(ids, -1)

		if checkTheatre(t) {
			messages = append(messages, INDENT.Message{Message: "Unsuccessful Some values are Empty."})
		} else {
			insert += "('" + t.Name + "','" + t.Location + "','" + t.Image + "','" + t.City + "'," + strconv.Itoa(t.Screen) + "),"
			messages = append(messages, INDENT.Message{Message: ""})
			ids[i] = 0
		}
	}
	if len(insert) != 0 {
		insert = insert[:len(insert)-1]
	}
	row, err := db.Query("Insert into theatre (Name,Location,Image,City,Screen) values " + insert + " RETURNING Id;")
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
	c.IndentedJSON(http.StatusOK, gin.H{"success": true, "Ids": ids, "Messages": messages})
}

func checkTheatre(t INDENT.Theatre) bool {
	if strings.Trim(t.Name, " ") == "" || strings.Trim(t.Location, " ") == "" || strings.Trim(t.Image, " ") == "" || strings.Trim(t.City, " ") == "" || t.Screen <= 0 {
		return true
	}
	return false
}
