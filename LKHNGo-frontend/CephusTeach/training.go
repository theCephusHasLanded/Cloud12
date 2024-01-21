// ----------------------------------------------------------------UDEMY CODING GoesToEleven
package main

import (
	"fmt"
)

func main() {
	//lets make a channrl
	//introducing "1" as a make param to make a buffer channel
	ch := make(chan int, 2)


	//when the shannel is blocked which is clear in the "fatal error: all goroutines are asleep - deadlock!" we will have to pass the value succesffully. we must use the go func() function -- the flow continues.
	go func() {
		ch <- 42
		ch <- 43


	}()

	fmt.Println((<-ch))

}
// ----------------------------------------------------------------UDEMY CODING GoesToEleven
// package main

// import (
// 	"fmt"
// 	"runtime"
// )

// func main() {
// 	fmt.Println("Show me my OS & Architecture:",runtime.GOOS, runtime.GOARCH)
// }


// ----------------------------------------------------------------UDEMY CODING GoesToEleven

// package main

// import (
// 	"fmt"
// 	"sync"
// 	"sync/atomic"
// )

// func main() {
// 	var wg sync.WaitGroup
// 	var incrementor int64
// 	gs := 100
// 	wg.Add(gs)

// 	for i := 0; i < gs; i++ {
// 		go func() {
// 			atomic.AddInt64(&incrementor, 1)
// 			fmt.Println(incrementor)
// 			wg.Done()
// 		}()

// 	}
// 	wg.Wait()
// 	fmt.Println("end value:", incrementor)

// }

// ----------------------------------------------------------------UDEMY CODING GoesToEleven
// package main

// import (
// 	"fmt"
// 	"runtime"
// 	"sync"
// )

// func main() {
// 	var wg sync.WaitGroup

// 	incrementor := 0
// 	gs := 100
// 	wg.Add(gs)
// 	var mu sync.Mutex

// 	for i := 0; i < gs; i++ {
// 		go func() {
// 			mu.Lock()
// 			//reassign the counter to the current counter/incrementor
// 			v := incrementor
// 			runtime.Gosched()
// 			v++
// 			incrementor = v
// 			fmt.Println(incrementor)
// 			mu.Unlock()

// 			wg.Done()

// 		}()

// 	}
// 	wg.Wait()
// 	fmt.Println("end value:",incrementor)

// }

// ----------------------------------------------------------------UDEMY CODING GoesToEleven

// package main

// import (
// 	"fmt"
// )

// type person struct {
// 	first string
// }

// type human interface {
// 	speak()
// }

// func (p *person) speak() {
// 	fmt.Println("Hello")
// }

// func sayHello(h human) {
// 	h.speak()
// }

// func main() {
// 	p1 := person{
// 		first: "Lamer",
// 	}
// 	sayHello(&p1)
// 	p1.speak()
// }

// ----------------------------------------------------------------UDEMY CODING GoesToEleven
// package main

// import (
// 	"fmt"
// 	"runtime"
// 	"sync"
// )

// func main() {

// 	fmt.Println("Begin CPU", runtime.NumCPU())
// 	fmt.Println("Begin Routines", runtime.NumGoroutine())

// 	var waitingGame sync.WaitGroup
// 	waitingGame.Add(2)

// 		go func() {
// 			fmt.Println("Hello From Thing 1")
// 			waitingGame.Done()
// 		}()

// 		go func() {
// 			fmt.Println("Hello From thing 2")
// 			waitingGame.Done()
// 		}()

// 		fmt.Println("Mid CPU", runtime.NumCPU())
// 		fmt.Println("Mid Routines", runtime.NumGoroutine())

// 		waitingGame.Wait()
// 		fmt.Println("exiting the game")

// 	fmt.Println("End CPU", runtime.NumCPU())
// 	fmt.Println("End Routines", runtime.NumGoroutine())

// 	}

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
