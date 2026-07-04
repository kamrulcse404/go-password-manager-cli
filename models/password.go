package models

type Password struct {
	ID 			int 	`json:"id"`
	Service 	string 	`json:"title"`
	Username 	string  `json:"username"`
	Password    string  `Json:"password"`
}