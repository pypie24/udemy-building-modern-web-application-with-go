package main

import (
	"encoding/json"
	"log"
)

type Person struct{
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog bool `json:"has_dog"`
}

func main(){
	myJson := `[
		{
			"first_name": "Clark",
			"last_name": "Kent",
			"hair_color": "black",
			"has_dog": true
		},
		{
			"first_name": "Bruce",
			"last_name": "Wayne",
			"hair_color": "black",
			"has_dog": false
		}
	]`
	var unmarshalled []Person

	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil{
		log.Println("Error unmarshalling json", err)
	}
	log.Printf("unmarshalled: %v", unmarshalled)

	m1 := Person{"Henry", "Hoang", "Black", false}
	m2 := Person{"Airi", "Le", "Black", false}
	mySlice := []Person{m1, m2}
	
	newJson, err := json.MarshalIndent(mySlice, "", "	")
	if err != nil{
		log.Println("Error marshalling json", err)
	}

	log.Println(string(newJson))
	log.Println(myJson)

}
