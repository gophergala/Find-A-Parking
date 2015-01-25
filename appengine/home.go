package home

import (
	"html/template"
	"net/http"
	"appengine"
    "appengine/user"
    "appengine/datastore"
    "time"
    "fmt"
    "encoding/json"
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
	u := user.Current(c)
	if u == nil {
		url, err := user.LoginURL(c, r.URL.String())
        if err != nil {
            http.Redirect(w, r, "/", http.StatusMovedPermanently)
        }
        w.Header().Set("Location", url)
        w.WriteHeader(http.StatusFound)
        return
	}
	d := map[string]interface{}{"Titulo": "Home"}
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
	u := user.Current(c)
	
	if u == nil {
		url, err := user.LoginURL(c, r.URL.String())
        if err != nil {
            http.Redirect(w, r, "/", http.StatusMovedPermanently)
        }
        w.Header().Set("Location", url)
        w.WriteHeader(http.StatusFound)
        return
	}

	d := map[string]interface{}{"Titulo": "Rent"}
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




