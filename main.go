package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/rlawnsxo131/madre-server-v2/database"
	"github.com/rlawnsxo131/madre-server-v2/server"
)

func main() {
	db, err := database.GetDB()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	database.ExcuteInitSQL(db)

	// db.MustExec("INSERT INTO user(uuid, email, username, display_name, photo_url) VALUES(?, ?, ?, ?, ?)", "uuid", "juntae", "juntae", "juntae", "juntae")
	// db.MustExec("INSERT INTO data (uuid, user_id, file_url, title, description) VALUES(?, ?, ?, ?, ?)", "uuid", 1, "https://devlog.juntae.kim", "title", "description")

	defer db.Close()
	s := server.New(db)
	s.Start()
}

func init() {
	log.Println("init main")
}
