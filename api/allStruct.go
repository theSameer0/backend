package api

type Movie struct {
	Id        int      `json:"id"`
	Name      string   `json:"name"`
	Language  string   `json:"language"`
	Image     string   `json:"image"`
	HeadImage string   `json:"headImage"`
	Tags      []string `json:"tags"`
	Comment   string   `json:"comment"`
}

type Theatre struct {
	Id       int    `json:"id"`
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
	Id        int    `json:"id"`
	MovieId   int    `json:"movieId"`
	TheatreId int    `json:"theatreId"`
	Date      string `json:"date"`
	Time      string `json:"timing"`
	Seats     string `json:"seats"`
	Screen    int    `json:"screen"`
}

type TheatreMovie struct {
	Id       int      `json:"movieId"`
	Name     string   `json:"movieName"`
	Language string   `json:"language"`
	Tags     []string `json:"tags"`
	Time     []Seat   `json:"time"`
}

type rawMovieTime struct {
	Id        int      `json:"id"`
	Name      string   `json:"name"`
	Tags      []string `json:"tags"`
	Image     string   `json:"image"`
	TheatreId int      `json:"theatreId"`
	ShowId    int      `json:"showId"`
	Time      string   `json:"time"`
	Date      string   `json:"date"`
	Seats     string   `json:"seats"`
}

type Seat struct {
	Timing string `json:"timing"`
	Seats  string `json:"seats"`
}

type MovieTime struct {
	Id        int      `json:"id"`
	Name      string   `json:"name"`
	Tags      []string `json:"tags"`
	Image     string   `json:"image"`
	TheatreId int      `json:"theatreId"`
	ShowId    int      `json:"showId"`
	Date      string   `json:"date"`
	Time      []Seat   `json:"time"`
}
