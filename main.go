package main

import (
	"github.com/rockkley/logme/logme"
	"github.com/rockkley/logme/logme/outputs"
)

func main() {
	logger := logme.NewLogMe()
	logger.SetLevel(logme.All)
	logger.AddOutput(&outputs.FileOutput{FilePath: "mylog.txt"})
	logger.AddOutput(&outputs.ConsoleOutput{})
	logger.Info("Из-за леса, из-за гор, показал мужик топор..")
	logger.Warning("..но не просто показал!..")
	logger.Debug("(прищурься)")
	logger.Critical("его к хую привязал!")
}
