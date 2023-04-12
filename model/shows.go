package model

type Show struct {
	Id        int    `json:"id" gorm:"primary_key"`
	Movieid   int    `json:"movieid"`
	Theatreid int    `json:"theatreid"`
	Date      string `json:"date"`
	Time      string `json:"timing"`
	Seats     string `json:"seats"`
	Screen    int    `json:"screen"`
}
