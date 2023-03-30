package theatre

import (
	INDENT "example/backend/api/struct"
	DB "example/backend/database"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetTheatre(c *gin.Context) {
	var getLang []string = c.QueryArray("lang[]")
	var loc string = c.Query("loc")
	var lang string
	var isAll bool = false
	var len int = len(getLang)

	for i, t := range getLang {
		if t == "All" {
			isAll = true
		}
		if i != len-1 {
			lang += "'" + strings.Trim(t, "") + "',"
		} else {
			lang += "'" + strings.Trim(t, "") + "'"
		}
	}
	db := DB.Connect()
	fmt.Printf("Great !!! we are connected to a browser\n")
	if c.Request.Method != "GET" {
		c.IndentedJSON(http.StatusMethodNotAllowed, gin.H{"success": false, "message": "This api is not valid. "})
		return
	}
	row, err := db.Query("select Id , Name , Image , Location from theatre where City = $1 and Id in (select distinct(s.TheatreId) from shows s,movies m where s.MovieId = m.Id and (m.Language in ("+lang+") or $2=true));", loc, isAll)
	if err != nil {
		c.IndentedJSON(http.StatusMethodNotAllowed, gin.H{"success": false, "message": "Some issue while fetching querying SQL.", "Error": err})
		return
	}
	defer func() {
		row.Close()
		db.Close()
	}()

	var theatreList []INDENT.Theatre
	for row.Next() {
		var tmpTheatre INDENT.Theatre
		err := row.Scan(&tmpTheatre.Id, &tmpTheatre.Name, &tmpTheatre.Image, &tmpTheatre.Location)
		if err != nil {
			log.Println(err)
			c.IndentedJSON(http.StatusMethodNotAllowed, gin.H{"success": false, "message": "Error in scanning the row.", "Error": err})
			return
		}
		theatreList = append(theatreList, tmpTheatre)
	}
	c.IndentedJSON(http.StatusOK, gin.H{"success": true, "theatreList": theatreList})
}