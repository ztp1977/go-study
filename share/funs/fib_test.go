package funs

import "testing"

func TestFib(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"0", args{0}, 0},
		{"1", args{1}, 1},
		{"2", args{2}, 1},
		{"3", args{3}, 2},
		{"4", args{4}, 3},
		{"5", args{5}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fib(tt.args.i); got != tt.want {
				t.Errorf("Fib() = %v, want %v", got, tt.want)
			}
		})
	}
}

var result int

func BenchmarkFib(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		r = Fib(10)
	}
	result = r
}


func benchmarkFib(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib(i)
	}
}

func BenchmarkFib10(b *testing.B) {
	benchmarkFib(10, b)
}
func BenchmarkFib20(b *testing.B) {
	benchmarkFib(20, b)
}
