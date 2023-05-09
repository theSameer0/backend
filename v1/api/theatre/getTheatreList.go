package theatre

import (
	"example/backend/model"
	INDENT "example/backend/v1/api/struct"
	"example/backend/v1/database"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetTheatre(c *gin.Context) {
	var getLang []string = c.QueryArray("lang[]")
	var loc string = c.Query("loc")
	var page = c.Query("page")
	var offset = c.Query("offset")
	currentPage, err := strconv.Atoi(page)
	database.CheckErr(err)
	currentOffset, err := strconv.Atoi(offset)
	database.CheckErr(err)
	var skipRows int = (currentPage - 1) * currentOffset
	var lang string
	var isAll bool = false
	var len int = len(getLang)
	var num int

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
	// db := DB.Connect()
	fmt.Printf("Great !!! we are connected to a browser\n")
	if c.Request.Method != "GET" {
		c.IndentedJSON(http.StatusMethodNotAllowed, gin.H{"success": false, "message": "This api is not valid. "})
		return
	}
	var theatreList []INDENT.TheatreList
	subQuery1 := model.DB.Model(&model.Movie{}).Select("Id").Where("Language In ("+lang+") or ? = true", isAll)
	subQuery2 := model.DB.Model(&model.Show{}).Where("MovieId IN (?) ", subQuery1).Select("Distinct(TheatreId)")
	model.DB.Offset(skipRows).Limit(currentOffset).Model(&model.Theatre{}).Where("Id in (?) and City = ?", subQuery2, loc).Find(&theatreList)
	model.DB.Model(&model.Theatre{}).Where("Id in (?) and City = ?", subQuery2, loc).Select("Count(Id)").Find(&num)
	// row, err := db.Query("select Id , Name , Image , Location from theatre where City = $1 and Id in (select distinct(s.TheatreId) from shows s,movies m where s.MovieId = m.Id and (m.Language in ("+lang+") or $2=true));", loc, isAll)
	// if err != nil {
	// 	c.IndentedJSON(http.StatusMethodNotAllowed, gin.H{"success": false, "message": "Some issue while fetching querying SQL.", "Error": err})
	// 	return
	// }
	// defer func() {
	// 	row.Close()
	// 	db.Close()
	// }()

	// var theatreList []INDENT.Theatre
	// for row.Next() {
	// 	var tmpTheatre INDENT.Theatre
	// 	err := row.Scan(&tmpTheatre.Id, &tmpTheatre.Name, &tmpTheatre.Image, &tmpTheatre.Location)
	// 	if err != nil {
	// 		log.Println(err)
	// 		c.IndentedJSON(http.StatusMethodNotAllowed, gin.H{"success": false, "message": "Error in scanning the row.", "Error": err})
	// 		return
	// 	}
	// 	theatreList = append(theatreList, tmpTheatre)
	// }
	c.JSON(http.StatusOK, gin.H{"success": true, "theatreList": theatreList, "totalTheatre": num})
}
