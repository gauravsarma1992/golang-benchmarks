package interfacecost

import (
	"fmt"
)

type (
	Person struct {
		Name string
		Age  int32
	}
	Chatgpt struct {
		Version float32
	}
	Animal struct {
		AnimalType string
	}
	Guest interface {
		SayHello() string
	}
)

func (person Person) SayHello() (result string) {
	result = fmt.Sprintf("Hello to %s, I am %d years old", person.Name, person.Age)
	return
}

func (bot Chatgpt) SayHello() (result string) {
	result = fmt.Sprintf("Hello, I am Chatgpt using GPT model %f", bot.Version)
	return
}

func (animal Animal) SayHello() (result string) {
	result = fmt.Sprintf("Hello, even though humans are also animals, I am still called out separately. I am %s", animal.AnimalType)
	return
}

func IntroduceAnimal(animal Animal) {
	animal.SayHello()
	return
}

func IntroducePeople(person Person) {
	person.SayHello()
	return
}

func IntroduceGreeter(guest Guest) {
	guest.SayHello()
	return
}

func IntroduceEveryone(guest interface{}) {
	switch guest.(type) {
	case Animal:
		guest.(Animal).SayHello()
	case Person:
		guest.(Person).SayHello()
	default:
		fmt.Println("No suitable type found")
	}
	return
}
