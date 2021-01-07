package main

import(
	"fmt"
	"log"
	"net/http"
)

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	
	retrun http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("Checking to see if Authorized header is set...")

		if val, ok := r.Header["Authorized"]; ok {
			fmt.Println(val)
			if val[0] == "true" {
				fmt.Println("Header is set! We can serve content!")
				endpoint(w, r)
			}
		} else {
			fmt.Println("Not Authorized!!")
			fmt.Fprintf(w, "Not Authorized!!")
		}
	})
}


func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: HomePage")
	fmt.fPrintf(w, "Welcome to the home page")
}

func handleRequests() {
	http.Handle("/", isAuthorized(homepage))
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() [
	handleRequests()
]