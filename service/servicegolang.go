package service

import (
	//"Myself/entity"
	"database/sql"
	"fmt"
	"time"

	"gitlab.com/pragmaticreviews/golang-gin-poc/entity"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "golang"
	password = "golang"
	dbname   = "golang"
)

// ServiceGolang ..
type ServiceGolang interface {
	AddGolang(entity.Golang) int
	GetAllGolang() []entity.Golang
}

type serviceGolang struct {
	users []entity.Golang
}

// New ServiceGolang ..
func New() ServiceGolang {
	return &serviceGolang{}
}

func (service *serviceGolang) AddGolang(golang entity.Golang) int {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	fmt.Print(err)
	if err != nil {
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected1!")
	db.SetConnMaxLifetime(time.Second * 5)
	db.SetMaxIdleConns(100)
	db.SetMaxOpenConns(100)
	defer db.Close()
	fmt.Print(golang)

	query := "INSERT INTO customer(id, name, age, phone) values($1,$2,$3,$4) RETURNING golang_id"
	id := 0
	err = db.QueryRow(query, golang.ID, golang.Name, golang.Age, golang.Phone).Scan(&id)
	if err != nil {
		panic(err.Error())
	}
	return id
}

func (service *serviceGolang) GetAllGolang() []entity.Golang {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	fmt.Print(err)
	if err != nil {
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
	db.SetConnMaxLifetime(time.Second * 5)
	db.SetMaxIdleConns(100)
	db.SetMaxOpenConns(100)
	defer db.Close()

	query := "Select * from customer"
	rows, err2 := db.Query(query)
	if err2 != nil {
		panic(err.Error())
	}

	golangs := []entity.Golang{}

	for rows.Next() {
		var g entity.Golang
		err = rows.Scan(&g.ID, &g.Name, &g.Age, &g.Phone)
		golangs = append(golangs, g)
	}
	return golangs

}
