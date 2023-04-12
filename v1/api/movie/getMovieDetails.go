package movie

import (
	"context"
	"example/backend/model"
	INDENT "example/backend/v1/api/struct"
	"example/backend/v1/database"
	pb "example/backend/v1/proto"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func GetMovieDetail(c *gin.Context) {
	var mId = c.Param("mId")
	fmt.Printf("Great !!! we are connected to a browser\n")
	if c.Request.Method != "GET" {
		c.IndentedJSON(http.StatusMethodNotAllowed, gin.H{"success": false, "message": "This api is not valid. "})
		return
	}
	var movies model.Movie
	var data INDENT.MovieData
	model.DB.Model(&movies).Where("Id = ?", mId).Find(&movies)
	// var ctx = context.Context
	// r, err = c.CapitalName(ctx, &pb.CapitalizeRequest{Name: movies.Name})
	// if err != nil {
	// 	log.Fatalf("could not greet: %v", err)
	// }

	// log.Printf("Greeting: %s", r.GetMessage())
	data.Id = movies.Id
	// data.Name = movies.Name
	data.HeadImage = movies.Headimage
	data.Tags = strings.Split(movies.Tags, ":")
	data.Comment = movies.Comment

	c1 := pb.NewCapitalizeClient(database.GRPC)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c1.CapitalName(ctx, &pb.CapitalizeRequest{Name: movies.Name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Printf("Recieved data from service running at port 50051: %v\n", r.GetName())
	data.Name = r.GetName()

	c.JSON(http.StatusOK, gin.H{"success": true, "movieData": data})
}
