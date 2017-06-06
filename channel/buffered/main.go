package main

var a string
var c = make(chan int)

func f() {
	a = "hello, world\n"
	<-c
}

func main() {
	go f()
	c <- 0
	print(a)
}
