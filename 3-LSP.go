// Liskov Substitution Principle (LSP)

// The Liskov Substitution Principle states that objects of a derived class should
// be able to replace objects of the base class without affecting the correctness of the program.
// In Golang, this principle applies to interfaces and their implementations,
// ensuring that the code remains consistent and reliable.

package main

type Bird interface {
	MakeSound() string
}

type FlyingBird interface {
	Bird
	Fly() string
}

type Pigeon struct{}

func (p *Pigeon) MakeSound() string {
	return "Coo"
}

func (p *Pigeon) Fly() string {
	return "Pigeon is flying."
}

type Penguin struct{}

func (p *Penguin) MakeSound() string {
	return "Squawk"
}

func main() {

}
