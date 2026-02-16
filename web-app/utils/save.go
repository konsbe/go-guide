package utils

import (
	"go-guide/web-app/types"
	"os"
)


// But what about persistent storage? We can address that by creating a save function for Page:
// It constructs a filename by appending ".txt" to the page title, 
// and then writes the page body to that file with permissions set to 0600 (read and write permissions for the owner only). 
// The function returns an error if there is a problem writing the file.
func SavePage(p *types.Page) error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}



// In Go, you can only define methods on types that are declared in the same package. 
// Since Page is declared in go-guide/web-app/types, 
// you can only add methods to it within the types package.
// If you want Save as a method, move it to types/types.go:
// package types

// import "os"

// type Page struct {
//     Title string
//     Body  []byte
// }

// func (p *Page) Save() error {
//     filename := p.Title + ".txt"
//     return os.WriteFile(filename, p.Body, 0600)
// }