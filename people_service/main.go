package main

import (
	"fmt"

	"github.com/comrade-pavel-sergeevich/TBD/people_service/service"
	"github.com/comrade-pavel-sergeevich/TBD/people_service/service/store"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	db := store.NewStore("admin:1234@tcp(localhost:3306)/people_service?parseTime=true&multiStatements=true")
	serv := service.Service{
		Store: db,
		Tax:   &t{},
	}

	fmt.Println(serv.ListPeople())
	fmt.Println(serv.GetPeopleByID(3)) // example

	fmt.Println("__________")
	serv.Store.CloseConn()
}

// simple fake tax realization
type t struct{}

func (t *t) GetTaxStatusByID(id int) (string, error) {
	return "", nil
}
