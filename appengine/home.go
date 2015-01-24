package home

import (
	"html/template"
	"net/http"
	"appengine"
    "appengine/user"
)

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

func rent() {
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