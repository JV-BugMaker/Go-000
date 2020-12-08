package main

import(
	"fmt"
	"sync"
//	"time"
)
var wg sync.WaitGroup
var counter int = 0


func main(){
	for routine:=1;routine<=2;routine++{
		wg.Add(1)
		go Routine(routine)
	}
	wg.Wait()
	fmt.Print(counter)
}

func Routine(id int){
	for count := 0; count < 2; count++{
		value := counter
//		time.Sleep(time.Nanosecond*1)
		value++
		counter = value
	}
	wg.Done()
}
