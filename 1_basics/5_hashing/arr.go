package main

import (
	"fmt"
	"runtime"
)

func main() {
	var m1, m2 runtime.MemStats
	runtime.ReadMemStats(&m1)

	a := 5
	fmt.Println(a)

	runtime.ReadMemStats(&m2)
	fmt.Printf("Heap alloc diff: %d bytes\n", m2.HeapAlloc-m1.HeapAlloc)
}
