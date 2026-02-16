package utils

import (
	"go-guide/web-app/types"
	"html/template"
	"net/http"
)

// templates is initialized once at program startup, parsing all template files.
// template.Must is a convenience wrapper that panics when passed a non-nil error value.
// A panic is appropriate here; if the templates can't be loaded the only sensible thing to do is exit the program.
var templates = template.Must(template.ParseFiles("templates/edit.html", "templates/view.html"))

// RenderTemplate renders the specified template with the given page data.
// The http.Error function sends a specified HTTP response code (in this case "Internal Server Error") and error message.
func RenderTemplate(w http.ResponseWriter, tmpl string, p *types.Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}