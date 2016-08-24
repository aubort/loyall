package main

import ( 
    "net/http"
    "fmt"
    
 //   //recaptcha imports
 //   "encoding/json"
	// "io/ioutil"
	// //"net/http"
	// "net/url"
	// "time"
    
    
    "appengine"
    "appengine/mail"
    // "appengine/urlfetch"
    
    // "./recaptcha"
    // "appengine/log"
)

func init() {
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.HandleFunc("/contactus/", handleContactus)
}

func handleContactus(w http.ResponseWriter, r *http.Request) {
    
    ctx := appengine.NewContext(r)
    
    name := r.FormValue("name")
    receiver := r.FormValue("email")
    request := r.FormValue("request")
    
    ctx.Infof("Name: %v", name)
    
    msg := &mail.Message{
		Sender:  "Loyall.ch Info <info@loyall.ch>",
		To:      []string{receiver, "info@loyall.ch"},
		Subject: "Your request to Loyall",
		Body:    fmt.Sprintf("Dear %v %v", name, request),
	}
	
	ctx.Infof("Trying to send message: %v", msg)
	
	if err := mail.Send(ctx, msg); err != nil {
	    ctx.Errorf("Couldn't send email: %v", err)
	} else {
	    ctx.Infof("An email has been sent to: %v", receiver)
	}
    return
}

// BELOW To be replaced when I figure out how to import friggin packages!

// R type represents an object of Recaptcha and has public property Secret,
// which is secret obtained from google recaptcha tool admin interface
// type R struct {
// 	Secret    string
// 	lastError []string
// }

// // Struct for parsing json in google's response
// type googleResponse struct {
// 	Success    bool
// 	ErrorCodes []string `json:"error-codes"`
// }

// // url to post submitted re-captcha response to
// var postURL = "https://www.google.com/recaptcha/api/siteverify"

// // Verify method, verifies if current request have valid re-captcha response and returns true or false
// // This method also records any errors in validation.
// // These errors can be received by calling LastError() method.
// func (r *R) Verify(req *http.Request) bool {
// 	r.lastError = make([]string, 1)
// 	response := req.FormValue("g-recaptcha-response")
// 	// client := &http.Client{Timeout: 20 * time.Second}
// 	ctx := appengine.NewContext(req)
//     client := urlfetch.Client(ctx)
// 	resp, err := client.Post(postURL,
// 		url.Values{"secret": {r.Secret}, "response": {response}})
// 	if err != nil {
// 		r.lastError = append(r.lastError, err.Error())
// 		return false
// 	}
// 	defer resp.Body.Close()
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		r.lastError = append(r.lastError, err.Error())
// 		return false
// 	}
// 	gr := new(googleResponse)
// 	err = json.Unmarshal(body, gr)
// 	if err != nil {
// 		r.lastError = append(r.lastError, err.Error())
// 		return false
// 	}
// 	if !gr.Success {
// 		r.lastError = append(r.lastError, gr.ErrorCodes...)
// 	}
// 	return gr.Success
// }

// // LastError returns errors occurred in last re-captcha validation attempt
// func (r R) LastError() []string {
// 	return r.lastError
// }
