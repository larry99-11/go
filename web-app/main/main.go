package main

/*
TODO:
- Have a route for /
-  Have a custom route for i.e /main or /login etc...
*/
import (
	"fmt"
	"log"
	"net/http"

	"github.com/larry-99/webGo/pkg/handlers"
)

const portNumber = ":8080"
/*
http.HandlerFunc to convert simple functions into handlers.
The function http.HandleFunc to tell the server which function to call to handle a request to the server.
*/
func main() {
	home := handlers.Home()
	// need to handle multiple path requests
	http.HandleFunc("/", home )

	fmt.Println("Starting server! Listening on port 8080")
	log.Fatal(http.ListenAndServe(portNumber, nil))
}
