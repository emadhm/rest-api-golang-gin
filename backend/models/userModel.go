package models


type Users struct {

  ID           uint		`json:"id"`
  Name         string	`json:"name"`
  Email        string `json:"email"`
  Role         string	`json:"role"`
  Password    string	`json:"password"`

}