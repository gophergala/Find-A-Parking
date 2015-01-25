package home

import (
	"html/template"
	"net/http"
	"appengine"
    "appengine/datastore"
    "time"
    "fmt"
    "encoding/json"
    "appengine/channel"
)


type Response map[string]interface{}

func (r Response) String() (s string) {
        b, err := json.Marshal(r)
        if err != nil {
                s = ""
                return
        }
        s = string(b)
        return
}

type Parking struct {
	Owner string
	Mail string
	Price float32
	Location appengine.GeoPoint
}

type Transactions struct {
	Park Parking
	Begin time.Time
	End time.Time
	Customer string
}

func init() {
	http.HandleFunc("/",index)
	http.HandleFunc("/home",home)
	http.HandleFunc("/rent",rent)

}

func index(w http.ResponseWriter, r *http.Request) {
	d := map[string]interface{}{"Titulo": "Find A Park"}
	t, err := template.ParseFiles("templates/index.html", "templates/base.html")
  	if err != nil {
    	http.Error(w, err.Error(), http.StatusInternalServerError)
  	} else{
    	err := t.ExecuteTemplate(w,"base", d)
	    if err != nil {
	      http.Error(w, err.Error(), http.StatusInternalServerError)
	    } 
  	}
}

func home(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	

	token, err := channel.Create(c, "customer")
	if err != nil {
	    token = ""
	}

	d := map[string]interface{}{"Titulo": "Home", "Token": token}
  	t, err := template.ParseFiles("templates/home.html", "templates/base.html")
  	if err != nil {
    	http.Error(w, err.Error(), http.StatusInternalServerError)
  	} else{
	    err := t.ExecuteTemplate(w,"base", d)
	    if err != nil {
	      http.Error(w, err.Error(), http.StatusInternalServerError)
	    } 
  	}
}

func rent(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	parking := &Parking{
		Owner:	"",
		Mail:	"",
		Price:	25.6,
	}

	d := map[string]interface{}{"Titulo": "Rent"}

	channel.SendJSON(c, "customer", parking)

  	t, err := template.ParseFiles("templates/renta.html", "templates/base.html")
  	if err != nil {
    	http.Error(w, err.Error(), http.StatusInternalServerError)
  	} else{
	    err := t.ExecuteTemplate(w,"base", d)
	    if err != nil {
	      http.Error(w, err.Error(), http.StatusInternalServerError)
	    } 
  	}	
}

func createParking(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	parking := &Parking{
		Owner:	"",
		Mail:	"",
		Price:	25.6,
	}

	key := datastore.NewIncompleteKey(c, "Parking", nil);
	_, err := datastore.Put(c,key,parking)
	if err != nil {
	    http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, Response{"success": err == nil})
    return
}




