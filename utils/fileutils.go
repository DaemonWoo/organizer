package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

/// Returns unique path for file in [dir].
func GetUniquePath(dir, name string) string {
	ext := filepath.Ext(name)
	base := strings.TrimSuffix(name, ext)
	targetPath := filepath.Join(dir, name)
	
	counter := 1
	for {
		// Проверяем, существует ли файл
		if _, err := os.Stat(targetPath); os.IsNotExist(err) {
			break // Путь свободен!
		}
		// Если занят — пробуем новое имя: base (1).ext
		newName := fmt.Sprintf("%s (%d)%s", base, counter, ext)
		targetPath = filepath.Join(dir, newName)
		counter++
	}
	return targetPath
}
