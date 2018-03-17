// Package handlers implements methods for returning lab availability information
// Uses the services package to get info from Kilburn
package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/colinx05/comp-availability/services"
	"github.com/gorilla/mux"
)

// Router is the router used to configure the routes
var Router = mux.NewRouter()

// ConfigureHandlers configures the routes
func ConfigureHandlers() {
	Router.HandleFunc("/", index)
	Router.HandleFunc("/{lab}", labInfo)

	Router.HandleFunc("/api/", labList)
	Router.HandleFunc("/api/{lab}", labInfoAPI)
}

// index page
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello!</h1>")
	fmt.Fprintf(w, "<h4>The available labs are:</h4>")
	fmt.Fprintf(w, "<ul>")
	for _, lab := range services.GetLabList() {
		fmt.Fprintf(w, "<li><a href='/%s'>%s</li>", lab, lab)
	}
	fmt.Fprintf(w, "</ul>")
}

// labInfo returns info about a page
func labInfo(w http.ResponseWriter, r *http.Request) {
	labInfo, err := services.GetLabInfo(mux.Vars(r)["lab"])

	if err != nil {
		fmt.Fprintf(w, "<h1>Lab doesn't exist!</h1>")
		return
	}
	fmt.Fprintf(w, "<h1>Lab: %s</h1>", labInfo.Name)
	fmt.Fprintf(w, "<p>Location: %s</p>", labInfo.Location)
	fmt.Fprintf(w, "<p>Description: %s</p>", labInfo.Description)
	fmt.Fprintf(w, "<p>Usage: %d (%d/%d)</p>", labInfo.Free/labInfo.Used, labInfo.Free, labInfo.Used)
	fmt.Fprintf(w, "<p>Computers:</p>")
	fmt.Fprintf(w, "<ul>")
	for _, computer := range labInfo.Computers {
		if computer.Used {
			fmt.Fprintf(w, "<li>%s: %s</li>", computer.Name, "used")
		} else {
			fmt.Fprintf(w, "<li>%s: %s</li>", computer.Name, "free")
		}
	}
	fmt.Fprintf(w, "</ul>")
}

// labList returns the list of labs as JSON
func labList(w http.ResponseWriter, r *http.Request) {
	// return JSON
	w.Header().Set("Content-Type", "application/json")

	json, _ := json.Marshal(struct {
		Name []string `json:"lab_list"`
	}{
		services.GetLabList(),
	})

	w.Write(json)
}

// lab returns info about a page as JSON
func labInfoAPI(w http.ResponseWriter, r *http.Request) {
	// return JSON
	w.Header().Set("Content-Type", "application/json")

	// retreive lab information
	labInfo, err := services.GetLabInfo(mux.Vars(r)["lab"])

	if err != nil {
		json, _ := json.Marshal(struct {
			Error string `json:"error"`
		}{
			err.Error(),
		})

		w.Write(json)
		return
	}
	json, _ := json.Marshal(labInfo)

	w.Write(json)
}
