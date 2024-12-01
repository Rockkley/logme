package outputs

type LogOutput interface {
	Write(message string) error
}
