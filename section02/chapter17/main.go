package main

import (
	"log"
	"Section02-Chapter17-Channels/helpers"
)

const numberPool = 10

func CalculateValue(intChan chan int){
	randomNumber := helpers.RandomNumber(numberPool)
	intChan <- randomNumber
}


func main(){
	intChan := make(chan int)

	defer close(intChan)
	
	go CalculateValue(intChan)

	num := <- intChan

	log.Println(num)
}