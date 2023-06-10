package main

import (
	// "MyGoWebserver/router"
	// "MyGoWebserver/services"
	// "MyGoWebserver/utils"
	"log"
	"net/http"

	"github.com/anandhu003/MyGoWebserver/router"
	"github.com/anandhu003/MyGoWebserver/services"
	"github.com/anandhu003/MyGoWebserver/utils"
)

func main() {
	log.Println("IN MAIN APP")
	var dbconn = utils.Getconnection()
	services.SetDB(dbconn)
	var approuter = router.CreateRouter()

	log.Println("❎ LISTENING ON PORT 8000 ❎ ")
	log.Fatal(http.ListenAndServe(":8000", approuter))

}
