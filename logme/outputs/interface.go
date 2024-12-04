package outputs

import "github.com/rockkley/logme/logme/entity"

type LogOutput interface {
	Write(message *entity.Message) error
}
