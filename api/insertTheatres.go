package api

import (
	"example/sameer_mbs/src/server/database"
	"fmt"
	"net/http"
	"strconv"

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
