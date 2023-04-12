package main

import (
	"example/backend/model"
	"example/backend/v1/api"
	"example/backend/v1/api/movie"
	"example/backend/v1/api/seat"
	"example/backend/v1/api/show"
	"example/backend/v1/api/theatre"
	"example/backend/v1/api/ticket"

	"example/backend/v1/database"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	model.ConnectDatabase()
	database.GrpcConnect()
	defer database.GRPC.Close()
	router.GET("/movieList", movie.GetMovieList)
	router.GET("/movieDetail/:mId", movie.GetMovieDetail)
	router.GET("/movieTheatres/:mId/:date", movie.GetMovieTheatres)
	router.GET("/theatreList", theatre.GetTheatre)
	router.GET("/theatreDetail/:tId", theatre.GetTheatreDetail)
	router.GET("/theatreMovies/:tId/:date", theatre.GetTheatreMovies)
	router.GET("/seats/:mId/:tId/:date/:time", seat.GetSeatsOfShow)
	router.GET("/ticket/:id", ticket.GetTicket)
	router.GET("/ticketList", ticket.GetTicketList)
	router.GET("/search/:type/:key", api.SearchKeyword)
	router.GET("/language", api.GetLanguageList)

	router.GET("/lastIds", api.GetLastIds)

	router.POST("/seats/:mId/:tId/:date/:time", seat.SetSeatsOfShow)
	router.POST("/ticket", ticket.SetTicket)

	router.POST("/movies", movie.InsertMovies)
	router.POST("/theatres", theatre.InsertTheatres)
	router.POST("/shows", show.InsertShows)

	router.Run(":8080")
}
