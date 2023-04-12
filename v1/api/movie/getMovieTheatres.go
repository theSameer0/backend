package movie

import (
	"example/backend/model"
	INDENT "example/backend/v1/api/struct"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMovieTheatres(c *gin.Context) {
	var mId = c.Param("mId")
	var date = c.Param("date")
	// id, err := strconv.Atoi(mId)
	// database.CheckErr(err)

	// db := DB.Connect()
	fmt.Printf("Great !!! we are connected to a browser\n")
	if c.Request.Method != "GET" {
		c.IndentedJSON(http.StatusMethodNotAllowed, gin.H{"success": false, "message": "This api is not valid. "})
		return
	}
	var data []INDENT.MovieShow

	type getData struct {
		Id        int    `json:"id"`
		Time      string `json:"time"`
		Seats     string `json:"seats"`
		Date      string `json:"date"`
		Movieid   int    `json:"movieid"`
		Theatreid int    `json:"theatreid"`
		Screen    int    `json:"screen"`
		Name      string `json:"name"`
		Location  string `json:"location"`
	}
	var fetchData []getData
	// showQuery := model.DB.Model(&model.Show{})
	// theatreQuery := model.DB.Model(&model.Theatre{})

	model.DB.Model(&model.Show{}).Order("shows.Id").Joins("left join theatres on theatres.Id = shows.TheatreId").Select("shows.Id", "shows.Time", "shows.Seats", "shows.Date", "shows.MovieId", "shows.Screen", "shows.theatreId", "theatres.Name", "theatres.Location").Where("shows.MovieId = ? and shows.Date = ?", mId, date).Find(&fetchData)
	for _, t := range fetchData {
		var isFound bool = false
		for ii, tt := range data {
			if t.Theatreid == tt.TheatreId {
				isFound = true
				data[ii].ShowTime = append(data[ii].ShowTime, INDENT.ShowTime{
					ShowId: t.Id,
					Time:   t.Time,
					Seat:   t.Seats,
				})
				break
			}
		}
		if !isFound {
			data = append(data, INDENT.MovieShow{
				TheatreId:       t.Theatreid,
				TheatreName:     t.Name,
				TheatreLocation: t.Location,
				ShowTime: []INDENT.ShowTime{
					{
						ShowId: t.Id,
						Time:   t.Time,
						Seat:   t.Seats,
					},
				},
				ShowDate:   t.Date,
				ShowScreen: t.Screen,
			})
		}
	}

	c.IndentedJSON(http.StatusOK, gin.H{"success": true, "theatreList": data})
	// row2, err2 := db.Query("select s.Id,s.Time,s.Seats,s.Date,s.MovieId,s.TheatreId,s.Screen,t.Name,t.Location from shows s,theatre t where s.TheatreId = t.Id and s.MovieId = $1 and s.Date = $2 order by Time,TheatreId", mId, date)
	// if err2 != nil {
	// 	c.IndentedJSON(http.StatusMethodNotAllowed, gin.H{"success": false, "message": "Some issue while fetching querying SQL.", "Error": err2.Error()})
	// 	return
	// }
	// defer func() {
	// 	row2.Close()
	// 	db.Close()
	// }()

	// var theatreList []INDENT.MovieShow
	// for row2.Next() {
	// 	var tmpShow INDENT.Show
	// 	var name, location string
	// 	err := row2.Scan(&tmpShow.Id, &tmpShow.Time, &tmpShow.Seats, &tmpShow.Date, &tmpShow.MovieId, &tmpShow.TheatreId, &tmpShow.Screen, &name, &location)
	// 	if err != nil {
	// 		log.Println(err)
	// 		c.IndentedJSON(http.StatusMethodNotAllowed, gin.H{"success": false, "message": "Error in scanning the row.", "Error": err.Error()})
	// 		return
	// 	}
	// 	var isFound bool = false
	// 	for i, t := range theatreList {
	// 		if t.TheatreId == tmpShow.TheatreId {
	// 			isFound = true
	// 			theatreList[i].ShowTime = append(theatreList[i].ShowTime, INDENT.ShowTime{
	// 				ShowId: tmpShow.Id,
	// 				Time:   tmpShow.Time,
	// 				Seat:   tmpShow.Seats,
	// 			})
	// 		}
	// 	}
	// 	if !isFound {
	// 		theatreList = append(theatreList, INDENT.MovieShow{
	// 			TheatreId:       tmpShow.TheatreId,
	// 			TheatreName:     name,
	// 			TheatreLocation: location,
	// 			ShowTime: []INDENT.ShowTime{
	// 				{
	// 					ShowId: tmpShow.Id,
	// 					Time:   tmpShow.Time,
	// 					Seat:   tmpShow.Seats,
	// 				},
	// 			},
	// 			ShowDate:   tmpShow.Date,
	// 			ShowScreen: tmpShow.Screen,
	// 		})
	// 	}
	// }

	// if len(theatreList) == 0 {
	// 	c.IndentedJSON(http.StatusNotFound, gin.H{"success": false, "message": "Hey not found the data!!"})
	// 	return
	// }
	// c.IndentedJSON(http.StatusOK, gin.H{"success": true, "theatreList": theatreList})
}
