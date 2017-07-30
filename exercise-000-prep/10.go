package main

import (
	"fmt"
	"math/rand"
	"time"
)

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func sleep(n int) {
	for {
		select {
		case <-time.After(time.Duration(n) * time.Second):
			return
		}
	}
}

func main() {
	// var c chan string = make(chan string)

	// go pinger(c)
	// go ponger(c)
	// go printer(c)

	// var input string
	// fmt.Scanln(&input)

	fmt.Println("first")
	sleep(5)
	fmt.Println("after wait")
}
