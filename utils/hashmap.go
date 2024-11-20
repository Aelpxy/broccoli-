package utils

import "sync"

type FileMap struct {
	mu   sync.RWMutex
	hash map[string]string
}

func NewFileMap() *FileMap {
	return &FileMap{
		hash: make(map[string]string),
	}
}

func (f *FileMap) Add(key, path string) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.hash[key] = path
}

func (f *FileMap) Get(key string) string {
	f.mu.RLock()
	defer f.mu.RUnlock()
	return f.hash[key]
}

func (f *FileMap) Delete(key string) {
	f.mu.Lock()
	defer f.mu.Unlock()
	delete(f.hash, key)
}

func (f *FileMap) Exists(key string) bool {
	f.mu.RLock()
	defer f.mu.RUnlock()
	_, exists := f.hash[key]
	return exists
}
