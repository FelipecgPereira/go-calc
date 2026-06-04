package iomanager

type IOManager interface {
	Read() (string, error)
	Write(data string) error
}
