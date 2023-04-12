package theatre

import (
	"example/backend/model"
	INDENT "example/backend/v2/api/struct"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func InsertTheatres(c *gin.Context) {
	// db := database.Connect()
	if c.Request.Method != "POST" {
		c.IndentedJSON(http.StatusMethodNotAllowed, gin.H{"success": false, "message": "This type of api is not allowed"})
		return
	}
	var data []model.Theatre
	if err := c.BindJSON(&data); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"success": false, "message": "cant bind the input data."})
		return
	}
	var ids []int
	var messages []INDENT.Message
	var theatreList []model.Theatre

	for i, t := range data {
		ids = append(ids, -1)

		if checkTheatre(t) {
			messages = append(messages, INDENT.Message{Message: "Unsuccessful Some values are Empty."})
		} else {
			theatreList = append(theatreList, model.Theatre{
				Name:     t.Name,
				Location: t.Location,
				Image:    t.Image,
				City:     t.City,
				Screen:   t.Screen,
			})
			messages = append(messages, INDENT.Message{Message: ""})
			ids[i] = 0
		}
	}
	fmt.Printf("%v", data)
	model.DB.Select("Name", "Image", "Location", "City", "Screen").Create(&theatreList)
	// row, err := db.Query("Insert into theatre (Name,Location,Image,City,Screen) values " + insert + " RETURNING Id;")
	// if err != nil {
	// 	c.IndentedJSON(http.StatusNotFound, gin.H{"successs": false, "message": "Error in fetching the sql query."})
	// 	return
	// }
	// defer func() {
	// 	row.Close()
	// 	db.Close()
	// }()
	var count int = 0
	for _, t := range theatreList {
		for {
			if ids[count] == 0 {
				break
			}
			count++
		}
		messages[count].Message = "Data Insert Was Successful for this record."
		ids[count] = t.Id
	}
	c.IndentedJSON(http.StatusOK, gin.H{"success": true, "Ids": ids, "Messages": messages})
}

func checkTheatre(t model.Theatre) bool {
	if strings.Trim(t.Name, " ") == "" || strings.Trim(t.Location, " ") == "" || strings.Trim(t.Image, " ") == "" || strings.Trim(t.City, " ") == "" || t.Screen <= 0 {
		return true
	}
	return false
}
