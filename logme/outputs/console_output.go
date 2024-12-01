package outputs

import "fmt"

type ConsoleOutput struct{}

func (c *ConsoleOutput) Write(message string) (err error) {
	_, err = fmt.Println(message)
	return
}
