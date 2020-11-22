package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/nocubicles/veloturg/src/middleware"
	"github.com/nocubicles/veloturg/src/routes"
	"github.com/nocubicles/veloturg/src/utils"
)

func confirmation(w http.ResponseWriter, r *http.Request) {

	utils.Render(w, "confirmation.html", nil)
}

func main() {
	err := godotenv.Load("../.env")

	if err != nil {
		fmt.Println("cannot find .env file")
	}

	router := mux.NewRouter()
	router.HandleFunc("/confirmation", confirmation).Methods("GET")
	router.HandleFunc("/logisisse", routes.RenderSignIn).Methods("GET")
	router.HandleFunc("/logisisse", routes.SendSignInEmail).Methods("POST")
	router.HandleFunc("/kuulutus", middleware.CheckIsUsedLoggedIn(routes.RenderAdForm)).Methods("GET")
	router.HandleFunc("/kuulutus", middleware.CheckIsUsedLoggedIn(routes.ReceiveAdForm)).Methods("POST")
	router.HandleFunc("/kuulutus/{adID}", routes.RenderAd).Methods("GET")
	router.HandleFunc("/", routes.RenderHome).Methods("GET")

	log.Println("Listening..")
	err = http.ListenAndServe(":3000", router)

	if err != nil {
		log.Fatal(err)
	}
}
