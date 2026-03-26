package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"organizer/menu"
	"organizer/utils"
)

type MoveTask struct {
	OldPath string
	NewPath string
	cat string
}

var presetDirs = map[string]string {
	"1": "Downloads",
	"2": "Desktop",
	"3": "Documents",
}

func main() {
	home, err := os.UserHomeDir()
	if err != nil {
		home = "/tmp"
	}
	dirToSort := menu.SelectDirectory()
	if dirToSort == "" {
		fmt.Println("No directory selected, exiting...", dirToSort)
		return
	} else {
		dirToSort = filepath.Join(home, dirToSort)
	}

	fastMap := ExtensionToCategory();

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
			catPath := filepath.Join(home, category);
            newPath := utils.GetUniquePath(catPath, fileName)
			
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
		catPath := filepath.Join(home, t.cat)
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
