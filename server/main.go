package main

import(
	"fmt"
	"log"
	"net/http"
	"github.com/Meruyert98/golang-react-todo/router"
)

func main(){
	r := router.Router()
	fmt.Println("starting the server on port 8081...")

	log.Fatal(http.ListenAndServe(":8081", r))
}
	
