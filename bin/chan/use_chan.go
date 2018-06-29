package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan string)
	go boring("boring!", c)
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c)
	}

	c2 := boring2("boring2!")
	for i := 0; i < 5; i++ {
		fmt.Printf("You Say: %q\n", <-c2)
	}
	fmt.Println("You're boring: I'm Leaving")

	joe := boring2("Joe")
	ann := boring2("Ann")
	for i := 0; i < 5; i++ {
		fmt.Println(<-joe)
		fmt.Println(<-ann)
	}

	// mulitplexing
	c2 = mergeChan(boring2("Joe"), boring2("Ann"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c2)
	}
	fmt.Println("You're both boring2; I'm leaving")
	c2 = mergeChan2(boring2("Joe"), boring2("Ann"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c2)
	}
	fmt.Println("You're both boring3; I'm leaving")

	// timeout
	c4 := boring2("Joe")
	for {
		select {
		case s := <-c4:
			fmt.Println(s)
		case <-time.After(1 * time.Second):
			fmt.Println("you're too slow")
			return
		}
	}

	//c3 := make(chan Message)
	//for i := 0; i < 5; i++ {
	//	msg1 := <-c3
	//	fmt.Println(msg1.str)
	//	msg2 := <-c3
	//	fmt.Println(msg2.str)
	//	msg1.wait <- true
	//	msg2.wait <- true
	//}

	//waitForIt := make(chan bool)
	//for i := 0; i < 5; i++ {
	//	c3 <- Message{fmt.Sprintf("%s: %d", "hello", i), waitForIt}
	//	time.Sleep(time.Duration(rand.Intn(2e3)) * time.Millisecond)
	//	<-waitForIt
	//}
}

func boring(s string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s, %d", s, i*10)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func boring2(s string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s, %d", s, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func mergeChan(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}

func mergeChan2(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()
	return c
}

type Message struct {
	str  string
	wait chan bool
}
