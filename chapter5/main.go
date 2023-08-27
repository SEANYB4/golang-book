package main

import "fmt"

func fib(n int) (ans int) {

	if n == 0 {
		return 0
	} else if n == 1{
		return 1 
	} else {
		return n + fib(n-1)
	}
}


func main() {

	fmt.Println(fib(5))

}