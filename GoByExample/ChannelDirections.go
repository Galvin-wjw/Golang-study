package main

import "fmt"

func ping (pings chan <- string,msg string)  {
	pings <- msg
}


//one channel for receives (pings) and a second for sends (pongs)
func pong(pings <- chan string ,pongs chan <- string)  {
	msg := <- pings
	pongs <- msg
}

func main() {
	pings := make(chan string,1)
	pongs := make(chan string,1)
	ping(pings,"passed message")
	pong(pings,pongs)
	fmt.Println(<-pongs)
}