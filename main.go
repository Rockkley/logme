package main

import (
	"github.com/rockkley/logme/logme"
	"github.com/rockkley/logme/logme/outputs/console_output"
)

func main() {

	logger := logme.NewLogMe()
	// logger.AddOutput(&file_output.FileOutput{FilePath: "mylog.txt"})
	logger.AddOutput(console_output.NewConsoleOutput()) // TODO добавление output из стандартного набора сделать проще

	logger.Info("Из-за леса, из-за гор, показал мужик топор..")
	logger.Warning("..но не просто показал!..")
	logger.Debug("(прищурься)")
	logger.Critical("его к хую привязал!")
}
