package main

import (
	"fmt"
	"net/http"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

func init(){
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/words", handleWords)
}

func handleIndex(res http.ResponseWriter, req http.Request){

}

type word struct {
	Term string
	Definition string
}

func handleWords(res http.ResponseWriter, req http.Request){
	if req.Method == "POST" {
		term := req.FormValue("word")
		definition := req.FormValue("definition")
		ctx := appengine.NewContext(req)
		key := datastore.NewIncompleteKey(ctx, "word", nil)
		entity := word{
			Term: term,
			Definition: definition,
		}
		_, err := datastore.Put(ctx, key, entity)
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}

	}
	fmt.Fprintln(res, `
	<form method= "POST" action="/words">
		<input type= "text" name= "term">
		<textarea name= "definition">
		<input type= "submit">
	</form>
	`)
}
