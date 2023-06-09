package api

import (
	"example/backend/model"
	INDENT "example/backend/v2/api/struct"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func isMatching(key string, name string) bool {
	key = strings.ToLower(key)
	name = strings.ToLower(name)
	var Len int = len(key)
	var nameLen int = len(name)

	for i := range name {
		var idx int = 0
		for ii := range key {
			if i+ii >= nameLen {
				break
			}
			if name[i+ii] == key[ii] {
				idx++
			} else {
				break
			}
		}
		if idx == Len {
			return true
		}
	}
	return false
}

func SearchKeyword(c *gin.Context) {
	var typeSearch = c.Param("type")
	var key string = c.Param("key")
	var pkey = "%" + strings.ToLower(key) + "%"
	// db := DB.Connect()
	fmt.Printf("Great !!! we are connected to a browser\n")
	if c.Request.Method != "GET" {
		c.IndentedJSON(http.StatusMethodNotAllowed, gin.H{"success": false, "message": "This api is not valid. "})
		return
	}
	if typeSearch == "Movies" {
		var data []model.Movie
		model.DB.Model(&model.Movie{}).Where("lower(Name) LIKE ? or ? = 'all'", pkey, key).Find(&data)
		// row, err := db.Query("select Id, Name, Language, Image, HeadImage, Tags, Comment from movies")
		// if err != nil {
		// 	c.IndentedJSON(http.StatusMethodNotAllowed, gin.H{"success": false, "message": "Some issue while fetching querying SQL."})
		// 	return
		// }
		// defer func() {
		// 	row.Close()
		// 	db.Close()
		// }()
		var searchMovie []INDENT.Movie
		for _, t := range data {
			searchMovie = append(searchMovie, INDENT.Movie{
				Id:        t.Id,
				Name:      t.Name,
				Language:  t.Language,
				Image:     t.Image,
				HeadImage: t.Headimage,
				Tags:      strings.Split(t.Tags, ":"),
				Comment:   t.Comment,
			})
		}
		c.IndentedJSON(http.StatusOK, gin.H{"success": true, "searchMovie": searchMovie})
	} else {
		var searchTheatre []model.Theatre
		model.DB.Model(&model.Theatre{}).Where("lower(Name) LIKE ? or ? = 'all'", pkey, key).Find(&searchTheatre)
		// row, err := db.Query("select Id , Name , Image , Location , City , Screen from theatre")
		// if err != nil {
		// 	c.IndentedJSON(http.StatusNotFound, gin.H{"success": false, "message": "Some issue while fetching querying SQL."})
		// 	return
		// }
		// var searchTheatre []Theatre
		// defer func() {
		// 	row.Close()
		// 	db.Close()
		// // }()
		// for row.Next() {
		// 	var tmp Theatre
		// 	err1 := row.Scan(&tmp.Id, &tmp.Name, &tmp.Image, &tmp.Location, &tmp.City, &tmp.Screen)
		// 	if err1 != nil {
		// 		c.IndentedJSON(http.StatusNotFound, gin.H{"success": false, "message": "Some issue while processing the query."})
		// 		return
		// 	}
		// 	if isMatching(key, tmp.Name) || key == "all" {
		// 		searchTheatre = append(searchTheatre, tmp)
		// 	}
		// }
		c.IndentedJSON(http.StatusOK, gin.H{"success": true, "searchTheatre": searchTheatre})
	}
}
