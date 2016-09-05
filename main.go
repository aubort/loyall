package main

import ( 
    "net/http"
    "encoding/json"
    // "strings"
    // "fmt"
	
    "bytes"
    
    "appengine"
    "appengine/urlfetch"
    
)

var (
    grooveApiKey = "31ec9b652af2605b87e51ca4acaed7e34ab2274cd588e5ab6fe4afe233816cdf"
)

type GrooveTicket struct {
		From string                 `json:"from"`
		To string                   `json:"to"`
		Subject string              `json:"subject"`
		Name string                 `json:"name"`
		Email string                `json:"email"`
		Body string                 `json:"body"`
		SendCopyToCustomer bool     `json:"send_copy_to_customer"`
}


// based on this: http://capotej.com/blog/2013/10/07/golang-http-handlers-as-middleware/

func OurLoggingHandler(h http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    // fmt.Println(*r.URL)
    //c := appengine.NewContext(r)
    
    h.ServeHTTP(w, r)
    
  })
}

func init() {
    
    fileHandler := http.FileServer(http.Dir("public"))
    wrappedHandler := OurLoggingHandler(fileHandler)
    http.Handle("/", wrappedHandler)
	http.HandleFunc("/contactus/", HandleContactus)
}



func HandleContactus(w http.ResponseWriter, r *http.Request) {
    
    c := appengine.NewContext(r)
    
    name := r.FormValue("name")
    receiver := r.FormValue("email")
    request := r.FormValue("request")
    
    c.Infof("Received values from form submit name: %v, email: %v, request: %v", name, receiver, request)
    
    ticket := GrooveTicket{receiver, "info@loyall.ch", "Your message to Loyall", name, receiver, request, true}
    b := new(bytes.Buffer)
    json.NewEncoder(b).Encode(ticket)
    
    req, err := http.NewRequest("POST", "https://api.groovehq.com/v1/tickets", b)
    
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    req.Header.Set("Authorization", "Bearer 31ec9b652af2605b87e51ca4acaed7e34ab2274cd588e5ab6fe4afe233816cdf")
	req.Header.Set("Content-Type", "application/json")
	
	client := urlfetch.Client(c)
	resp, err := client.Do(req)
	defer resp.Body.Close()
	
	if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    if resp.StatusCode != 201 {
        http.Error(w, "Error while creating ticket:" + resp.Status, resp.StatusCode)
        c.Errorf("Error while creating ticket: %v", resp.Status)
        return
    }
    
    // body, _ := ioutil.ReadAll(resp.Body)
    // c.Debugf(string(body))
    
    c.Infof("Groove ticket creation responded with: %s", resp.Status)
    
//     msg := &mail.Message{
// 		Sender:  "Loyall.ch Info <info@loyall.ch>",
// 		To:      []string{receiver, "info@loyall.ch"},
// 		Subject: "Your request to Loyall",
// 		Body:    fmt.Sprintf("Dear %v %v", name, request),
// 	}
	
// 	c.Infof("Trying to send message: %v", msg)
	
// 	if err := mail.Send(c, msg); err != nil {
// 	    c.Errorf("Couldn't send email: %v", err)
// 	} else {
// 	    c.Infof("An email has been sent to: %v", receiver)
// 	}
    // http.Redirect(w, r, "http://www.google.com", 301)
    
    // createGrooveTicket(w, c, receiver, request, name)	
    
}

func createGrooveTicket(w http.ResponseWriter, c appengine.Context, sender string, request string, name string){
    
    // url := "https://api.groovehq.com/v1/me?access_token=31ec9b652af2605b87e51ca4acaed7e34ab2274cd588e5ab6fe4afe233816cdf"

    // resp, err := client.Get(url)
    //resp, err := client.Post("https://api.groovehq.com/v1/me?access_token=31ec9b652af2605b87e51ca4acaed7e34ab2274cd588e5ab6fe4afe233816cdf", "application/json", nil)
    //var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast."}`)
    
    
    
    // ticket := GrooveTicket{sender, "info@loyall.ch", "Your message to Loyall", name, sender, request, true}
    // b := new(bytes.Buffer)
    // json.NewEncoder(b).Encode(ticket)
    
    // req, err := http.NewRequest("POST", "https://api.groovehq.com/v1/tickets", b)
    
    
    // json := `{"body":"`+request+`",
    //     "from":"`+sender+`",
    //     "to":"info@loyall.ch", 
    //     "subject":"Your message to Loyall",
    //     "name":"`+name+`",
    //     "email": "`+sender+`",
    //     "send_copy_to_customer": true,
    //     "body": "`+request+`"}`
    
    // c.Infof("Trying to create a Groove Ticket with following input:%v", json)
    
    // 
    // if err != nil {
    //     http.Error(w, err.Error(), http.StatusInternalServerError)
    //     return
    // }
    
// 	req.Header.Set("Authorization", "Bearer 31ec9b652af2605b87e51ca4acaed7e34ab2274cd588e5ab6fe4afe233816cdf")
// 	req.Header.Set("Content-Type", "application/json")
	
// 	client := urlfetch.Client(c)
// 	resp, err := client.Do(req)
	
//     defer resp.Body.Close()
    
    // if err != nil {
    //     http.Error(w, err.Error(), http.StatusInternalServerError)
    //     return
    // }
    
    // if resp.StatusCode != 201 {
    //     http.Error(w, "Error while creating ticket:" + resp.Status, resp.StatusCode)
    //     c.Errorf("Error while creating ticket: %v", resp.Status)
    //     return
    // }
    
    // body, _ := ioutil.ReadAll(resp.Body)
    // c.Debugf(string(body))
    
    // c.Infof("Groove ticket creation responded with: %s", resp.Status)
}