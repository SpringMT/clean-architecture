package usecase

import "fmt"

type ConsoleWriter struct {
}

func (r *ConsoleWriter) Write(encrypted string) {
	fmt.Println(encrypted)
}