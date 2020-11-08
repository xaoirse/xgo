package main

import (
	"flag"
	"log"

	"github.com/xaoirse/xgo/repository"
	// It's for init hook mutatehooks
	_ "github.com/xaoirse/xgo/graph/plugin"
	"github.com/xaoirse/xgo/router"
)

const defaultPort = "4000"

func main() {

	port := flag.String("p", "4000", "Port")
	secret := flag.String("s", "", "Secret")
	flag.Parse()

	repository.InitDB()
	db := repository.GetDB()
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalln("Error while closing db:", err)
		}
	}()

	router.Start(db, port, []byte(*secret))
}
