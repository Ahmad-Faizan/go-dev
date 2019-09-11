package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Ahmad-Faizan/go-dev/src/api_connection_test/app"
)

func main() {
	router := app.GetRouter()
	//start the server at port 3030
	fmt.Println("Server is live")

	log.Fatal(http.ListenAndServe(":3030", router))
}
