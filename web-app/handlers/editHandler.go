
package handlers

import (
    "net/http"
    "go-guide/web-app/utils"
    "go-guide/web-app/types"
)

func EditHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/edit/"):]
    p, err := utils.LoadPage(title)
    if err != nil {
        p = &types.Page{Title: title}
    }
    utils.RenderTemplate(w, "edit", p)

}