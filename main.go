package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/fatih/color"
)

func getDir() string {
	path, _ := os.Getwd()
	return path
}

func welcome() {
	welcome := `  __ _   ___          ___   ___  _ __ __   __  ___  _ __ 
 / _' | / _ \  _____ / __| / _ \| '__|\ \ / / / _ \| '__|
| (_| || (_) ||_____|\__ \|  __/| |    \ V / |  __/| |   
 \__, | \___/        |___/ \___||_|     \_/   \___||_|   
 |___/ `
	fmt.Println(color.GreenString(welcome))
}

func getPort() string {
	ct := os.Args
	var target string
	if len(ct) == 1 {
		target = "8000"
	} else if len(ct) != 4 {
		fmt.Println("ERROR: BAD PORT NUMBER")
		os.Exit(1)
	} else {
		target = ct[1]
	}
	port := ":" + target
	return port
}

func handler(w http.ResponseWriter, r *http.Request) {
	http.FileServer(http.Dir(getDir())).ServeHTTP(w, r)
	request := color.HiWhiteString("Request: ")
	var method string
	switch r.Method {
		case "GET":
		method = color.HiYellowString("GET")
		case "POST":
		method = color.HiGreenString("POST")
		case "PUT":
		method = color.BlueString("PUT")
		case "DELETE":
		method = color.RedString("DELETE")
		default :
		method = color.RedString("UNKNOWN")
		}
	log.Printf(request + method +" "+ r.URL.Path)
}

func main() {
	welcome()
	port := getPort()

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	start := color.HiYellowString("[START]")
	link := color.HiCyanString("http://localhost:" + port)
	log.Print(start + " Server is running on " + link + " ...")
	log.Fatal(http.ListenAndServe(port, mux))
}

