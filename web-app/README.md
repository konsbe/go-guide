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

### Introducing the net/http package (an interlude)¶
