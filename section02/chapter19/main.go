package main

import (
	"Section02-Chapter19-Testings/utils"
	"log"
)

func main(){
	result, err := utils.Divide(100.0, 0)
	if err != nil{
		log.Println(err)
		return
	}
	log.Println("result of division is:", result)
}