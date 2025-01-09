package main

import (
	"github.com/rockkley/logme/logme"
)

func main() {
	logger := logme.NewLogMe()
	logger.SetLevel()
	logger.AddOutput().ConsoleOutput()
	logger.AddOutput().FileOutput("mylog.txt")
	logger.AddHook("server", "her")
	logger.AddHook("AUTHOR", "ME")
	//logger.Info(12, 21, 219, 45)
	logger.Info("Из-за леса, из-за гор, показал мужик топор..")
	//logger.Warning("..но не просто показал!..")
	//logger.Debug("(прищурься)")
	//logger.Critical("его к хую привязал!")
}

func test(logger *logme.LogMe) {
	logger.Debug("start ")

	for i := range 100000 {
		if i%2 == 0 {
			logger.Info(i)
		} else {
			logger.Warning(i)
		}
	}
	logger.Debug("end ")
}
