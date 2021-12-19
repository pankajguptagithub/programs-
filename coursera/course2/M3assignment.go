package main

import (
	"fmt"
	"strings"
)
type Animal struct{
	food string
	locomotion string
	noise string	
}
func (a Animal) Eat(name string){
	switch name {
    case "Cow":
        fmt.Println("grass")
    case "Bird":
        fmt.Println("worms")
    case "Snake":
        fmt.Println("mice")
	default:
		fmt.Println("Incorrect animal name")
   	}
}
func (a Animal) Move(name string){
	switch name {
    case "Cow":
        fmt.Println("walk")
    case "Bird":
        fmt.Println("fly")
    case "Snake":
        fmt.Println("slither")
	default:
		fmt.Println("Incorrect animal name")
    }
}
func (a Animal) Speak(name string){
	switch name {
    case "Cow":
        fmt.Println("moo")
    case "Bird":
        fmt.Println("peep")
    case "Snake":
        fmt.Println("hsss")
	default:
		fmt.Println("Incorrect animal name")
    }
}
func main(){
	var inp string
	for{
		fmt.Println("> Enter the animal name (cow/bird/snake) followed by information needed(eat/move/speak) Eg Cow-Move(initial letter Caps)")
		fmt.Scanf("%s",&inp)
		animal_name := strings.Split(inp,"-")[0]
		information := strings.Split(inp,"-")[1]
		var a Animal
		switch information {
		case "Eat":
			 a.Eat(animal_name)
		case "Move":
			 a.Move(animal_name)
		case "Speak":
			a.Speak(animal_name)
		default:
			fmt.Println("Incorrect input for animal information")
		}
	}
}