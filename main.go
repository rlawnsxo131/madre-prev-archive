package main

import (
	"log"

	"github.com/rlawnsxo131/madre-server-v2/database"
	"github.com/rlawnsxo131/madre-server-v2/server"
)

func main() {
	db, err := database.GetDB()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	database.ExcuteInitSQL(db)
	defer db.Close()
	s := server.New(db)
	s.Start()
}

func init() {
	log.Println("init main")
}
