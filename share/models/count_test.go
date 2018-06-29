package models

import "testing"

func BenchmarkCount(b *testing.B) {
	for n := 0; n < b.N; n++ {
		count(n)
	}
}
