package iomanager

type IOManager interface {
	Read() ([]string, error)
	Write(data interface{}) error
}
