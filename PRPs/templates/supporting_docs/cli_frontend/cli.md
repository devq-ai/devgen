// Required CLI structure pattern
// cli/internal/ui/styles.go
package ui

import "github.com/charmbracelet/lipgloss"

var (
    // Required base styles for consistency
    BaseStyle = lipgloss.NewStyle().
        PaddingLeft(1).
        PaddingRight(1)
    
    ErrorStyle = BaseStyle.Copy().
        Foreground(lipgloss.Color("196")).
        Bold(true)
    
    SuccessStyle = BaseStyle.Copy().
        Foreground(lipgloss.Color("46")).
        Bold(true)
)

// Required Bubble Tea component pattern
// cli/internal/ui/components/input.go
package components

import (
    "github.com/charmbracelet/bubbles/textinput"
    tea "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/lipgloss"
)

type InputModel struct {
    textInput textinput.Model
    err       error
}

func NewInputModel() InputModel {
    ti := textinput.New()
    ti.Focus()
    return InputModel{textInput: ti}
}

func (m InputModel) Init() tea.Cmd {
    return textinput.Blink
}

func (m InputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    // Required: handle all input properly
    var cmd tea.Cmd
    m.textInput, cmd = m.textInput.Update(msg)
    return m, cmd
}

func (m InputModel) View() string {
    return m.textInput.View()
}

// Required logging pattern
// cli/internal/commands/base.go
package commands

import (
    "github.com/charmbracelet/log"
    "github.com/charmbracelet/lipgloss"
)

// Configure charm log with Logfire backend
func setupCLILogging() {
    logger := log.NewWithOptions(os.Stderr, log.Options{
        ReportCaller: true,
        ReportTimestamp: true,
        TimeFormat: time.Kitchen,
    })
    
    // Custom styles for CLI
    styles := log.DefaultStyles()
    styles.Levels[log.InfoLevel] = lipgloss.NewStyle().
        SetString("INFO").
        Foreground(lipgloss.Color("86"))
    
    logger.SetStyles(styles)
    log.SetDefault(logger)
}
