package main

import (
	"net/http"
	"fmt"
	"sync"
	"time"
)

func GET(url string) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Println(url, resp.Body)
}

func StresTest() {
	go GET("http://127.0.0.1:9000/gateway/health")
	go GET("http://127.0.0.1:9000/gateway/version")
}



type mutexCounter struct {
	mu sync.RWMutex
	x  int64
}

var mc = mutexCounter{
	mu:sync.RWMutex{},
	x: 1,
}

func read() {
	//mc.mu.RLock()
	//mc.mu.Lock()
	mc.x = 2
	fmt.Println(mc.x)

}

func write() {
	mc.mu.Lock()
	defer mc.mu.Unlock()
	mc.x = 3
	time.Sleep(time.Second * 10)
}

func write2() {
	mc.mu.Lock()
	defer mc.mu.Unlock()
	mc.x = 5
	fmt.Println(mc.x)
	//time.Sleep(time.Second * 5)
}


func main() {



	read()

	go write()

	go write2()

	read()


	select {}
}