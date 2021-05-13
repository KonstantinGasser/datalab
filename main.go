package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"

	"github.com/KonstantinGasser/datalab/libraries/pool"
)

func main() {

	p := pool.New(2)
	cancel := p.Start()
	for i := 0; i < 2; i++ {
		p.Jobs <- func(i int) (interface{}, error) {
			if rand.Int()%2 == 0 {
				fmt.Println("Job will sleep 500mil: ", i)
				time.Sleep(5000 * time.Millisecond)
				fmt.Println("Done sleeping: ", i)
			}
			return "hello", nil
		}
	}
	ticker := time.NewTicker(200 * time.Millisecond)
	for {
		select {
		case res := <-p.Results:
			fmt.Println(res)
		case <-ticker.C:
			fmt.Println("Canceling Context")
			cancel()

			time.Sleep(2 * time.Second)
			fmt.Println(runtime.NumGoroutine())
			return
		}
	}

}
