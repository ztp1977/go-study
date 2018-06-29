package main

import (
	"time"
	"net/http"
	"sync"
	"io/ioutil"
	"fmt"
)

// Requestの構造体
type request struct{
	r *http.Request
	response chan []byte
}

func main() {
	// サイズは100のchan
	requests := make(chan request, 100)

	// サーバーを起動する
	go startServer(requests)
	// processを起動する
	go process(requests)

	// 6msの間隔で500個のmessageを作る
	makeRequests(500, 6*time.Millisecond)
}

func makeRequests(count int, cooldown time.Duration) {
	wg := sync.WaitGroup{}
	for i := 0; i < count; i++ {
		go func() {
			wg.Add(1)
			defer wg.Done()
			response, _ := http.Get("http://localhost:9000/requests")
			defer response.Body.Close()
			b, _ := ioutil.ReadAll(response.Body)
			fmt.Printf(string(b))
		}()
		time.Sleep(cooldown)
	}
	wg.Wait()
}

func process(rq chan request) {
	for r := range rq {
		r.process()
	}
}

func startServer(rq chan request) {
	http.HandleFunc("/requests", handle(rq))
	http.ListenAndServe(":9000", nil)
}

func handle(rq chan request) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if len(rq) < cap(rq) {
			r := newRequest(r)
			rq <- r
			w.Write(<-r.response)
		} else {
			w.Write([]byte("X"))
		}
	}
}

func newRequest(r *http.Request) request {
	return request{r, make(chan []byte)}
}

func (r request) process() {
	time.Sleep(10 * time.Millisecond)
	r.response <- []byte("√")
}


