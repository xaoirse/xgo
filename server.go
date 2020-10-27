package main

import (
	"flag"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/xaoirse/xgo/graph/model"
	_ "github.com/xaoirse/xgo/graph/plugin"
	"github.com/xaoirse/xgo/router"
)

func getDb() *gorm.DB {
	db, err := gorm.Open("sqlite3", "data.db")
	if err != nil {
		log.Fatalln("Error in opening db:", err)
	}
	db.AutoMigrate(model.GetModels()...)
	return db
}

const defaultPort = "4000"

func main() {

	port := flag.String("p", "4000", "Port")
	secret := flag.String("s", "", "Secret")
	flag.Parse()

	db := getDb()
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalln("Error when closing db:", err)
		}
	}()

	router.Start(db, port, []byte(*secret))
}
