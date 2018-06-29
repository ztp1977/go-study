package funs

import "github.com/k0kubun/pp"

func DoChan() {

	ch := make(chan int)
	chBuff := make(chan int, 10)

	// sent message
	ch <- 1
	// receive message
	if v, ok := <-ch; ok {
		pp.Println(v)
	} else {
		close(ch)
	}

	select {
	case v := <-chBuff:
		pp.Println(v)
	default:
		pp.Println("could not receive unknown channel")
	}

	close(chBuff)

}
