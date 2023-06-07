package main

import (
	"fmt"
	"reflect"
)

type Movie struct {
	name string
	date int
}

func main() {
	var x float32 = 3.14159
	describe(x)

	movieTitle := "Blade Runner"
	describe(movieTitle)

	movie := Movie{
		name: movieTitle,
		date: 1982,
	}
	describe(movie)

}

func describe(x interface{}) {
	val := reflect.ValueOf(x)
	typeOf := reflect.TypeOf(x)
	fmt.Println("value:", val)
	fmt.Println("type:", typeOf)
	fmt.Println("kind:", typeOf.Kind())
	if val.Kind() == reflect.Struct {
		fieldCount := val.NumField()
		fmt.Println("fields in struct:", fieldCount)
		for i := 0; i < fieldCount; i++ {
			fmt.Printf("%d. type: %T / value: %#v / kind: %v\n",
				(i + 1), val.Field(i), val.Field(i), val.Field(i).Kind())
		}
	}
	fmt.Println("------------------")
}
