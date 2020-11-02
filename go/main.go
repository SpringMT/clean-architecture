package main

import (
	"github.com/SpringMT/clean-architecture/go/domain"
	"github.com/SpringMT/clean-architecture/go/usecase"
)

func main() {
	domain.Encryption(&usecase.ConsoleReader{}, &usecase.ConsoleWriter{})

}