package types

//  A wiki consists of a series of interconnected pages, 
// each of which has a title and a body (the page content).
// The type []byte means "a byte slice". 
// The Body element is a []byte rather than string because that is the type 
// expected by the io libraries we will use, as you'll see below.
// The Page struct describes how page data will be stored in memory. 
type Page struct {
    Title string
    Body  []byte
}