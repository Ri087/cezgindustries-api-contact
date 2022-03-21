package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Contact struct {
	Id          int
	Name        string
	Age         int
	Post        string
	Mail        string
	UrlImage    string
	Cities      string
	Competences []string
}
type Contacts struct {
	CoFounder      []Contact
	HangmanWeb     []Contact
	GroupieTracker []Contact
}

func AllContact(w http.ResponseWriter, r *http.Request) {
	Jerem := Contact{Id: 1, Name: "Jeremy", Age: 18, Post: "Dev Back & Front", Mail: "jeremy.dura@ynov.com", UrlImage: "https://cezgindustries-api-contact.herokuapp.com/ressources/img/photo_jeremy.jpg", Cities: "Lyon", Competences: []string{"Golang", "CSS", "web", "Font End Developement", "HTML", "Back End Developement"}}
	Benjamin := Contact{Id: 2, Name: "Benjamin", Age: 18, Post: "Dev Backend", Mail: "benjamin.vernet@ynov.com", UrlImage: "https://cezgindustries-api-contact.herokuapp.com/ressources/img/photo_benjamin.png", Cities: "Lyon", Competences: []string{"Golang", "Font End Developement", "HTML", "CSS", "web", "Back End Developement", "Main ADC"}}
	Sean := Contact{Id: 3, Name: "Sean", Age: 23, Post: "Dev FrontEnd", Mail: "sean.cappe@ynov.com", UrlImage: "https://cezgindustries-api-contact.herokuapp.com/ressources/img/photo_sean.jpg", Cities: "Lyon", Competences: []string{"HTML", "CSS", "Back End Developement", "Web", "Font End Developement", "Garen IRL"}}
	Contacts := Contacts{
		CoFounder: []Contact{
			Jerem,
			Benjamin,
		},
		HangmanWeb: []Contact{
			Jerem,
			Benjamin,
		},
		GroupieTracker: []Contact{
			Jerem,
			Benjamin,
			Sean,
		},
	}
	json.NewEncoder(w).Encode(Contacts)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HomePage !")
}
func handleRequest() {
	port := os.Getenv("PORT")
	http.HandleFunc("/", AllContact)
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/ressources/", http.StripPrefix("/ressources/", fileServer))
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
func main() {
	handleRequest()

}
