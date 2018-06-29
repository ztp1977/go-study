package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	results := Google("golang")
	elapsed := time.Since(start)

	fmt.Println(results)
	fmt.Println(elapsed)
}

type Result struct{}

func Google(s string) (results []Result) {

	results = append(results, Web(s))
	results = append(results, Image(s))
	results = append(results, Video(s))

	return
}

func Google2(q string) (results []Result) {
	c := make(chan Result)
	go func() { c <- Web(q) }()
	go func() { c <- Image(q) }()
	go func() { c <- Video(q) }()

	for i := 0; i < 3; i++ {
		result := <-c
		results = append(results, result)
	}
	return
}

func Google3(q string) (results []Result) {
	c := make(chan Result)
	go func() { c <- Web(q) }()
	go func() { c <- Image(q) }()
	go func() { c <- Video(q) }()

	timeout := time.After(80 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("timed out")
			return
		}
	}
	return
}

//func First(q string, replicas ...Search) Result {
//	c := make(chan Result)
//	searchReplics := func(i int) { c <- replicas[i](q) }
//	for i := range replicas {
//
//	}
//}

func Video(s string) Result {
	return Result{}
}
func Image(s string) Result {
	return Result{}
}
func Web(s string) Result {
	return Result{}
}
