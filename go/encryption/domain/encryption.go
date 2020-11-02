package domain

type CharReader interface {
	Read() string
}

type CharWriter interface {
	Write(string)
}

func Encryption(reader CharReader, writer CharWriter) {
	input := reader.Read()
	encrypted := translate(input)
	writer.Write(encrypted)
}

func translate(input string) string {
	return input
}

