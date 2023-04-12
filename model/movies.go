package model

type Movie struct {
	Id        int    `json:"id" gorm:"primary_key"`
	Name      string `json:"name"`
	Language  string `json:"language"`
	Image     string `json:"image"`
	Headimage string `json:"headimage"`
	Tags      string `json:"tags"`
	Comment   string `json:"comment"`
}
