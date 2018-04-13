package main

import (
	"fmt"
	"sync"
	"time"
	"github.com/k0kubun/pp"
)

var unresolved []int64
var l sync.Mutex

func main() {

	parallelism := 5
	processDuration := 5 * time.Second
	requestEvery := 1 * time.Second

	rqs := make(chan request, parallelism)
	work := make(chan request, parallelism)
	rts := make(chan int64, parallelism)
	orts := make(chan int64)

	go makeRequests(rqs, requestEvery)
	go readRequests(rqs, work)
	go orderResults(rts, orts)

	for i := 0; i < parallelism; i++ {
		go processRequests(work, rts, processDuration)
	}

	for r := range orts {
		pp.Printf("orts size: %v\n", len(orts) )
		fmt.Println(r, "-")
	}
}

func makeRequests(rqs chan request, requestEvery time.Duration) {
	order := int64(0)
	for {
		pp.Println("in makeRequests")
		order++
		rqs <- request{order}
		time.Sleep(requestEvery)
	}
}

func readRequests(rqs chan request, work chan request) {
	for r := range rqs {
		pp.Println("in readRequests")
		l.Lock()
		unresolved = append(unresolved, r.order)
		l.Unlock()
		work <- r
	}
}

func processRequests(work chan request, rts chan int64, processDuration time.Duration) {
	pp.Println("in processRequests")
	for r := range work {
		rts <- r.process(processDuration)
	}
}

func orderResults(rts chan int64, orts chan int64) {
	pp.Println("in orderResults")
	rtBuf := make(map[int64]int64)
	for rt := range rts {
		rtBuf[rt] = rt
	loop:
		if len(unresolved) > 0 {
			u := unresolved[0]
			if rtBuf[u] != 0 {
				l.Lock()
				unresolved = unresolved[1:]
				l.Unlock()
				orts <- rtBuf[u]
				delete(rtBuf, u)
				goto loop
			}
		}
	}
}

type request struct {
	order int64
}

func (r request) process(processDuration time.Duration) int64 {
	time.Sleep(processDuration)
	return r.order
}