package main

import (
	"github.com/SpringMT/clean-architecture/go/encryption/domain"
	"github.com/SpringMT/clean-architecture/go/encryption/usecase"
)

func main() {
	domain.Encryption(&usecase.ConsoleReader{}, &usecase.ConsoleWriter{})
}