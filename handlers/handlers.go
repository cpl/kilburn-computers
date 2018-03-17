// Package handlers implements methods for returning lab availability information
// Uses the services package to get info from Kilburn
package handlers

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/colinx05/kilburn-computers/services"
	"github.com/colinx05/kilburn-computers/utils"
	"github.com/gorilla/mux"
)

// Router is the router used to configure the routes
var Router = mux.NewRouter()

// ConfigureHandlers configures the routes
func ConfigureHandlers() {
	// Configure routes
	Router.HandleFunc("/", index)
	Router.HandleFunc("/about", aboutPage)
	Router.HandleFunc("/{lab}", labInfoPage)

	// Configure API routes
	Router.HandleFunc("/api/list", labList)
	Router.HandleFunc("/api/{lab}", labInfoAPI)

	// This will serve files under /assets/<filename>
	Router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	http.Handle("/", Router)
}

// index displays the index page
func index(w http.ResponseWriter, r *http.Request) {
	// return HTML
	w.Header().Set("Content-Type", "text/html")

	t, _ := template.ParseFiles("templates/index.html", "templates/layout.html")
	t.ExecuteTemplate(w, "layout", struct {
		LabList []string
	}{
		LabList: services.GetLabList(),
	})
}

// aboutPage displays the about page
func aboutPage(w http.ResponseWriter, r *http.Request) {
	// return HTML
	w.Header().Set("Content-Type", "text/html")

	t, _ := template.ParseFiles("templates/about.html", "templates/layout.html")
	t.ExecuteTemplate(w, "layout", struct {
		LabList []string
	}{
		LabList: services.GetLabList(),
	})
}

// labInfoPage displays info about a particular lab
func labInfoPage(w http.ResponseWriter, r *http.Request) {
	// return HTML
	w.Header().Set("Content-Type", "text/html")

	// get lab info
	labInfo, err := services.GetLabInfo(mux.Vars(r)["lab"])

	// redirect to '/' if lab doesn't exist
	if err != nil {
		index(w, r)
		//fmt.Fprintf(w, "<script>window.location.href='/'</script>")
		return
	}

	// display lab info
	t, _ := template.ParseFiles("templates/lab.html", "templates/layout.html")
	t.ExecuteTemplate(w, "layout", struct {
		LabList    []string
		LabInfo    services.LabInfo
		UsageLevel string
	}{
		LabList:    services.GetLabList(),
		LabInfo:    labInfo,
		UsageLevel: utils.GetPercentage(labInfo.Used, labInfo.Count), // compute usage level here, probably not possible in template
	})
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
