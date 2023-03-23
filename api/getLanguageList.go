package api

import (
	"example/backend/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetLanguageList(c *gin.Context) {
	var loc string = c.Param("loc")
	db := database.Connect()
	if c.Request.Method != "GET" {
		c.IndentedJSON(http.StatusMethodNotAllowed, gin.H{"success": false, "message": "This Type of Api not allowed"})
		return
	}
	var languageList []string
	row, err := db.Query("select distinct(Language) from movies m , shows s, theatre t where m.Id = s.MovieId and s.TheatreId = t.Id and t.City = $1;", loc)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"success": false, "message": "Some error while fetching the query."})
		return
	}
	languageList = append(languageList, "All")
	for row.Next() {
		var tmp string
		err := row.Scan(&tmp)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"success": false, "message": "some error while processing the query"})
			return
		}
		languageList = append(languageList, tmp)
	}
	c.IndentedJSON(http.StatusOK, gin.H{"success": true, "languageList": languageList})
}
