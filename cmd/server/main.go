package main

import (
	"log"

	"github.com/anil1226/go-gRPC/internal/db"
	"github.com/anil1226/go-gRPC/internal/rocket"
)

func main() {

	err := Run()
	if err != nil {
		log.Fatal(err)
	}
}

func Run() error {

	db, err := db.NewDatabase()
	if err != nil {
		return err
	}
	serv := rocket.New(db)
	_ = serv
	return nil
}
