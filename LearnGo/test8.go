//Interface
package main

import (
	"errors"
	"fmt"
)

// type Taxi interface {
// 	isEmpty() bool
// 	avliable() bool
// 	showCusNumber() int
// }

type Robot interface {
	PowerOn() error
	Talk() string
}

type T850 struct {
	Name string
}

func (a *T850) PowerOn() error {
	return nil
}

func (a *T850) Talk() string {
	return "I'm T850"
}

type R2D2 struct {
	Broken bool
}

func (r *R2D2) PowerOn() error {
	if r.Broken {
		return errors.New("R2D2 is broken")
	} else {
		return nil
	}
}

func (r *R2D2) Talk() string {
	return "I'm R2D2"
}

func Boot(robot Robot) (string, error) {
	return robot.Talk(), robot.PowerOn()
}

func main() {
	t := T850{
		Name: "Frank",
	}

	r := R2D2{
		Broken: true,
	}

	r_talk, err := Boot(&r)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Robot is power on!")
	}
	fmt.Println(r_talk)

	t_talk, err2 := Boot(&t)

	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println("Robot is power on!")
	}
	fmt.Println(t_talk)
}
