package main

//./bench  65.81s user 1.71s system 265% cpu 25.459 total
import (
	"log"
	"time"
)

func main() {
	c := time.Now()
	runTreeBench()
	log.Println(time.Now().UnixMilli()-c.UnixMilli(), "ms")
}
