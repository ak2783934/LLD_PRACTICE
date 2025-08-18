package main

import "fmt"

type Car interface {
	Assemble()
}

type CarSpecification interface {
	Display()
}

// Assemble a sedan for north america.
type Sedan struct{}

func (s *Sedan) Assemble() {
	fmt.Println("assmble a sedan")
}

type NorthAmericaSpecification struct{}

func (n *NorthAmericaSpecification) Display() {
	fmt.Println("North America Specification: Safety features compliant with local regulations.")
}

type Hatchback struct{}

func (h *Hatchback) Assemble() {
	fmt.Println("Assembling Hatchback car.")
}

type EuropeSpecification struct{}

func (e *EuropeSpecification) Display() {
	fmt.Println("Europe Specification: Fuel efficiency and emissions compliant with EU standards.")
}

type CarFactory interface {
	CreateCar() Car
	CreateSpecification() CarSpecification
}

type NorthAmericaCarFactory struct{}

func (n *NorthAmericaCarFactory) CreateCar() Car {
	return &Sedan{}
}

func (n *NorthAmericaCarFactory) CreateSpecification() CarSpecification {
	return &NorthAmericaSpecification{}
}

type EuropeCarFactory struct{}

func (f *EuropeCarFactory) CreateCar() Car {
	return &Hatchback{}
}

func (f *EuropeCarFactory) CreateSpecification() CarSpecification {
	return &EuropeSpecification{}
}

func main() {
	var factory CarFactory

	// Producing for North America
	factory = &NorthAmericaCarFactory{}
	car := factory.CreateCar()
	spec := factory.CreateSpecification()

	car.Assemble()
	spec.Display()

	// Producing for Europe
	factory = &EuropeCarFactory{}
	car = factory.CreateCar()
	spec = factory.CreateSpecification()

	car.Assemble()
	spec.Display()
}
