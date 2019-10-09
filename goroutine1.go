//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//func main(){  //one go routine
//	fmt.Println("Hello World")
//	go hello() // 2nd go routine
//	time.Sleep(1 * time.Second)
//}
//
//func hello() {
//	fmt.Println("Hello Goroutine")
//}
//
//
//------------------------------------------------------------------------------------------
package main

//import (
//	"fmt"
//)
//
//func main(){
//	var done = make(chan bool)
//	fmt.Println("done")
//	go hello_goroutine1(done)
//	<- done
//}
//
//func hello_goroutine1(done chan bool){
//	fmt.Println("Channel1")
//	var done1 = make(chan bool)
//	go hello_routine2(done1)
//	<- done1
//	done <- true
//}
//
//func hello_routine2(done1 chan bool){
//	fmt.Println("Channel2")
//	done1 <- true
//}

//===========================================================================

func main(){
	//ch := make(chan int)
	//go func() {
	//	for {
	//		//fmt.Println(<-ch)
	//		_, ok := <-ch
	//		if !ok {
	//			fmt.Println("Recieved")
	//			break
	//		}
	//		fmt.Println("11111")
	//	}
	//}()
	//ch <- 5
	//close(ch)
	//time.Sleep(1 * time.Microsecond)

	// =====================Buffered Channel====================

	//ch := make(chan int,2)  // if we have 3 channel , then it will throw error as it has size of 2 only
	//ch <- 1
	//ch <- 2
	//ch <- 3 // throws an error
	//fmt.Println(<-ch)
	//fmt.Println(<-ch)

	// ==========================================================

}