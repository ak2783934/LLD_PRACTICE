package main

import (
	"fmt"
)

type Computer struct {
	CPU     string
	RAM     string
	Storage string
	GPU     string // Optional
}

type ComputerBuilder interface {
	SetCPU()
	SetRAM()
	SetStorage()
	SetGPU() // Optional
	GetComputer() *Computer
}

type GamingComputerBuilder struct {
	computer *Computer
}

func NewGamingComputerBuilder() *GamingComputerBuilder {
	return &GamingComputerBuilder{&Computer{}}
}

func (b *GamingComputerBuilder) SetCPU() {
	b.computer.CPU = "Intel Core i9"
}

func (b *GamingComputerBuilder) SetRAM() {
	b.computer.RAM = "32GB"
}

func (b *GamingComputerBuilder) SetStorage() {
	b.computer.Storage = "1TB NVMe SSD"
}

func (b *GamingComputerBuilder) SetGPU() {
	b.computer.GPU = "NVIDIA RTX 4080"
}

func (b *GamingComputerBuilder) GetComputer() *Computer {
	return b.computer
}

type ComputerBuilderChain struct {
	computer *Computer
}

func NewComputerBuilderChain() *ComputerBuilderChain {
	return &ComputerBuilderChain{&Computer{}}
}

func (b *ComputerBuilderChain) SetCPU(cpu string) *ComputerBuilderChain {
	b.computer.CPU = cpu
	return b
}

func (b *ComputerBuilderChain) SetRAM(ram string) *ComputerBuilderChain {
	b.computer.RAM = ram
	return b
}

func (b *ComputerBuilderChain) SetStorage(storage string) *ComputerBuilderChain {
	b.computer.Storage = storage
	return b
}

func (b *ComputerBuilderChain) SetGPU(gpu string) *ComputerBuilderChain {
	b.computer.GPU = gpu
	return b
}

func (b *ComputerBuilderChain) Build() *Computer {
	return b.computer
}

// Usage

func main() {
	computer := NewComputerBuilderChain().
		SetCPU("AMD Ryzen 9").
		SetRAM("64GB").
		SetStorage("2TB SSD").
		SetGPU("NVIDIA RTX 4090").
		Build()

	fmt.Printf("Custom Built Computer: %+v\n", computer)
}
