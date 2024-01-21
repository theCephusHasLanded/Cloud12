package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	fmt.Println("how many CPU's running?:\t", runtime.NumCPU())
	fmt.Println("how many go Routines running?:\t", runtime.NumGoroutine())

	counter := 0

	const gz = 10
	var wait sync.WaitGroup
	wait.Add(gz)

	for i := 0; i < gz; i++ {
		go func() {
			v := counter
			//if you want to wait or have a timeout -- import time and give it some rest
			time.Sleep(time.Second)
			runtime.Gosched()
			v++
			counter = v
			//now set your wait group to done
			wait.Done()
		}()
		fmt.Println("how many go Routines running?:\t", runtime.NumGoroutine())

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
