// ----------------------------------------------------------------UDEMY CODING GoesToEleven
package main

import (
	"fmt"
)

func main(){
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	//send
	go send(eve, odd, quit)

	//recieve
	recieve(eve, odd, quit)
}
func recieve(e,o,q <-chan int){
	for {
		select {
		case v := <- e:
			fmt.Println("recieved from the even chan:", v)
		case v := <- o:
			fmt.Println("recieved from the odd chan:", v)
		case v := <- q:
			fmt.Println("recieved from the quit chan:", v)
		// close(q)
		return
	}
	}
}

func send(e,o,q chan <- int){
	for i := 0; i <100; i++ {
		if i % 2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(e)
	close(o)
	q <- 0
}


// ----------------------------------------------------------------UDEMY CODING GoesToEleven

// package main

// import(
// 	"fmt"
// ) 

// func main() {
// 	ch := make(chan int)

// 	//send
// 	go func(){
// 		for i := 0; i < 77; i++ {
// 			ch <- i
// 		}
// 		close(ch)
// 	}()

// 	//recieve 
// 	for value := range ch {
// 		fmt.Println(value)
// 	} 
// 	fmt.Println("We leaving!")
// 	fmt.Printf("ch\t%T\n", ch)


// }

// ----------------------------------------------------------------UDEMY CODING GoesToEleven
/*
package main

import "fmt"

func main() {

	c := make(chan int)

	//send
	go dragon(c)

	//recieve
	horse(c)

	fmt.Println("We about to end it here...")
}

//send
func dragon(c chan<- int){
	c <- 42
}

//recieve
func horse(c <-chan int){
fmt.Println(<- c)
}
// ----------------------------------------------------------------UDEMY CODING GoesToEleven

/*
package main

import (
	"fmt"
)

func main() {
	//lets make a directional channel -- one you can send only -- read from left to right
	ch := make(chan int) //recieve
	cr := make(<-chan int) //recieve
	cs := make(chan<- int) //send

	fmt.Println("*****************")
	fmt.Printf("ch\t%T\n", ch)
	fmt.Printf("cr\t%T\n", cr)
	fmt.Printf("cs\t%T\n", cs)

	//we can take a bi-driectional channel and assign it to a recieve  -- going from general to specific works but not so much the other way.

	//when the shannel is blocked which is clear in the "fatal error: all goroutines are asleep - deadlock!" we will have to pass the value succesffully. we must use the go func() function -- the flow continues.
		ch <- 42
		ch <- 43



		// fmt.Println((<-ch))
		// fmt.Println((<-ch))
		// fmt.Println("*****************")
}		// fmt.Printf("%T\n", ch)
*/



// ----------------------------------------------------------------UDEMY CODING GoesToEleven
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	//lets make a channrl
// 	//introducing "1" as a make param to make a buffer channel
// 	ch := make(chan int, 2)


// 	//when the shannel is blocked which is clear in the "fatal error: all goroutines are asleep - deadlock!" we will have to pass the value succesffully. we must use the go func() function -- the flow continues.
// 	go func() {
// 		ch <- 42
// 		ch <- 43


// 	}()

// 	fmt.Println((<-ch))

// }
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
