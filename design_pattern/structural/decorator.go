package main

import "fmt"

/*
	https://refactoring.guru/design-patterns/decorator
	https://github.com/iamgoangle/design-patterns/blob/master/Decorator.md
*/
func main() {
	decorator := &Facebook{
		&Component{},
	}

	decorator.Operation()
	decorator.component.Operation()
}

type Facebook struct {
	component *Component
}

func (f *Facebook) Operation() {
	fmt.Println("Send Facebook")
}

type Slack struct {
	component *Component
}

func (s *Slack) Operation() {
	fmt.Println("Send Slack")
}

type Component struct {
}

func (c *Component) Operation() {
	fmt.Println("Send e-mail")
}
