package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type MoveTask struct {
	OldPath string
	NewPath string
	cat string
}
func reverseMap(src map[string][]string) map[string]string {
    res := make(map[string]string)
    for category, exts := range src {
        for _, ext := range exts {
            res[ext] = category
        }
    }
    return res
}

/// Returns unique path for file in [dir].
func getUniquePath(dir, name string) string {
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

var extensions = map[string][]string {
	"Pictures": {".jpg", ".png", ".gif", ".jpeg"},
	"Music": {".mp3", ".wav", ".flac", ".ogg"},
	"Media": {".mp4", ".mov", ".avi", ".mkv"},
	"Docs": {".pdf", ".docx", ".txt", ".xlsx"},
	"Archive":  {".zip", ".rar", ".7z", ".tar", ".gz"},
}

func main() {
	home, err := os.UserHomeDir()
	if err != nil {
		home = "/tmp"
	}
	basePath := home + string(filepath.Separator)
	/// TODO: config.json
	dirToSort := basePath + "Downloads"
    fastMap := reverseMap(extensions);

	if len(os.Args) > 1 {
		dirToSort = os.Args[1]
	}
	info, err := os.Stat(dirToSort)
	if err != nil || !info.IsDir() {
		fmt.Printf("❌ Ошибка: путь '%s' не найден\n", dirToSort)
		return
	}

    files, err := os.ReadDir(dirToSort)
    if err != nil {
		fmt.Println("❌ Ошибка чтения папки:", err)
		return
	}
	var tasks []MoveTask

	stats := make(map[string]int)

    fmt.Println("🚀 Начинаю сортировку...")

    for _, file := range files {
        if file.IsDir() {
            continue
        }
        fileName := file.Name()
        ext := strings.ToLower(filepath.Ext(fileName))

        if category, ok := fastMap[ext]; ok {
            
            

			oldPath := filepath.Join(dirToSort, fileName)
			catPath := filepath.Join(basePath + category);
            newPath := getUniquePath(catPath, fileName)
			
			tasks = append(tasks, MoveTask{
				OldPath: oldPath, 
				NewPath: newPath,
				cat: category,
			})
			stats[category]++

            
        } else {
			fmt.Printf("⏩ Пропущен: %s\n", fileName)
		}
    }

	if len(tasks) == 0 {
		fmt.Println("✨ Файлов для сортировки не найдено.")
		return
	}

	fmt.Printf("📂 Анализ папки: %s\n", dirToSort)
	fmt.Println("📊 План перемещения:")
	for cat, count := range stats {
		fmt.Printf("  - %-10s: %d шт.\n", cat, count)
	}
	fmt.Printf("\nИтого к перемещению: %d\n", len(tasks))
	fmt.Print("✅ Продолжить? (y/n): ")

	var confirm string
	fmt.Scanln(&confirm)
	if strings.ToLower(confirm) != "y" {
		fmt.Println("🛑 Операция отменена.")
		return
	}

	fmt.Println("\n🚀 Начинаю работу...")

	for _, t := range tasks {
		catPath := filepath.Join(basePath + t.cat)
		os.MkdirAll(catPath, 0755)

		err := os.Rename(t.OldPath, t.NewPath)
		if err != nil {
			fmt.Printf("❌ Ошибка %s: %v\n", filepath.Base(t.OldPath), err)
		} else {	
			fmt.Printf("📦 %s -> %s\n", filepath.Base(t.OldPath), t.cat)
		}
	}

	fmt.Println("\n✨ Сортировка завершена!")
}
