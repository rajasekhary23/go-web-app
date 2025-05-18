package main

import (
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	// Render the home html page from static folder
	log.Println("Accessed path:", r.URL.Path)
	http.ServeFile(w, r, "static/home.html")
}

func coursePage(w http.ResponseWriter, r *http.Request) {
	// Render the course html page
	log.Println("Accessed path:", r.URL.Path)
	http.ServeFile(w, r, "static/courses.html")
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
	// Render the about html page
	log.Println("Accessed path:", r.URL.Path)
	http.ServeFile(w, r, "static/about.html")
}

func contactPage(w http.ResponseWriter, r *http.Request) {
	// Render the contact html page
	log.Println("Accessed path:", r.URL.Path)
	http.ServeFile(w, r, "static/contact.html")
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/home", homePage)
	http.HandleFunc("/courses", coursePage)
	http.HandleFunc("/about", aboutPage)
	http.HandleFunc("/contact", contactPage)

	// Serve static assets (css, js, images) from /static/ URL path
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	
	log.Println("Starting server on :8081")
	
	err := http.ListenAndServe("0.0.0.0:8081", nil)
	if err != nil {
		log.Fatal(err)
	}
}
