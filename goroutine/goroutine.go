package goroutine

import (
	. "fmt"
	"runtime"
)

func GoroutineImpl() {
	go say("world")
	say("hello")
}

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		Println(s)
	}
}