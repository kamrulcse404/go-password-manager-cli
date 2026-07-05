package models

type Password struct {
	ID 			int 	`json:"id"`
	Service 	string 	`json:"service"`
	Username 	string  `json:"username"`
	Password    string  `Json:"password"`
}