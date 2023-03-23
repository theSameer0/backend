package main

import (
	api "example/backend/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/movies", api.GetMovies)
	router.GET("/theatres/:loc", api.GetTheatre)
	router.GET("/filterMovies", api.GetMoviesByFilter)
	router.GET("/seats/:mId/:tId/:date/:time", api.GetSeatsOfShow)
	router.GET("/theatre/:tId/:date", api.GetTheatreMovies)
	router.GET("/show/:mId/:date", api.GetMovieShows)
	router.GET("/ticket", api.GetTicket)
	router.GET("/search/:type/:key", api.SearchKeyword)
	router.GET("/language/:loc", api.GetLanguageList)

	router.GET("/lastIds", api.GetLastIds)

	router.POST("/seats/:mId/:tId/:date/:time", api.SetSeatsOfShow)
	router.POST("/ticket", api.SetTicket)

	router.POST("/movies", api.InsertMovies)
	router.POST("/theatres", api.InsertTheatres)
	router.POST("/shows", api.InsertShows)

	router.Run(":8080")
}
