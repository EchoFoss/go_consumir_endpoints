package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	r := http.DefaultServeMux

	r.HandleFunc("/", fetchCatFacts)
	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		panic(err)
	}
}

func fetchCatFacts(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://cat-fact.herokuapp.com/facts")
	if err != nil {
		panic(err)
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Println("não foi possivel fechar o body")
		}
	}()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error handling response body:", err)
	}

	fmt.Println(string(body))
	_, err = w.Write([]byte(body))
	if err != nil {
		return
	}
}
