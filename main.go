package main

import ( 
    "net/http"
    "fmt"
    
 //   //recaptcha imports
 //   "encoding/json"
// 	"io/ioutil"
	// //"net/http"
	// "net/url"
	// "time"
	"strings"
    
    
    "appengine"
    // "appengine/mail"
    "appengine/urlfetch"
    
    // "./recaptcha"
    // "appengine/log"
)

var (
    grooveApiKey = "31ec9b652af2605b87e51ca4acaed7e34ab2274cd588e5ab6fe4afe233816cdf"
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
    
//     msg := &mail.Message{
// 		Sender:  "Loyall.ch Info <info@loyall.ch>",
// 		To:      []string{receiver, "info@loyall.ch"},
// 		Subject: "Your request to Loyall",
// 		Body:    fmt.Sprintf("Dear %v %v", name, request),
// 	}
	
// 	ctx.Infof("Trying to send message: %v", msg)
	
// 	if err := mail.Send(ctx, msg); err != nil {
// 	    ctx.Errorf("Couldn't send email: %v", err)
// 	} else {
// 	    ctx.Infof("An email has been sent to: %v", receiver)
// 	}
    
    createGrooveTicket(w, ctx, receiver, request)	
}

func createGrooveTicket(w http.ResponseWriter, ctx appengine.Context, sender string, request string){
    ctx.Infof("Trying to create a Groove Ticket")
    client := urlfetch.Client(ctx)
    // url := "https://api.groovehq.com/v1/me?access_token=31ec9b652af2605b87e51ca4acaed7e34ab2274cd588e5ab6fe4afe233816cdf"

    // resp, err := client.Get(url)
    //resp, err := client.Post("https://api.groovehq.com/v1/me?access_token=31ec9b652af2605b87e51ca4acaed7e34ab2274cd588e5ab6fe4afe233816cdf", "application/json", nil)
    //var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast."}`)
    
    json := `{"body":"`+request+`", "from":"`+sender+`", "to":"info@loyall.ch"}`
    buf := strings.NewReader(json)
    
    req, err := http.NewRequest("POST", "https://api.groovehq.com/v1/tickets", buf)
    
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
	req.Header.Set("Authorization", "Bearer 31ec9b652af2605b87e51ca4acaed7e34ab2274cd588e5ab6fe4afe233816cdf")
	req.Header.Set("Content-Type", "application/json")
   
	resp, err := client.Do(req)
	ctx.Infof("Response: %v", resp)

    
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    fmt.Fprintf(w, "Hi there, I love %s!", resp.Status)
}