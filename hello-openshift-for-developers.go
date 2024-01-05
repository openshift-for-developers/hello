package main

import (
	"fmt"
	"net/http"
	"html/template"
	"os"
        "embed"
)

//go:embed static/*
var content embed.FS

func helloHandler(w http.ResponseWriter, r *http.Request) {
	response := os.Getenv("RESPONSE")
	if len(response) == 0 {
		response = "I love Application Platforms!"
	}

	tmpl := `
	<!DOCTYPE html>
	<html>
	<head>
		<style>
			.container {
				display: flex;
				flex-direction: column;
				align-items: center;
				justify-content: center;
				height: 100vh;
			}
		</style>
	</head>
	<body>
		<div class="container">
			<img src="https://raw.githubusercontent.com/andyrepton/hello/main/static/openshift.jpg" alt="OpenShift" style="max-width: 100%; max-height: 50%;">
			<h1>{{.Response}}</h1>
		</div>
	</body>
	</html>
	`

	t, err := template.New("webpage").Parse(tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Response string
	}{
		Response: response,
	}

	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Println("Servicing an impatient beginner's request.")
}

func listenAndServe(port string) {
	fmt.Printf("serving on %s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

func main() {
	http.HandleFunc("/", helloHandler)
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	go listenAndServe(port)

	select {}
}

