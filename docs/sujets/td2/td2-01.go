package main

import (
	"fmt"
	"sync"
	"time"
)

var d int
var mtx sync.RWMutex

func lecteur() {
	mtx.RLock()
	fmt.Println(d)
	mtx.RUnlock()
}
func redacteur() {
	mtx.Lock()
	d++
	mtx.Unlock()
}

func main() {
	go func() {
		for i := 0; i < 10; i++ {
			lecteur()
			time.Sleep(100 * time.Millisecond)
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			redacteur()
			time.Sleep(150 * time.Millisecond)
		}
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("Fin")
}
