package INDENT

type Movie struct {
	Id        int      `json:"id"`
	Name      string   `json:"name"`
	Language  string   `json:"language"`
	Image     string   `json:"image"`
	HeadImage string   `json:"headImage"`
	Tags      []string `json:"tags"`
	Comment   string   `json:"comment"`
}

type MovieList struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Language string `json:"language"`
	Image    string `json:"image"`
}

type MovieData struct {
	Id        int      `json:"id"`
	Name      string   `json:"name"`
	HeadImage string   `json:"headImage"`
	Tags      []string `json:"tags"`
	Comment   string   `json:"comment"`
}
type ShowTime struct {
	ShowId int    `json:"showId"`
	Time   string `json:"time"`
	Seat   string `json:"seat"`
}
type MovieShow struct {
	TheatreId       int        `json:"theatreId"`
	TheatreName     string     `json:"theatreName"`
	TheatreLocation string     `json:"theatreLocation"`
	ShowTime        []ShowTime `json:"showTime"`
	ShowDate        string     `json:"showDate"`
	ShowScreen      int        `json:"showScreen"`
}

type TheatreList struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Image    string `json:"image"`
}
type TheatreShow struct {
	MovieId       int        `json:"movieId"`
	MovieName     string     `json:"movieName"`
	MovieLanguage string     `json:"movieLanguage"`
	MovieImage    string     `json:"movieImage"`
	MovieTags     []string   `json:"movieTags"`
	ShowDate      string     `json:"showDate"`
	ShowScreen    int        `json:"showScreen"`
	ShowTime      []ShowTime `json:"showTime"`
}
type Theatre struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Image    string `json:"image"`
	City     string `json:"city"`
	Screen   int    `json:"screen"`
}

type Show struct {
	Id        int    `json:"id"`
	MovieId   int    `json:"movieId"`
	TheatreId int    `json:"theatreId"`
	Date      string `json:"date"`
	Time      string `json:"timing"`
	Seats     string `json:"seats"`
	Screen    int    `json:"screen"`
}

type Ticket struct {
	Id        string `json:"id"`
	Date      string `json:"date"`
	Time      string `json:"time"`
	Seats     string `json:"seats"`
	Screen    int    `json:"screen"`
	SeatCount int    `json:"seatCount"`
	MovieId   int    `json:"movieId"`
	TheatreId int    `json:"theatreId"`
	ShowId    int    `json:"showId"`
	TimeStamp string `json:"timeStamp"`
}
type TicketList struct {
	Id      string `json:"id"`
	Date    string `json:"date"`
	Time    string `json:"time"`
	Movieid int    `json:"movieId"`
}

type GetSeat struct {
	Seats string `json:"seats"`
}

type Message struct {
	Message string `json:"message"`
}
