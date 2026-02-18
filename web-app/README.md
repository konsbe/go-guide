## [Writing Web Applications](https://go.dev/doc/articles/wiki/)

### Introduction

Covered in this tutorial:
- Creating a data structure with load and save methods
- Using the net/http package to build web applications
- Using the html/template package to process HTML templates
- Using the regexp package to validate user input
- Using closures
- Assumed knowledge:

Programming experience
- Understanding of basic web technologies (HTTP, HTML)
- Some UNIX/DOS command-line knowledge

### Compile and executing

After compiling and executing this code, a file named TestPage.txt would be created, containing the contents of p1. The file would then be read into the struct p2, and its Body element printed to the screen.

You can compile and run the program like this:

    go build wiki.go
    ./wiki

### Introducing the net/http package (an interlude)

Here's a full working example of a simple web server:

    //go:build ignore

    package main

    import (
        "fmt"
        "log"
        "net/http"
    )

    func handler(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
    }

    func main() {
        http.HandleFunc("/", handler)
        log.Fatal(http.ListenAndServe(":8080", nil))
    }

The main function begins with a call to http.HandleFunc, which tells the http package to handle all requests to the web root ("/") with handler.

It then calls http.ListenAndServe, specifying that it should listen on port 8080 on any interface (":8080"). (Don't worry about its second parameter, nil, for now.) This function will block until the program is terminated.

ListenAndServe always returns an error, since it only returns when an unexpected error occurs. In order to log that error we wrap the function call with log.Fatal.

The function handler is of the type http.HandlerFunc. It takes an http.ResponseWriter and an http.Request as its arguments.

An http.ResponseWriter value assembles the HTTP server's response; by writing to it, we send data to the HTTP client.

An http.Request is a data structure that represents the client HTTP request. r.URL.Path is the path component of the request URL. The trailing [1:] means "create a sub-slice of Path from the 1st character to the end." This drops the leading "/" from the path name.

If you run this program and access the URL:

http://localhost:8080/monkeys

### Future tasks

Visiting http://localhost:8080/view/ANewPage should present you with the page edit form. You should then be able to enter some text, click 'Save', and be redirected to the newly created page.

Other tasks¶
Here are some simple tasks you might want to tackle on your own:

Store templates in tmpl/ and page data in data/.
Add a handler to make the web root redirect to /view/FrontPage.
Spruce up the page templates by making them valid HTML and adding some CSS rules.
Implement inter-page linking by converting instances of [PageName] to
<a href="/view/PageName">PageName</a>. (hint: you could use regexp.ReplaceAllFunc to do this)
