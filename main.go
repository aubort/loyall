package main

import ( 
    "net/http"
    
    "appengine"
)

func init() {
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.HandleFunc("/contactus/", handleContactus)
}

func handleContactus(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)
    content := r.FormValue("firstname")
    c.Debugf("The message: %s", content)
    return
}