package main

import (
	"fmt"
	"errors"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	name       string
	food       string
	locomotion string
	noise      string
}

type Bird struct {
	name       string
	food       string
	locomotion string
	noise      string
}

type Snake struct {
}


func (cow Cow) Eat() {
	fmt.Println(cow.food)
}

func (cow Cow) Move() {
	fmt.Println(cow.locomotion)
}
func (cow Cow) Speak() {
	fmt.Println(cow.noise)
}

func (bird Bird) Eat() {
	fmt.Println(bird.food)
}

func (bird Bird) Move() {
	fmt.Println(bird.locomotion)
}
func (bird Bird) Speak() {
	fmt.Println(bird.noise)
}

func NewCow(name string) *Cow {
	return &Cow{name: name, food: "grass", locomotion: "walk", noise: "moo"}
}


func NewBird(name string) *Bird {
	return &Bird{name: name, food: "worms", locomotion: "fly", noise: "peep"}
}


func FindAnimalByName(zoo []Animal, name string) (Animal, error) {
	for _, animal := range zoo {
		switch sh := animal.(type) {
		case Cow:
			if sh.name == name{
				return sh,nil
			} else{
				return nil, errors.New("not found")
			} 
		case Bird:
			if sh.name == name{
				return sh,nil
			} else{
				return nil, errors.New("not found")
			} 
		default:
			return nil, errors.New("animal not found")
		}

	}
	return nil,errors.New("animal not found")
}

func main() {

	var cow = NewCow("vaca lechera")
	var bird = NewBird("piolin")

	zoo := []Animal{}

	zoo = append(zoo, *cow)

	zoo = append(zoo, *bird)

	found, err := FindAnimalByName(zoo, "vaca lechera")

	fmt.Println(found)

	if (err == nil){
		found.Eat()
	}


}
