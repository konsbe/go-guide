
package handlers

import (
    "net/http"
    "go-guide/web-app/utils"
    "go-guide/web-app/types"
)

// The returned function is called a closure because it encloses values defined outside of it. 
// In this case, the variable fn (the single argument to makeHandler) is enclosed by the closure. 
// The variable fn will be one of our save, edit, or view handlers.

// The function saveHandler will handle the submission of forms located on the edit pages.
// After uncommenting the related line in main, let's implement the handler:
func SaveHandler(w http.ResponseWriter, r *http.Request, title string) {

    body := r.FormValue("body")
    p := &types.Page{Title: title, Body: []byte(body)}
    err := utils.SavePage(p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    http.Redirect(w, r, "/view/"+title, http.StatusFound)
}