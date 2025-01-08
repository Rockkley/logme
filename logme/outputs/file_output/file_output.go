package file_output

import (
	"github.com/rockkley/logme/logme/entity"
	"os"
	"strings"
	"time"
)

type FileOutput struct {
	FilePath string
}

func NewFileOutput(filePath string) *FileOutput {
	return &FileOutput{FilePath: filePath}
}

func (f *FileOutput) Write(message *entity.Message) error {
	file, err := os.OpenFile(f.getCurrentDate()+" "+f.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	str := strings.Join([]string{message.Timestamp, message.Level.String(), message.Text, "\n"}, " ")
	_, err = file.WriteString(str)

	return err
}

func (f *FileOutput) getCurrentDate() string {
	return time.Now().Format(time.DateOnly)
}
