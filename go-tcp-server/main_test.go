package main

import (
	"fmt"
	"net/http"
	"sync"
	"testing"
)

func TestTcpServer(t *testing.T) {
	total := 10
	var wg sync.WaitGroup

	for i := 0; i < total; i++ {
		wg.Add(1)
		//go net.Dial("localhost", ":8090")
		go get()
		defer wg.Done()
		//go http.Get("http://localhost:8090")
	}
	wg.Wait()
	return
}

func get() {
	resp, err := http.Get("http://localhost:8090/")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	fmt.Println(resp)
}
