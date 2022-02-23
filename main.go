package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/rlawnsxo131/madre-server-v2/database"
	"github.com/rlawnsxo131/madre-server-v2/server"
)

func main() {
	db := database.GetDB()
	database.ExcuteInitSQL(db)
	// db.MustExec("INSERT INTO user (email, username, display_name, photo_url) VALUES(?, ?, ?, ?)", "juntae", "juntae", "juntae", "juntae")
	// db.MustExec("INSERT INTO data (user_id, file_url, title, description) VALUES(?, ?, ?, ?)", "user_id", "https://devlog.juntae.kim", "title", "description")
	defer db.Close()
	s := server.New(db)
	s.Start()
}

func init() {
	log.Println("init main")
}
