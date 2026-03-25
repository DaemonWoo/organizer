package main

import (
	"encoding/json"
	"os"
)

// DefaultRules определяет маппинг категорий и расширений
var DefaultRules = map[string][]string {
	"Pictures": {".jpg", ".png", ".gif", ".jpeg", ".svg", ".webp"},
	"Music":    {".mp3", ".wav", ".flac", ".ogg", ".aac"},
	"Media":    {".mp4", ".mov", ".avi", ".mkv", ".webm"},
	"Docs":     {".pdf", ".docx", ".doc", ".txt", ".xlsx", ".xls", ".pptx"},
	"Archive":  {".zip", ".rar", ".7z", ".tar", ".gz", ".bz2"},
	"Exec":     {".exe", ".msi", ".dmg", ".pkg"},
}

// getConfig пытается прочитать config.json, иначе возвращает дефолт
func getConfig(filename string) map[string][]string {
	file, err := os.ReadFile(filename)
	if err != nil {
		// Если файла нет, возвращаем дефолт
		return DefaultRules
	}

	var config map[string][]string
	if err := json.Unmarshal(file, &config); err != nil {
		// Если файл битый, тоже лучше откатиться на дефолт
		return DefaultRules
	}

	return config
}

// ExtensionToCategory создаёт обратный маппинг для быстрого поиска
func ExtensionToCategory() map[string]string {
	rules := getConfig("config.json")
	res := make(map[string]string)
	for category, exts := range rules {
		for _, ext := range exts {
			res[ext] = category
		}
	}
	return res
}