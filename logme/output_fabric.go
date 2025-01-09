package logme

import (
	"github.com/rockkley/logme/logme/outputs"
	"github.com/rockkley/logme/logme/outputs/console_output"
	"github.com/rockkley/logme/logme/outputs/file_output"
	"sync"
)

type OutputFabric struct {
	outputs []outputs.LogOutput
	mu      sync.RWMutex
}

func NewOutputFabric() *OutputFabric {
	return &OutputFabric{}
}

func (of *OutputFabric) ConsoleOutput() {
	output := console_output.NewConsoleOutput()
	of.mu.Lock()
	defer of.mu.Unlock()
	of.outputs = append(of.outputs, output)

}

func (of *OutputFabric) FileOutput(filepath string) {
	output := file_output.NewFileOutput(filepath)

	of.mu.Lock()
	defer of.mu.Unlock()
	of.outputs = append(of.outputs, output)
}

func (of *OutputFabric) GetOutputs() []outputs.LogOutput {
	return of.outputs
}
