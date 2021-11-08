/* making this file exec */
package main

/* imports */
import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

/* main func */
func main() {

	/* first http handler */
	http.HandleFunc("/index", func(r http.ResponseWriter, rq *http.Request) {
		log.Println("Request")
		body, err := ioutil.ReadAll(rq.Body)

		if err != nil {
			http.Error(r, "OOPS", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(r, "Hello %s", body)
	})

	/* second example handler */
	http.HandleFunc("/bye", func(http.ResponseWriter, *http.Request) {
		log.Println("End..")
	})

	http.ListenAndServe(":8080", nil)
}
