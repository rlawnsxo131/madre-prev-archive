package main

import (
	"log"

	"github.com/rlawnsxo131/madre-server-v2/database"
	"github.com/rlawnsxo131/madre-server-v2/server"
)

func main() {
	db, err := database.GetDatabaseInstance()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	defer db.DB.Close()
	database.ExcuteInitSQL(db.DB)
	s := server.New()
	s.Start()
}

func init() {
	log.Println("init main")
}
