// task-1 FizzBuzz

package main

import(
	"fmt"
)

func main(){
	for i:=1; i<=100; i++ {
		FizzBuzz(i)
	}
}

func FizzBuzz(n int){
	if n%3==0 && n%5==0{
	fmt.Println("FizzBuzz")
	} else if n%3==0{
	fmt.Println("Fizz")
	} else if n%5==0{
	fmt.Println("Buzz")
	} else {fmt.Println(n)}
}

//task-2 Find weekday from given date
// package main
// import(
// 	"fmt"
// 	"time"
// )

// func main(){
// 	dobStr := "06.09.2022"
// 	givenDate, err := time.Parse("02.01.2006", dobStr)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("%s is %s\n", givenDate.Format("02-01-2006"), FindWeekday(givenDate))

// }

// func FindWeekday(date time.Time) (weekday string) {
// 	weekday = date.Weekday().String()
// 	return
// }

// task-3
// package main

// import "fmt"

// func main() {
// 	DisplayNumberInReverseOrderWithDefer()
// }

// func DisplayNumberInReverseOrderWithDefer() {
// 	for i := 0; i < 100; i++ {
// 		defer fmt.Println(i)
// 	}
// }
