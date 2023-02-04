package models


type Users struct {

  ID           uint		`json:"id"`
  Name         string	`json:"name"`
  Role         string	`json:"role"`
  Passsword    string	`json:"password"`

}