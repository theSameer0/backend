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

	v2_api "example/backend/v2/api"
	v2_movie "example/backend/v2/api/movie"
	v2_seat "example/backend/v2/api/seat"
	v2_show "example/backend/v2/api/show"
	v2_theatre "example/backend/v2/api/theatre"
	v2_ticket "example/backend/v2/api/ticket"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	model.ConnectDatabase()
	database.GrpcConnect()
	defer database.GRPC.Close()
	database.InitRedisCache()

	router.GET("/movieList/v1", movie.GetMovieList)
	router.GET("/movieDetail/v1/:mId", movie.GetMovieDetail)
	router.GET("/movie/total", movie.GetTotalMovie)
	router.GET("/movieTheatres/v1/:mId/:date", movie.GetMovieTheatres)
	router.GET("/theatreList/v1", theatre.GetTheatre)
	router.GET("/theatreDetail/v1/:tId", theatre.GetTheatreDetail)
	router.GET("/theatre/total", theatre.GetTotalTheatre)
	router.GET("/theatreMovies/v1/:tId/:date", theatre.GetTheatreMovies)
	router.GET("/seats/v1/:mId/:tId/:date/:time", seat.GetSeatsOfShow)
	router.GET("/ticket/v1/:id", ticket.GetTicket)
	router.GET("/ticketList/v1", ticket.GetTicketList)
	router.GET("/search/v1/:type/:key", api.SearchKeyword)
	router.GET("/language/v1", api.GetLanguageList)

	router.GET("/lastIds/v1", api.GetLastIds)

	router.POST("/seats/v1/:mId/:tId/:date/:time", seat.SetSeatsOfShow)
	router.POST("/ticket/v1", ticket.SetTicket)

	router.POST("/movies/v1", movie.InsertMovies)
	router.POST("/theatres/v1", theatre.InsertTheatres)
	router.POST("/shows/v1", show.InsertShows)

	router.GET("/movieList/v2", movie.GetMovieList)
	router.GET("/movieDetail/v2/:mId", v2_movie.GetMovieDetail)
	router.GET("/movieTheatres/v2/:mId/:date", v2_movie.GetMovieTheatres)
	router.GET("/theatreList/v2", v2_theatre.GetTheatre)
	router.GET("/theatreDetail/v2/:tId", v2_theatre.GetTheatreDetail)
	router.GET("/theatreMovies/v2/:tId/:date", v2_theatre.GetTheatreMovies)
	router.GET("/seats/v2/:mId/:tId/:date/:time", v2_seat.GetSeatsOfShow)
	router.GET("/ticket/v2/:id", v2_ticket.GetTicket)
	router.GET("/ticketList/v2", v2_ticket.GetTicketList)
	router.GET("/search/v2/:type/:key", v2_api.SearchKeyword)
	router.GET("/language/v2", v2_api.GetLanguageList)

	router.GET("/lastIds/v2", v2_api.GetLastIds)
	router.POST("/seats/v2/:mId/:tId/:date/:time", v2_seat.SetSeatsOfShow)
	router.POST("/ticket/v2", v2_ticket.SetTicket)

	router.POST("/movies/v2", v2_movie.InsertMovies)
	router.POST("/theatres/v2", v2_theatre.InsertTheatres)
	router.POST("/shows/v2", v2_show.InsertShows)

	router.Run(":8080")
}
