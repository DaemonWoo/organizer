package organizer

// CategoryRules определяет маппинг категорий и расширений
var CategoryRules = map[string][]string{
	"Pictures": {".jpg", ".png", ".gif", ".jpeg", ".svg", ".webp"},
	"Music":    {".mp3", ".wav", ".flac", ".ogg", ".aac"},
	"Media":    {".mp4", ".mov", ".avi", ".mkv", ".webm"},
	"Docs":     {".pdf", ".docx", ".doc", ".txt", ".xlsx", ".xls", ".pptx"},
	"Archive":  {".zip", ".rar", ".7z", ".tar", ".gz", ".bz2"},
	"Exec":     {".exe", ".msi", ".dmg", ".pkg"},
}

// ExtensionToCategory создаёт обратный маппинг для быстрого поиска
func ExtensionToCategory() map[string]string {
	res := make(map[string]string)
	for category, exts := range CategoryRules {
		for _, ext := range exts {
			res[ext] = category
		}
	}
	return res
}