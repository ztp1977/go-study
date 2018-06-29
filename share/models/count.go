package models

import "sync"

func count(n int) {
	var wg sync.WaitGroup
	wg.Add(n)
	m := map[int]int{}
	var mu sync.Mutex
	for i := 1; i <= n; i++ {
		go func(i int) {
			for j := 0; j < i; j++ {
				mu.Lock()
				m[i]++
				mu.Unlock()
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}
