package funs

func Fib (i int) int {
	if i < 2 {
		return i
	}
	return Fib(i -1) + Fib(i - 2)
}
