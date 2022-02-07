package hot100

import (
	"fmt"
	"testing"
	"time"
)

func Test_permute(t *testing.T) {
	fmt.Println(permute([]int{1, 2, 3}))
}

func Test_asssd(t *testing.T) {
	d := make(chan int, 10)
	go func() {
		d <- 1
		d <- 2
	}()
	go func() {
		for v := range d {
			fmt.Println(v)
		}
	}()
	time.Sleep(time.Second*2)
	d<-33
	time.Sleep(time.Second * 10)
}
