package menu

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
    choices  []string        
    cursor   int              
    selected string          
}

func initialModel() model {
    return model{
        choices: []string{"Downloads", "Desktop", "Documents", "Custom Path"},
    }
}

func (m model) Init() tea.Cmd {
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "ctrl+c", "q", "й":
            return m, tea.Quit
        case "up", "k":
            if m.cursor > 0 {
                m.cursor--
            } else {
                m.cursor = len(m.choices) - 1
            }
        case "down", "j":
            if m.cursor < len(m.choices)-1 {
                m.cursor++
            } else {
                m.cursor = 0
            }
        case "enter", " ":
            m.selected = m.choices[m.cursor]
            return m, tea.Quit
        }
    }
    return m, nil
}

func (m model) View() string {
    var s strings.Builder; 
    s.WriteString("📂 Выберите папку для сортировки:\n\n")

    for i, choice := range m.choices {
        cursor := " "
        if m.cursor == i {
            cursor = ">"
        }
        fmt.Fprintf(&s, "%s [%s]\n", cursor, choice)
    }

    s.WriteString("\n(нажмите q для выхода)\n")
    return s.String()
}

func SelectDirectory() (string) {
	p := tea.NewProgram(initialModel())
    m, err := p.Run()
    if err != nil {
        fmt.Printf("Ошибка запуска: %v", err)
        os.Exit(1)
    }

    finalModel := m.(model)
    
	if finalModel.selected == "Custom Path" {
		fmt.Print("Введите полный путь: ")
		var customPath string
		fmt.Scanln(&customPath)
		customPath = strings.TrimSpace(customPath)
		
		if strings.HasPrefix(customPath, "~") {
			home, err := os.UserHomeDir()
			if err == nil {
				customPath = filepath.Join(home, customPath[1:])
			}
		}
		return customPath
	}
	
	return finalModel.selected
}