package api

type Movie struct {
	Id        string   `json:"id"`
	Name      string   `json:"name"`
	Language  string   `json:"language"`
	Image     string   `json:"image"`
	HeadImage string   `json:"headImage"`
	Tags      []string `json:"tags"`
	Comment   string   `json:"comment"`
}

type Theatre struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Image    string `json:"image"`
	City     string `json:"city"`
	Screen   int    `json:"screen"`
}

type Seats struct {
	Time  string       `json:"time"`
	Seats [12][12]bool `json:"seats"` //n*m B4 C7
}
type Show struct {
	Id        string `json:"id"`
	MovieId   string `json:"movieId"`
	TheatreId string `json:"theatreId"`
	Date      string `json:"date"`
	Time      string `json:"timing"`
	Seats     string `json:"seats"`
	Screen    int    `json:"screen"`
}

type rawMovieTime struct {
	Id        string   `json:"id"`
	Name      string   `json:"name"`
	Tags      []string `json:"tags"`
	Image     string   `json:"image"`
	TheatreId string   `json:"theatreId"`
	ShowId    string   `json:"showId"`
	Time      string   `json:"time"`
	Date      string   `json:"date"`
	Seats     string   `json:"seats"`
}

type Seat struct {
	Timing string `json:"timing"`
	Seats  string `json:"seats"`
}

type MovieTime struct {
	Id        string   `json:"id"`
	Name      string   `json:"name"`
	Tags      []string `json:"tags"`
	Image     string   `json:"image"`
	TheatreId string   `json:"theatreId"`
	ShowId    string   `json:"showId"`
	Date      string   `json:"date"`
	Time      []Seat   `json:"time"`
}

type Ticket struct {
	Id        string `json:"id"`
	Date      string `json:"date"`
	Time      string `json:"time"`
	Seats     string `json:"seats"`
	Screen    int    `json:"screen"`
	SeatCount int    `json:"seatCount"`
	MovieId   string `json:"movieId"`
	TheatreId string `json:"theatreId"`
	ShowId    string `json:"showId"`
}

type getSeat struct {
	Seats string `json:"seats"`
}
