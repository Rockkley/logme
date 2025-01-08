package main

import (
	"github.com/rockkley/logme/logme"
)

func main() {
	logger := logme.NewLogMe()
	logger.SetLevel()
	logger.AddOutput().ConsoleOutput()
	logger.AddOutput().FileOutput("mylog.txt")

	logger.Info(12, 21, 219, 45)
	logger.Info("Из-за леса, из-за гор, показал мужик топор..")
	logger.Warning("..но не просто показал!..")
	logger.Debug("(прищурься)")
	logger.Critical("его к хую привязал!")
}
