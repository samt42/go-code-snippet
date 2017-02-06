package main

//import "sync"

func main() {
	const workers = 3
	m := map[int]int{}
	//var mutex sync.RWMutex
	for i := 1; i <= workers; i++ {
		go func(i int) {
			for j := 0; j < i; j++ {
				//mutex.Lock()
				m[i]++
				//mutex.Unlock()
			}
		}(i)
	}
}
