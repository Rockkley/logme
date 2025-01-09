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
	ls.setLevel(levels.Info)
}

func (ls *LevelSetter) Warning() {
	ls.setLevel(levels.Warning)
}

func (ls *LevelSetter) Debug() {
	ls.setLevel(levels.Debug)
}

func (ls *LevelSetter) Critical() {
	ls.setLevel(levels.Critical)
}

func (ls *LevelSetter) All() {
	ls.setLevel(levels.All)
}

func (ls *LevelSetter) setLevel(level levels.LogLevel) {
	ls.mu.Lock()
	defer ls.mu.Unlock()
	ls.level = level
}
