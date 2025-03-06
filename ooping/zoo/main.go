package main

import "fmt"

func main() {
	var myZoo []Animal
	leo := Lion{Name: "Leo", location: LatLong{Lat: 10.40, Long: 11.5}}
	peggy := Pegion{Name: "Peggy", location: LatLong{Lat: 10.40, Long: 11.5}}
	myZoo = append(myZoo, &leo)
	myZoo = append(myZoo, &peggy)

	for _, animal := range myZoo {
		fmt.Println("Name:", animal.GetName(),"Says:", animal.Speak())		
	}

}