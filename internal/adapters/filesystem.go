package adapters

import (
	"os"
	"path/filepath"
)

type FileSystem struct{}

func (FileSystem) WriteFile(path string, content []byte) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	return os.WriteFile(path, content, 0o644)
}

func (FileSystem) ReadFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}