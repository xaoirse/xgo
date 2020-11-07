package main

import (
	"flag"
	"log"

	"github.com/xaoirse/xgo/graph/model"
	_ "github.com/xaoirse/xgo/graph/plugin"
	"github.com/xaoirse/xgo/router"
)

const defaultPort = "4000"

func main() {

	port := flag.String("p", "4000", "Port")
	secret := flag.String("s", "", "Secret")
	flag.Parse()

	db := model.GetDb()
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalln("Error while closing db:", err)
		}
	}()

	router.Start(db, port, []byte(*secret))
}
