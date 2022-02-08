package main

import "fmt"

type animal interface {
	mover
	eat
}

type mover interface {
	move()
}

type eat interface {
	eat(string)
}

type cat struct {
	name string
	feet int
}

type chicken struct {
	feet int
}

func (c cat) move() {
	fmt.Println("走猫步")
}
func (c cat) eat(food string) {
	fmt.Println("吃", food)
}

func (c chicken) move() {
	println("鸡冻")
}

func (c chicken) eat(food string) {
	println("吃", food)
}

func main() {
	var a1 animal
	bc := cat{name: "淘气", feet: 4}

	a1 = bc
	a1.eat("鱼")
	a1.move()

	c1 := chicken{feet: 2}
	a1 = c1
	a1.eat("虫子")
	a1.move()

	c2 := &chicken{feet: 2}
	a1 = c2
	a1.eat("饲料")
	a1.move()

}
