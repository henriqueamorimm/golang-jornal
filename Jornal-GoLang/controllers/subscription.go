package controllers
import (
	"websitego/db" 
)

type Subscription struct {
	Nome  string
	Email string
}

func Create(nome string, email string) error {
	s := Subscription{Nome: nome, Email: email}

	return db.Insert("jornal", s)
}
