// building my first web server in golang!

package main

// using packages needed
import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// this is our entrypoint to our new application
func main() {

	//this is a multiplexer to route incomming http requests to our handlers in code
	m := http.NewServeMux()

	// handle any request to '/' root with our handlepage function
	m.HandleFunc("/", handlePage)
	//now define a server
	const addr = ":8000"
	//os.Getenv("PORT", "8000")

	srv := http.Server{
		Handler:      m,
		Addr:         addr,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	// print message to the console
	fmt.Println("server started on port", addr)

	// this will flag an error if somethng goes wrong
	err := srv.ListenAndServe()
	log.Fatal(err)
}

// creating a handler function
func handlePage(w http.ResponseWriter, r *http.Request) {
	// setting response header
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	const page = `<html>
	<body>
	<p> Hi Docker, I pushed a new version! </p>
	</body>
	</html>
	`
	// 200: Response code
	w.WriteHeader(200)
	// write the actual html data
	w.Write([]byte(page))
}
