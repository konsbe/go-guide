package main

import (
    "fmt"
	"go-guide/web-app/types"
	"go-guide/web-app/utils"
)

func wiki() {
    p1 := &types.Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
    utils.SavePage(p1)
    p2, _ := utils.LoadPage("TestPage")
    fmt.Println(string(p2.Body))
}