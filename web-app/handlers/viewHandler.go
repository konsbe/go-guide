package handlers

import (
    "net/http"
	"go-guide/web-app/utils"
)

// First, this function extracts the page title from r.URL.Path, 
// the path component of the request URL. The Path is re-sliced with [len("/view/"):] 
// to drop the leading "/view/" component of the request path. 
// This is because the path will invariably begin with "/view/", 
// which is not part of the page's title.

// What if you visit /view/APageThatDoesntExist? 
// You'll see a page containing HTML. 
// This is because it ignores the error return value from loadPage 
// and continues to try and fill out the template with no data. Instead, 
// if the requested Page doesn't exist, it should redirect the client to the edit Page so the content may be created:
func ViewHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/view/"):]
    p, err := utils.LoadPage(title)
    if err != nil {
        http.Redirect(w, r, "/edit/"+title, http.StatusFound)
        return
    }
    utils.RenderTemplate(w, "view", p)
}