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
	data.Id = movies.Id
	data.Name = responseFromGrpcServer(movies.Name)
	data.HeadImage = movies.Headimage
	data.Tags = strings.Split(movies.Tags, ":")
	data.Comment = movies.Comment
	c.JSON(http.StatusOK, gin.H{"success": true, "movieData": data})
}

func responseFromGrpcServer(name string) string {
	c1 := pb.NewCapitalizeClient(database.GRPC)
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
	defer cancel()
	response, err := c1.CapitalName(ctx, &pb.CapitalizeRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Printf("Recieved data from service running at port 50051: %v\n", response.GetName())
	return string(response.GetName())
}
