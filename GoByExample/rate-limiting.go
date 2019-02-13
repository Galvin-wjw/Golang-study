package main

import (
	"time"
	"fmt"
)

func main() {
	requests := make(chan int ,5)
	for i := 1; i <= 5; i++{
		requests <- i
	}
	close(requests)


	// This limiter channel will receive a value every 2000 milliseconds
	limiter := time.Tick(time.Millisecond * 2000)
	for req := range requests{
		<- limiter
		fmt.Println("requests",req, time.Now())
	}

	//This burstyLimiter channel will allow bursts of up to 3 events
	burstyLimiter := make(chan time.Time,3)


	//Fill up the channel to represent allowed bursting.
	for i := 0;i < 3;i ++{
		burstyLimiter <- time.Now()
	}


	//Every 2000 milliseconds we’ll try to add a new value to burstyLimiter, up to its limit of 3
	go func() {
		for t := range time.Tick(time.Millisecond * 2000){
			burstyLimiter <- t
		}
	}()

	//simulate 5 more incoming requests
	burstyRequests := make(chan int,5)
	for i := 1 ; i <= 5 ; i++  {
		burstyRequests <- i
	}
	close(burstyRequests)

	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request",req,time.Now())
	}


}
