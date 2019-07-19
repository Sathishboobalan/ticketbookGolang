package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func helloworld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello")
}
func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.PathPrefix("/temp-images/").Handler(http.StripPrefix("/temp-images/", http.FileServer(http.Dir("."+"/temp-images/"))))
	// myRouter.Handle("/temp-images/", http.StripPrefix("/temp-images/", http.FileServer(http.Dir("temp-images"))))
	myRouter.HandleFunc("/", helloworld).Methods("GET")
	myRouter.HandleFunc("/login/{mail}", Login).Methods("POST")
	myRouter.HandleFunc("/showdetails", ShowDetails).Methods("POST")
	myRouter.HandleFunc("/myBookings/{userid}", MyBooking).Methods("POST")
	myRouter.HandleFunc("/user/{userid}/{fname}/{lname}/{dob}/{mobile}/{gender}/{typeoflogin}/{mail}/{password}", Newuser).Methods("POST")
	myRouter.HandleFunc("/bookevent/{userid}/{fname}/{lname}/{usermail}/{bimage}/{eventid}/{eventname}/{eventtype}/{stadiumtype}/{eventdate}/{eventlocation}/{ticketbooked}/{amountpaid}", BookEvent).Methods("POST")
	myRouter.HandleFunc("/eventbook", EventBook).Methods("POST")
	// myRouter.HandleFunc("/user/{name}", Deleteuser).Methods("DELETE")
	// myRouter.HandleFunc("/user/{name}/{mail}", Updateuser).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	fmt.Println("Started")
	InitialMigration()
	handleRequest()
}
