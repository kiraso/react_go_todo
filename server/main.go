package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/kiraso/react_go_todo/router"
)

func main() {
	r :=  router.Router()
	fmt.Println("Server started on: http://localhost:9000")
	http.ListenAndServe(":9000",r)
	log.Fatal(http.ListenAndServe(":9000", router.Router()))
}
