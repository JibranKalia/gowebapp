package main

import "fmt"

type Speaker interface {
	Speak()
}

type Dog struct {
	Name       string
	IsMammal   bool
	PackFactor int
}

func (d Dog) Speak() {
	fmt.Println("Woof!",
		"My name is", d.Name,
		", it is", d.IsMammal,
		"I am a mammal with a pack factor of", d.PackFactor)
}

type Horse struct {
	Name       string
	IsMammal   bool
	BarnFactor int
}

func (h Horse) Speak() {
	fmt. Println("Baah!",
		"My name is", h.Name,
		", it is", h.IsMammal,
		"I am a mamal with Barn factor of", h.BarnFactor)
}

type Cat struct {
	Name        string
	IsMammal    bool
	ClimbFactor int
}


func (c Cat) Speak() {
	fmt.Println("Meow!",
		"My name is", c.Name,
		", it is", c.IsMammal,
		"I am a mammal with a climb factor of", c.ClimbFactor)
}

func main() {

	// Create a list of Animals that know how to speak.
	speakers := []Speaker{

		Dog{
			Name:       "Fido",
			IsMammal:   true,
			PackFactor: 5,
		},

		Cat{
			Name:        "Milo",
			IsMammal:    true,
			ClimbFactor: 4,
		},

		Horse{
			Name:		"Harry",
			IsMammal: 	true,
			BarnFactor: 55,
		},
	}

	// Have the Animals speak.
	for _, spkr := range speakers {
		spkr.Speak()
	}
}
