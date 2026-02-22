package main

import "fmt"

type Walking interface {
	walk()
}

type WalkingFly interface {
	Walking
	fly()
}

type Panguin struct {
	name string
}

type Bird struct {
	name string
}

func (p *Panguin) walk() {
	fmt.Printf("%v can walk\n", p.name)
}

func (b *Bird) walk() {
	fmt.Printf("%v can walk\n", b.name)
}

func (b *Bird) fly() {
	fmt.Printf("%v can fly\n", b.name)
}

func CheckBird(bird interface{}) {
	switch data := bird.(type) {
	case WalkingFly:
		data.fly()
		data.walk()
	case Walking:
		data.walk()

	}
}

func main() {
	bird := Bird{
		name: "mamad",
	}
	CheckBird(&bird)

	bird2 := Panguin{
		name: "javad",
	}
	CheckBird(&bird2)
}
