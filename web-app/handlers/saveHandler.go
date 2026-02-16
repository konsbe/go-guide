
package handlers

import (
    "net/http"
    "go-guide/web-app/utils"
    "go-guide/web-app/types"
)

// The function saveHandler will handle the submission of forms located on the edit pages.
// After uncommenting the related line in main, let's implement the handler:
func SaveHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/save/"):]
    body := r.FormValue("body")
    p := &types.Page{Title: title, Body: []byte(body)}
    utils.SavePage(p)
    http.Redirect(w, r, "/view/"+title, http.StatusFound)
}