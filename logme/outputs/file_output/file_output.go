package file_output

import (
	"os"
	"time"
)

type FileOutput struct {
	defaultFolder string
	FilePath      string
}

func (f *FileOutput) Write(message string) error {
	return os.WriteFile(f.getCurrentDate()+" "+f.FilePath, []byte(f.getCurrentDateTime()+message+"\n"), 0644)
}

func (f *FileOutput) getCurrentDateTime() string {
	return time.Now().Format(time.DateTime)
}

func (f *FileOutput) getCurrentDate() string {
	return time.Now().Format(time.DateOnly)
}
