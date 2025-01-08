package logme

import (
	"github.com/rockkley/logme/logme/entity/levels"
	"sync"
)

type LevelSetter struct {
	level levels.LogLevel
	mu    sync.RWMutex
}

func (ls *LevelSetter) Info() {
	ls.mu.Lock()
	defer ls.mu.Unlock()
	ls.level = levels.Info

}

func (ls *LevelSetter) Warning() {
	ls.mu.Lock()
	defer ls.mu.Unlock()
	ls.level = levels.Warning
}

func (ls *LevelSetter) Debug() {
	ls.mu.Lock()
	defer ls.mu.Unlock()
	ls.level = levels.Debug
}

func (ls *LevelSetter) Critical() {
	ls.mu.Lock()
	defer ls.mu.Unlock()
	ls.level = levels.Critical
}

func (ls *LevelSetter) All() {
	ls.mu.Lock()
	defer ls.mu.Unlock()
	ls.level = levels.All
}
