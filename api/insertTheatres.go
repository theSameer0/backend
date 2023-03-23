package api

import (
	"example/backend/database"
	"fmt"
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
	var theatreList []Theatre
	if err := c.BindJSON(&theatreList); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"success": false, "message": "cant bind the input data."})
		return
	}
	var ids []string
	var insert string
	for i, t := range theatreList {
		ids = append(ids, t.Id)
		insert += "("
		insert += "'" + t.Id + "','" + t.Name + "','" + t.Location + "','" + t.Image + "','" + t.City + "'," + strconv.Itoa(t.Screen) + ")"
		if i != len(theatreList)-1 {
			insert += ","
		}
		if checkTheatre(t) {
			c.IndentedJSON(http.StatusMethodNotAllowed, gin.H{"successs": false, "message": "Some fields are empty."})
			return
		}
	}
	fmt.Printf("%v", insert)
	row, err := db.Query("Insert into theatre (Id,Name,Location,Image,City,Screen) values " + insert + ";")
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

func checkTheatre(t Theatre) bool {
	if strings.Trim(t.Id, " ") == "" || strings.Trim(t.Name, " ") == "" || strings.Trim(t.Location, " ") == "" || strings.Trim(t.Image, " ") == "" || strings.Trim(t.City, " ") == "" || t.Screen <= 0 {
		return true
	}
	return false
}
