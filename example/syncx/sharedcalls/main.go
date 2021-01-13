package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/zjz894251se/go-zero/core/stringx"
	"github.com/zjz894251se/go-zero/core/syncx"
)

func main() {
	const round = 5
	var wg sync.WaitGroup
	barrier := syncx.NewSharedCalls()

	wg.Add(round)
	for i := 0; i < round; i++ {
		go func() {
			defer wg.Done()
			val, err := barrier.Do("once", func() (interface{}, error) {
				time.Sleep(time.Second)
				return stringx.RandId(), nil
			})
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(val)
			}
		}()
	}

	wg.Wait()
}
