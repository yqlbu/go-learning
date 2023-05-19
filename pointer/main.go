package main

import "fmt"

func main() {
	wallace := "Uses a ninja star"
	fmt.Println(wallace) // Prints "Uses a ninja star"
	changeGear(wallace)
	fmt.Println(wallace) // Prints "Uses a ninja star"
	changeGearPointer(&wallace)
	fmt.Println(wallace) // Prints "Uses a ninja sword"
}

func changeGear(wallaceCopy string) {
	wallaceCopy = "Uses a ninja sword"
}

func changeGearPointer(wallaceCopy *string) {
	*wallaceCopy = "Uses a ninja sword"
}
