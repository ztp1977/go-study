package main

import (
	"fmt"
	"log"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if p, ok := w.(http.Pusher); ok {
		err := p.Push("/style.css", nil)
		if err != nil {
			log.Printf("could not push %v", err)
		}
	}
	fmt.Fprintln(w, "123")
}
