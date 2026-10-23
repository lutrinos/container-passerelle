package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Redirection struct {
	Id  string
	Url string
}

var redirections []Redirection
var client = &http.Client{Timeout: 5 * time.Second}

func read() {
	body, read_err := os.ReadFile("/config.yaml")

	if read_err != nil {
		log.Fatal("unable to read config file", read_err)
	}

	err := yaml.Unmarshal([]byte(body), &redirections)

	if err != nil {
		log.Fatal("parsing error")
	}

	log.Println(len(redirections), "redirections will be applied :")

	for i := 0; i < len(redirections); i++ {
		log.Println("\t-", redirections[i].Id, ":", redirections[i].Url)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {

	for i := 0; i < len(redirections); i++ {
		if redirections[i].Id == r.URL.Path[1:] {
			res, err := client.Get(redirections[i].Url)

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprint(w, "Not OK")
				return
			}

			defer res.Body.Close()
			io.Copy(io.Discard, res.Body)

			w.WriteHeader(res.StatusCode)
			fmt.Fprintf(w, "OK")
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "Not OK")
}

func main() {
	read()

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
