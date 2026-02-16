package utils

import (
	"go-guide/web-app/types"
	"os"
)


// In addition to saving pages, we will want to load pages, too:
// The function loadPage constructs the file name from the title parameter, reads the file's contents into a new variable body, 
// and returns a pointer to a Page literal constructed with the proper title and body values.
// Functions can return multiple values. The standard library function os.ReadFile returns []byte and error. 
// In loadPage, error isn't being handled yet; the "blank identifier" represented by the underscore (_) symbol 
// is used to throw away the error return value (in essence, assigning the value to nothing).
// But what happens if ReadFile encounters an error? For example, the file might not exist. 
// We should not ignore such errors. Let's modify the function to return *Page and error.
func LoadPage(title string) (*types.Page, error) {
    filename := title + ".txt"
    body, err := os.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &types.Page{Title: title, Body: body}, nil
}
// Callers of this function can now check the second parameter; 
// if it is nil then it has successfully loaded a Page. 
// If not, it will be an error that can be handled by the caller.