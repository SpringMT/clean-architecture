package usecase

type ConsoleReader struct {
}

func (r *ConsoleReader) Read() string  {
	return "dummy string"
}
