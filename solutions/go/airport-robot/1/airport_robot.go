package airportrobot

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.

import "fmt"

type Greeter interface {
    LanguageName() string
    Greet (name string) string
}


type Italian struct {
    Name string
}

func (i Italian) LanguageName() string {
    return fmt.Sprintf("Italian")
}

func (i Italian) Greet (name string) string {
    return fmt.Sprintf("I can speak Italian: Ciao %s!", name)
}


type Portuguese struct {
    Name string
}

func (i Portuguese) LanguageName() string {
    return fmt.Sprintf("Portuguese")
}

func (i Portuguese) Greet (name string) string {
    return fmt.Sprintf("I can speak Portuguese: Olá %s!", name)
}

func SayHello(name string, g Greeter) string {
    return g.Greet(name)
}




