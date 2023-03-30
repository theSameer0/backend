package main

import (
	api "example/backend/api"
	"example/backend/api/movie"
	"example/backend/api/seat"
	"example/backend/api/show"
	"example/backend/api/theatre"
	"example/backend/api/ticket"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/movieList", movie.GetMovieList)
	router.GET("/movieDetail/:mId", movie.GetMovieDetail)
	router.GET("/movieTheatres/:mId/:date", movie.GetMovieTheatres)
	router.GET("/theatreList", theatre.GetTheatre)
	router.GET("/theatreDetail/:tId", theatre.GetTheatreDetail)
	router.GET("/theatreMovies/:tId/:date", theatre.GetTheatreMovies)
	router.GET("/seats/:mId/:tId/:date/:time", seat.GetSeatsOfShow)
	router.GET("/ticket", ticket.GetTicket)
	router.GET("/search/:type/:key", api.SearchKeyword)
	router.GET("/language/:loc", api.GetLanguageList)

	router.GET("/lastIds", api.GetLastIds)

	router.POST("/seats/:mId/:tId/:date/:time", seat.SetSeatsOfShow)
	router.POST("/ticket", ticket.SetTicket)

	router.POST("/movies", movie.InsertMovies)
	router.POST("/theatres", theatre.InsertTheatres)
	router.POST("/shows", show.InsertShows)

	router.Run(":8080")
}
