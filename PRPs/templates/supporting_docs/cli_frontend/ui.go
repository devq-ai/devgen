package main

import (
    "fmt"
    "strings"
    "time"

    "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/glamour"
    "github.com/charmbracelet/lipgloss"
)

var (
    titleStyle = lipgloss.NewStyle().
        Bold(true).
        Foreground(lipgloss.Color("205")).
        Padding(1, 2)

    branchStyle = lipgloss.NewStyle().
        Border(lipgloss.RoundedBorder()).
        Padding(1).
        Margin(1, 0)

    completedStyle = lipgloss.NewStyle().
        Foreground(lipgloss.Color("10"))

    runningStyle = lipgloss.NewStyle().
        Foreground(lipgloss.Color("11"))

    pendingStyle = lipgloss.NewStyle().
        Foreground(lipgloss.Color("8"))

    failedStyle = lipgloss.NewStyle().
        Foreground(lipgloss.Color("9"))
)

type model struct {
    playbook *Playbook
    engine   *Engine
    markdown *glamour.TermRenderer
    logs     []string
    running  bool
}

type tickMsg time.Time

func newUI(playbook *Playbook) *model {
    renderer, _ := glamour.NewTermRenderer(
        glamour.WithAutoStyle(),
        glamour.WithWordWrap(80),
    )

    return &model{
        playbook: playbook,
        engine:   newEngine(playbook),
        markdown: renderer,
        logs:     []string{},
        running:  false,
    }
}

func (m *model) Init() tea.Cmd {
    return tea.Batch(
        m.tick(),
        tea.Sequence(
            tea.Tick(time.Second, func(t time.Time) tea.Msg {
                return tickMsg(t)
            }),
        ),
    )
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "ctrl+c", "q":
            return m, tea.Quit
        case " ", "enter":
            if !m.running && !m.engine.isPlaybookComplete() {
                m.running = true
                return m, m.executeNextSteps()
            }
        }

    case tickMsg:
        return m, tea.Tick(time.Second, func(t time.Time) tea.Msg {
            return tickMsg(t)
        })

    case stepCompleteMsg:
        m.addLog(fmt.Sprintf("✓ %s: %s completed", 
            m.playbook.Branches[msg.branchIdx].Name,
            m.playbook.Branches[msg.branchIdx].Steps[msg.stepIdx].Agent))
        
        if m.engine.isPlaybookComplete() {
            m.running = false
            m.addLog("🎉 Playbook execution completed!")
        } else {
            return m, m.executeNextSteps()
        }
    }

    return m, nil
}

func (m *model) View() string {
    var b strings.Builder

    // Title
    title := fmt.Sprintf("# %s\n", m.playbook.Name)
    rendered, _ := m.markdown.Render(title)
    b.WriteString(rendered)

    // Controls
    if !m.running && !m.engine.isPlaybookComplete() {
        b.WriteString("Press SPACE to execute next steps, 'q' to quit\n\n")
    } else if m.engine.isPlaybookComplete() {
        b.WriteString("Execution complete! Press 'q' to quit\n\n")
    } else {
        b.WriteString("Executing... Press 'q' to quit\n\n")
    }

    // Branches
    for _, branch := range m.playbook.Branches {
        branchContent := m.renderBranch(branch)
        b.WriteString(branchStyle.Render(branchContent))
        b.WriteString("\n")
    }

    // Logs
    if len(m.logs) > 0 {
        b.WriteString("## Recent Activity\n")
        for i := len(m.logs) - 5; i < len(m.logs); i++ {
            if i >= 0 {
                b.WriteString(fmt.Sprintf("• %s\n", m.logs[i]))
            }
        }
    }

    return b.String()
}

func (m *model) renderBranch(branch Branch) string {
    var b strings.Builder
    
    b.WriteString(fmt.Sprintf("**%s**\n\n", branch.Name))
    
    for _, step := range branch.Steps {
        status := m.getStatusIcon(step.Status)
        style := m.getStatusStyle(step.Status)
        
        line := fmt.Sprintf("%s %s: %s (condition: %s)", 
            status, step.Agent, step.Action, step.Condition)
        b.WriteString(style.Render(line))
        b.WriteString("\n")
    }
    
    return b.String()
}

func (m *model) getStatusIcon(status StepStatus) string {
    switch status {
    case StatusCompleted:
        return "✓"
    case StatusRunning:
        return "⟳"
    case StatusFailed:
        return "✗"
    default:
        return "○"
    }
}

func (m *model) getStatusStyle(status StepStatus) lipgloss.Style {
    switch status {
    case StatusCompleted:
        return completedStyle
    case StatusRunning:
        return runningStyle
    case StatusFailed:
        return failedStyle
    default:
        return pendingStyle
    }
}

func (m *model) tick() tea.Cmd {
    return tea.Tick(time.Second, func(t time.Time) tea.Msg {
        return tickMsg(t)
    })
}

func (m *model) executeNextSteps() tea.Cmd {
    return func() tea.Msg {
        steps := m.engine.getNextExecutableSteps()
        
        if len(steps) == 0 {
            m.running = false
            return nil
        }

        // Execute first available step
        step := steps[0]
        err := m.engine.executeStep(step.BranchIndex, step.StepIndex)
        if err != nil {
            m.addLog(fmt.Sprintf("✗ Error: %s", err.Error()))
            return nil
        }

        return stepCompleteMsg{
            branchIdx: step.BranchIndex,
            stepIdx:   step.StepIndex,
        }
    }
}

func (m *model) addLog(message string) {
    m.logs = append(m.logs, message)
    if len(m.logs) > 20 {
        m.logs = m.logs[1:]
    }
}

func (m *model) run() error {
    p := tea.NewProgram(m, tea.WithAltScreen())
    _, err := p.Run()
    return err
}

type stepCompleteMsg struct {
    branchIdx int
    stepIdx   int
}