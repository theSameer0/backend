package api

import (
	DB "example/sameer_mbs/src/server/database"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func isMatching(key string, name string) bool {
	key = strings.ToLower(key)
	name = strings.ToLower(name)
	fmt.Printf("%v,%v", key, name)
	var len int = len(key)
	var idx int = 0
	for i := range name {
		if name[i] == key[idx] {
			idx++
			if idx == len {
				return true
			}
		}
	}
	return false
}

func SearchKeyword(c *gin.Context) {
	var typeSearch = c.Param("type")
	var key string = c.Param("key")
	db := DB.Connect()
	fmt.Printf("Great !!! we are connected to a browser\n")
	if c.Request.Method != "GET" {
		c.IndentedJSON(http.StatusMethodNotAllowed, gin.H{"success": false, "message": "This api is not valid. "})
		return
	}
	if typeSearch == "Movies" {
		row, err := db.Query("select Id, Name, Language, Image, HeadImage, Tags, Comment from movies")
		if err != nil {
			c.IndentedJSON(http.StatusMethodNotAllowed, gin.H{"success": false, "message": "Some issue while fetching querying SQL."})
			return
		}
		defer func() {
			row.Close()
			db.Close()
		}()
		var searchMovie []Movie
		for row.Next() {
			var tmp Movie
			var tags string
			err1 := row.Scan(&tmp.Id, &tmp.Name, &tmp.Language, &tmp.Image, &tmp.HeadImage, &tags, &tmp.Comment)
			if err1 != nil {
				c.IndentedJSON(http.StatusNotFound, gin.H{"success": false, "message": "Error in processing the quries."})
				return
			}
			tmp.Tags = strings.Split(tags, ":")
			if isMatching(key, tmp.Name) || key == "all" {
				searchMovie = append(searchMovie, tmp)
			}
		}
		c.IndentedJSON(http.StatusOK, gin.H{"success": true, "searchMovie": searchMovie})
	} else {
		row, err := db.Query("select Id , Name , Image , Location , City , Screen from theatre")
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"success": false, "message": "Some issue while fetching querying SQL."})
			return
		}
		var searchTheatre []Theatre
		defer func() {
			row.Close()
			db.Close()
		}()
		for row.Next() {
			var tmp Theatre
			err1 := row.Scan(&tmp.Id, &tmp.Name, &tmp.Image, &tmp.Location, &tmp.City, &tmp.Screen)
			if err1 != nil {
				c.IndentedJSON(http.StatusNotFound, gin.H{"success": false, "message": "Some issue while processing the query."})
				return
			}
			if isMatching(key, tmp.Name) || key == "all" {
				searchTheatre = append(searchTheatre, tmp)
			}
		}
		c.IndentedJSON(http.StatusOK, gin.H{"success": true, "searchTheatre": searchTheatre})
	}
}