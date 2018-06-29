package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/k0kubun/pp"
)

func main() {
	var m sync.Map

	// 1.0 ここで三つの要素をmapに入れる。順番ランダム
	for i := 0; i < 3; i++ {
		// ここは3つのgoroutineで無限ループしてる。
		// どのような時使うのか？
		go func(i int) {
			for j := 0; ; j++ {
				m.Store(i, j) // ここは同期
			}
			// こちらは表示しないわけ
			pp.Println("store", m)
		}(i)
	}

	// 2. 10回その時の値を表示
	for i := 0; i < 10; i++ {
		// 2.1 それぞれのキーに適用するロジック
		m.Range(func(key, v interface{}) bool {
			fmt.Printf("%d: %d\t", key, v)
			return true
		})
		fmt.Println()
		time.Sleep(time.Second)
	}
}
