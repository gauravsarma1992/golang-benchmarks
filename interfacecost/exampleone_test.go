package interfacecost

import (
	"testing"
)

func BenchmarkPerson(b *testing.B) {
	testMap := Person{
		Name: "Gary",
		Age:  34,
	}
	for i := 0; i < b.N; i++ {
		IntroducePeople(testMap)
	}
}

func BenchmarkPersonInterface(b *testing.B) {
	testMap := Person{
		Name: "Gary",
		Age:  34,
	}
	for i := 0; i < b.N; i++ {
		IntroduceGreeter(testMap)
	}
}

func BenchmarkPersonEmptyInterface(b *testing.B) {
	testMap := Person{
		Name: "Gary",
		Age:  34,
	}
	for i := 0; i < b.N; i++ {
		IntroduceEveryone(testMap)
	}
}

func BenchmarkAnimal(b *testing.B) {
	testMap := Animal{
		AnimalType: "Tiger",
	}
	for i := 0; i < b.N; i++ {
		IntroduceAnimal(testMap)
	}
}

func BenchmarkAnimalInterface(b *testing.B) {
	testMap := Animal{
		AnimalType: "Tiger",
	}
	for i := 0; i < b.N; i++ {
		IntroduceGreeter(testMap)
	}
}

func BenchmarkAnimalEmptyInterface(b *testing.B) {
	testMap := Animal{
		AnimalType: "Tiger",
	}
	for i := 0; i < b.N; i++ {
		IntroduceEveryone(testMap)
	}
}
