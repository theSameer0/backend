package model

type Ticket struct {
	Id        string `json:"id" gorm:"primary_key"`
	Date      string `json:"date"`
	Time      string `json:"time"`
	Seats     string `json:"seats"`
	Screen    int    `json:"screen"`
	Seatcount int    `json:"seatcount"`
	Movieid   int    `json:"movieid"`
	Theatreid int    `json:"theatreid"`
	Showid    int    `json:"showid"`
	Timestamp string `json:"TimeStamp"`
}

func (Ticket) TableName() string {
	return "ticket"
}
