package utils

import (
    "net/http"
    "regexp"
    "errors"
)

// As you may have observed, this program has a serious security flaw:
// a user can supply an arbitrary path to be read/written on the server. 
// To mitigate this, we can write a function to validate the title with a regular expression.

// First, add "regexp" to the import list. 
// Then we can create a global variable to store our validation expression:
var ValidPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

// The function regexp.MustCompile will parse and compile the regular expression, 
// and return a regexp.Regexp. MustCompile is distinct from Compile 
// in that it will panic if the expression compilation fails, 
// while Compile returns an error as a second parameter.

// This function that uses the validPath expression to validate path and extract the page title:
func GetTitle(w http.ResponseWriter, r *http.Request) (string, error) {
    m := ValidPath.FindStringSubmatch(r.URL.Path)
    if m == nil {
        http.NotFound(w, r)
        return "", errors.New("invalid Page Title")
    }
    return m[2], nil // The title is the second subexpression.
}
// If the title is valid, it will be returned along with a nil error value. 
// If the title is invalid, the function will write a "404 Not Found" error to the HTTP connection, 
// and return an error to the handler. To create a new error, we have to import the errors package.