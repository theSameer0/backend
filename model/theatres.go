package model

type Theatre struct {
	Id       int    `json:"id" gorm:"primary_key"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Image    string `json:"image"`
	City     string `json:"city"`
	Screen   int    `json:"screen"`
}
