package main

import (
	"Section02-Chapter16-Packages/helpers"
	"log"
)


func main(){
	log.Println("Hello")
	var someType helpers.SomeType
	someType.TypeName = "Some name"
	log.Println(someType.TypeName)
}