package movie

import (
	"encoding/json"
	"example/backend/model"
	INDENT "example/backend/v1/api/struct"
	"example/backend/v1/database"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func GetMovieList(c *gin.Context) {
	link := c.Request.Referer()
	var finalResult INDENT.MovieListCacheType
	result, err := database.Cache.GetMovieListCache(link)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "error in checking in cache"})
		log.Fatal(err)
	}
	if result != nil {
		err := json.Unmarshal(result, &finalResult)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "error in fetching in cache"})
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, gin.H{"success": true, "movieList": finalResult.MoviesList, "totalMovies": finalResult.TotalMovies})
		return
	}

	var getLang []string = c.QueryArray("lang[]")
	var loc string = c.Query("loc")
	var page = c.Query("page")
	var offset = c.Query("offset")

	currentPage, err := strconv.Atoi(page)
	database.CheckErr(err)
	currentOffset, err := strconv.Atoi(offset)
	var rowsToSkip int = (currentPage - 1) * currentOffset
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
	fmt.Printf("Great !!! we are connected to a browser Guiz\n")
	if c.Request.Method != "GET" {
		c.IndentedJSON(http.StatusMethodNotAllowed, gin.H{"success": false, "message": "This api is not valid. "})
		return
	}

	var movieList []INDENT.MovieList
	subQuery1 := model.DB.Model(&model.Theatre{}).Select("Id").Where("City = ?", loc)
	// fmt.Printf("%v", subQuery1)
	subQuery2 := model.DB.Model(&model.Show{}).Where("TheatreId IN (?) ", subQuery1).Select("Distinct(MovieId)")
	model.DB.Limit(currentOffset).Offset(rowsToSkip).Model(&model.Movie{}).Where("(Language in ("+lang+") or ? = true) and Id in (?)", isAll, subQuery2).Find(&movieList)
	model.DB.Model(&model.Movie{}).Where("(Language in ("+lang+") or ? = true) and Id in (?)", isAll, subQuery2).Select("Count(Id)").Find(&num)

	finalResult.MoviesList = movieList
	finalResult.TotalMovies = num

	err = database.Cache.SetMovieListCache(link, finalResult, 1*time.Minute)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "error in updating in cache"})
		log.Fatal(err)
	}
	// row, err := db.Query("select Id, Name, Language, Image , HeadImage, Tags,Comment from movies where (Language in ("+lang+") or $1 = true) and Id in (select distinct(s.MovieId) from theatre t,shows s where s.TheatreId = t.Id and t.City = $2);", isAll, loc)
	// if err != nil {
	// 	c.IndentedJSON(http.StatusMethodNotAllowed, gin.H{"success": false, "message": "Some issue while fetching querying SQL.", "Error": err})
	// 	return
	// }
	// defer func() {
	// 	row.Close()
	// 	db.Close()
	// }()

	// var movieList []st.Movie
	// for row.Next() {
	// 	var tmpMovie st.Movie
	// 	var tags string
	// 	row.Scan(&tmpMovie.Id, &tmpMovie.Name, &tmpMovie.Language, &tmpMovie.Image, &tmpMovie.HeadImage, tags, &tmpMovie.Comment)
	// 	if err != nil {
	// 		log.Println(err)
	// 		c.IndentedJSON(http.StatusMethodNotAllowed, gin.H{"success": false, "message": "Error in scanning the row.", "Error": err})
	// 		return
	// 	}
	// 	tmpMovie.Tags = strings.Split(tags, ":")
	// 	movieList = append(movieList, tmpMovie)
	// }
	// c.IndentedJSON(http.StatusOK, gin.H{"success": true, "movieList": movieList})
	c.JSON(http.StatusOK, gin.H{"success": true, "movieList": movieList, "totalMovies": num})
}
