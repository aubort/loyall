package main

import ( 
    "net/http"
    "fmt"
    
    "appengine"
    "appengine/mail"
    
    // "./recaptcha"
    "github.com/haisum/recaptcha"
    // "appengine/log"
)

func init() {
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.HandleFunc("/contactus/", handleContactus)
}

func handleContactus(w http.ResponseWriter, r *http.Request) {
    
    ctx := appengine.NewContext(r)
    
	re := recaptcha.R{
		Secret: "6LeURygTAAAAAApwnTDNSQ7fSHAf5vfCPCPuAWlJ",
	}
    
    isValid := re.Verify(*r)
	if isValid {
		ctx.Infof("Valid")
	} else {
		ctx.Errorf("Invalid! These errors ocurred: %v", re.LastError())
	}

    
    name := r.FormValue("name")
    receiver := r.FormValue("email")
    request := r.FormValue("request")
    
    msg := &mail.Message{
		Sender:  "Loyall.ch Info <info@loyall.ch>",
		To:      []string{receiver},
		Subject: "Your request to Loyall",
		Body:    fmt.Sprintf("Dear %v %v", name, request),
	}
	
	if err := mail.Send(ctx, msg); err != nil {
	    ctx.Errorf("Couldn't send email: %v", err)
	} else {
	    ctx.Infof("An email has been sent to: %v", receiver)
	}
    return
}