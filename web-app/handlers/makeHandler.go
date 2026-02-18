package handlers

import (
	"net/http"
	"go-guide/web-app/utils"
)


// wrapper function that takes a function of the above type, 
// and returns a function of type http.HandlerFunc (suitable to be passed to the function http.HandleFunc):
func MakeHandler(fn func (http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Here we will extract the page title from the Request,
        // and call the provided handler 'fn'
		m := utils.ValidPath.FindStringSubmatch(r.URL.Path)
        if m == nil {
            http.NotFound(w, r)
            return
        }
        fn(w, r, m[2])
    }
}