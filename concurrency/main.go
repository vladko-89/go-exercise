package main

import (
	"fmt"

	"math/rand"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	res := []int{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go getRandNum(ch)
	}
	go func() {
		wg.Wait()
		for _, v := range res {
			fmt.Print(``, v, ' ')
		}
		close(ch)

	}()

	for num := range ch {
		go func() {
			res = append(res, getScuadNum(num))

			wg.Done()
		}()
	}

}

func getRandNum(ch chan int) {
	ch <- rand.Intn(100)
}

func getScuadNum(num int) int {
	return num * num
}
