package main

import (
	"fmt"
	s "sayhello/say"
)

func main() {

	fmt.Println("say Hello!")

	var saynoformat s.Greeter
	fmt.Println(saynoformat.Hello("Irfan"))

	say := s.Greeter{
		Format: "Hello There, %s",
	}
	fmt.Println(say.Hello("Irfan"))

}
