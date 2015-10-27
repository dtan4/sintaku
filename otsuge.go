package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/codegangsta/negroni"
)

const (
	TerminationTimeURL = "http://169.254.169.254/latest/meta-data/spot/termination-time"
)

func checkTerminationTime() (int, string) {
	var body string

	resp, _ := http.Get(TerminationTimeURL)
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		byteArray, _ := ioutil.ReadAll(resp.Body)
		body = string(byteArray)
	} else {
		body = ""
	}

	return resp.StatusCode, body
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		statusCode, body := checkTerminationTime()
		w.WriteHeader(statusCode)
		fmt.Fprintln(w, body)
	})

	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(":8080")
}
