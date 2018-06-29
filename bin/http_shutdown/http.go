package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func handler(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("1231"))
}

func main() {

	srv := &http.Server{Addr: ":9999", Handler: http.DefaultServeMux}

	http.HandleFunc("/", handler)
	err := srv.ListenAndServe()
	if err != http.ErrServerClosed {
		log.Fatalf("Listen: %s\n", err)
	}
	log.Println("Server gracefully stopped")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		log.Printf("Shuting down server...")
		ctx, err := context.WithTimeout(context.Background(), 100*time.Millisecond)
		if err != nil {
			log.Fatalf("could not create context")
		}
		if err := srv.Shutdown(ctx); err != nil {
			log.Fatalf("could not shutdown, err: %v", err)
		}
	}()

}
