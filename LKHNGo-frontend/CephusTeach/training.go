package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("how many CPU's running?:\t", runtime.NumCPU())
	fmt.Println("how many go Routines running?:\t", runtime.NumGoroutine())

	var counter int64

	const gz = 100
	var wait sync.WaitGroup
	wait.Add(gz)

	for i := 0; i < gz; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			//this allows to read and write tot he counter in memory without race conditions
			fmt.Println("counter is safe no race condition:\t", atomic.LoadInt64(&counter))
			runtime.Gosched()
			wait.Done()
		}()
		fmt.Println("how many go Routines running?:\t", runtime.NumGoroutine())
		fmt.Println("count:\t", counter)

	}
	//create the race condition below.
	wait.Wait()
	fmt.Println("how many go Routines running?:\t", runtime.NumGoroutine())
	fmt.Println("count:\t", counter)

}

// ----------------------------------------------------------------PAK CODING 1:1
// package main

// import (
// 	"fmt"
// 	"runtime"
// 	"sync"
// )

// var wg sync.WaitGroup

// // func main() {
// 	var name string = "Cephus"
// 	fmt.Printf("Hello World I'm, %s\t\n", name)

// 	fmt.Println("OS\t", runtime.GOOS)
// 	fmt.Println("Architecture\t", runtime.GOARCH)
// 	fmt.Println("CPU's running\t", runtime.NumCPU())
// 	fmt.Println("Live  Go Routines", runtime.NumGoroutine())

// 	energy := []int{45, 42, 41, 40}
// 	for _, v := range energy {
// 		fmt.Printf("checking Energy Levels -- @%v\n", v)
// 		fmt.Printf("the Type is -- %T\n", v)

// 	}
// 	isBool, energyLevel := activateGate(39)
// 	fmt.Println(isBool, energyLevel)

// 	wg.Add(1)
// 	go activateGate(40)
// 	go activateGate(41)
// 	go activateGate(42)
// 	go activateGate(43)
// 	wg.Wait()

// }

// var energyLevel int = 5
// var energyPointer *int = &energyLevel

// var quantumCodes [5]int = [5]int{11, 33, 44, 5, 66}

// func activateGate(energyLevel int) (bool, int) {
// 	if energyLevel > 40 {
// 		fmt.Println("Energylevels are a go")
// 		return true, energyLevel
// 	} else {
// 		fmt.Println("Energylevels are not a go!")
// 		return false, energyLevel
// 	}
// // wg.Wait()
// // return false, energyLevel
// }
