package main

import "fmt"

type Animal struct {
    Name string
    mean bool
}

type Cat struct {
    Basics Animal
    MeowStrength int
}

type Dog struct {
    Animal
    BarkStrength int
}

type AnimalSounder interface {
	MakeNoise()
}

func (dog *Dog) MakeNoise() {
	barkStrength := dog.BarkStrength

	if dog.mean == true {
		barkStrength = barkStrength * 5
	}

	for i := 0; i < barkStrength; i++{
		fmt.Printf("Bark ")
	}
}

func (cat *Cat) MakeNoise() {
	meowStrength := cat.MeowStrength

	if cat.Basics.mean == true {
		meowStrength = meowStrength * 5
	}

	for i := 0; i < meowStrength; i++{
		fmt.Printf("Meow ")
	}
}



func main() {
	myDog := &Dog{
		Animal{
			"Rover",
			false,
		},
		2,
	}

	myCat := &Cat{
		Basics: Animal{
			Name: "Julius",
			mean: true,
		},
		MeowStrength: 3,
	}

	MakeSomeNoise(myDog)
    MakeSomeNoise(myCat)
}


func MakeSomeNoise(animalSounder AnimalSounder){
	animalSounder.MakeNoise()
}
