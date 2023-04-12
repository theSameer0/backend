package api

import (
	"example/backend/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetLanguageList(c *gin.Context) {
	// db := database.Connect()
	if c.Request.Method != "GET" {
		c.IndentedJSON(http.StatusMethodNotAllowed, gin.H{"success": false, "message": "This Type of Api not allowed"})
		return
	}
	var languageList []string
	model.DB.Model(&model.Movie{}).Select("Distinct(Language)").Find(&languageList)
	languageList = append(languageList, "All")
	languageList[len(languageList)-1] = languageList[0]
	languageList[0] = "All"
	// row, err := db.Query("select distinct(Language) from movies m , shows s, theatre t where m.Id = s.MovieId and s.TheatreId = t.Id and t.City = $1;", loc)
	// if err != nil {
	// 	c.IndentedJSON(http.StatusNotFound, gin.H{"success": false, "message": "Some error while fetching the query."})
	// 	return
	// }
	// for row.Next() {
	// 	var tmp string
	// 	err := row.Scan(&tmp)
	// 	if err != nil {
	// 		c.IndentedJSON(http.StatusNotFound, gin.H{"success": false, "message": "some error while processing the query"})
	// 		return
	// 	}
	// 	languageList = append(languageList, tmp)
	// }
	c.IndentedJSON(http.StatusOK, gin.H{"success": true, "languageList": languageList})
}
