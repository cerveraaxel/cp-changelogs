package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const HelpMessage = `
cPanel changelogs API:

Perform a query with the full cPanel version including major, minor and build versions
for which you would like to receive changelogs for.


Basic usage:

curl api.cpanel.axelcervera.com/11.102.0.5


Using the "parent" version value is optional:

curl api.cpanel.axelververa.com/102.0.5

`

func handleRequests() {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/", root)
	r.HandleFunc("/{version}", runner)

	log.Fatal(http.ListenAndServe(":8080", r))
}

func main() {
	handleRequests()
}

// Version Endpoint function runner
func runner(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	v := args["version"]

	fmt.Printf("Endpoint received version search: %s\n", v)

	output, err := GetLogs(v)
	if err != nil {
		fmt.Fprintf(w, "Error: %v\n", err)
		fmt.Printf("Error: %v\n", err)
	} else {
		data, _ := json.MarshalIndent(output, "", "  ")
		fmt.Fprint(w, string(data))
	}

}

// Root Enddpoint function with instructions on how to use
func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, HelpMessage)
	fmt.Println("Endpoint Hit: Root")
}
