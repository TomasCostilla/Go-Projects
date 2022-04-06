package models

type user struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Useremail string `json:"useremail"`
}

var Users = []user{
	{ID: "1", Username: "Tomas", Useremail: "tomas@gmail.com"},
	{ID: "2", Username: "Pedro", Useremail: "Pedro@gmail.com"},
	{ID: "3", Username: "Pablo", Useremail: "Pablo@gmail.com"},
}
