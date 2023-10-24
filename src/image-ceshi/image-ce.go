package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func check(err error) {
	if err != nil {
		fmt.Printf("报错：%s \n", err.Error())
	}
}

func Image(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("view.html")
	check(err)
	writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = html.Execute(writer, nil)

	check(err)
}

func main() {
	http.HandleFunc("/image", Image)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("image"))))
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("获取当前工作目录失败：", err)
	} else {
		fmt.Println("当前工作目录：", dir)
	}
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
