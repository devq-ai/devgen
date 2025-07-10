## PRP 3: CLI with Charm.sh Components

```markdown
# PRP: Machina CLI with Charm.sh Components

## Goal
Build a unified Go CLI for Machina using Charm.sh components (Bubble Tea, Lip Gloss, Gum) that provides interactive management of MCP servers and registry operations.

## Why
- **Unified Interface**: Single CLI for all Machina operations across the DevQ.ai ecosystem
- **User Experience**: Beautiful, interactive terminal interface using Charm.sh components
- **Developer Productivity**: Quick access to MCP server management, health monitoring, and registry operations
- **Consistency**: Follows DevQ.ai CLI patterns with trustlis authentication integration

## What
A Go CLI application with:
- Interactive TUI for MCP server management using Bubble Tea
- Consistent styling with Lip Gloss using Cyber & Pastel Design System
- Shell script integration with Gum components
- Charm Log integration with Logfire backend
- Subcommands for machina operations within the unified devqai CLI

## All Needed Context

### CLI Structure (from devqai-tree.txt)
```
cli/
├── internal/
│   ├── commands/
│   │   ├── machina/      # machina subcommands
│   │   ├── agentical/    # agentical subcommands  
│   │   ├── ptolemies/    # ptolemies subcommands
│   │   └── trustlis/     # trustlis subcommands
│   ├── auth/
│   │   └── trustlis.go   # trustlis integration
│   └── ui/               # charm ui components
│       ├── styles.go     # lip gloss styles
│       ├── components/   # reusable bubble tea components
│       └── tui/         # full tui applications
```

### Charm.sh Patterns (from cli.md)
```go
import (
    "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/lipgloss"
    "github.com/charmbracelet/log"
)

// Required base styles for consistency
var BaseStyle = lipgloss.NewStyle().
    PaddingLeft(1).
    PaddingRight(1)

// Required Bubble Tea component pattern
type InputModel struct {
    textInput textinput.Model
    err       error
}

func (m InputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var cmd tea.Cmd
    m.textInput, cmd = m.textInput.Update(msg)
    return m, cmd
}
```

### Design System (from style_guide.md)
```go
// Cyber Black Palette
var (
    CyberVoidBlack   = lipgloss.Color("#000000")
    CyberMatrixGreen = lipgloss.Color("#00ff00") 
    CyberNeonPink    = lipgloss.Color("#ff0080")
    CyberElectricCyan = lipgloss.Color("#00ffff")
)

// Pastel Black Palette  
var (
    PastelMintGreen   = lipgloss.Color("#a8e6a3")
    PastelBlushPink   = lipgloss.Color("#ffb3ba")
    PastelSkyBlue     = lipgloss.Color("#b3e5fc")
    PastelCreamYellow = lipgloss.Color("#fff9c4")
)
```

### Authentication (from trustlis integration)
- CLI must integrate with trustlis for authenticated operations
- No API calls to DevQ.ai backend without proper authentication
- Trustlis client integration for auth token management

## Implementation Blueprint

### Phase 1: Core CLI Structure
```go
// cli/internal/commands/machina/root.go
package machina

import (
    "github.com/spf13/cobra"
    "github.com/charmbracelet/log"
    "github.com/charmbracelet/lipgloss"
)

var (
    // Cyber theme styles for machina
    titleStyle = lipgloss.NewStyle().
        Foreground(lipgloss.Color("#00ffff")).
        Bold(true).
        Padding(1, 2)
    
    errorStyle = lipgloss.NewStyle().
        Foreground(lipgloss.Color("#ff0080")).
        Bold(true)
)

func NewMachinaCommand() *cobra.Command {
    cmd := &cobra.Command{
        Use:   "machina",
        Short: "MCP registry and server management",
        Long:  titleStyle.Render("🤖 Machina - MCP Registry Management"),
    }

    // Add subcommands
    cmd.AddCommand(
        newRegistryCommand(),
        newServersCommand(), 
        newHealthCommand(),
        newTUICommand(),
    )

    return cmd
}
```

### Phase 2: Registry Management Commands
```go
// cli/internal/commands/machina/registry.go
package machina

import (
    "context"
    "fmt"
    "github.com/spf13/cobra"
    "github.com/charmbracelet/log"
    "devqai/cli/pkg/client"
)

func newRegistryCommand() *cobra.Command {
    cmd := &cobra.Command{
        Use:   "registry",
        Short: "MCP registry operations",
    }

    cmd.AddCommand(
        &cobra.Command{
            Use:   "status",
            Short: "Show registry status",
            RunE:  registryStatusCommand,
        },
        &cobra.Command{
            Use:   "list",
            Short: "List all registered MCP servers", 
            RunE:  registryListCommand,
        },
    )

    return cmd
}

func registryStatusCommand(cmd *cobra.Command, args []string) error {
    with logfire.span("cli-registry-status"):
        client := client.NewMachinaClient()
        
        status, err := client.GetRegistryStatus(context.Background())
        if err != nil {
            log.Error("Failed to get registry status", "error", err)
            return err
        }

        // Render status with Lip Gloss styling
        fmt.Println(titleStyle.Render("Registry Status"))
        fmt.Printf("Health: %s\n", getHealthStyle(status.Health).Render(status.Health))
        fmt.Printf("Services: %d\n", status.ServiceCount)
        fmt.Printf("Healthy Services: %d\n", status.HealthyCount)
        
        return nil
}

func getHealthStyle(health string) lipgloss.Style {
    switch health {
    case "healthy":
        return lipgloss.NewStyle().Foreground(lipgloss.Color("#00ff00"))
    case "unhealthy":
        return lipgloss.NewStyle().Foreground(lipgloss.Color("#ff0080"))
    default:
        return lipgloss.NewStyle().Foreground(lipgloss.Color("#ffff00"))
    }
}
```

### Phase 3: Interactive TUI Application
```go
// cli/internal/ui/tui/machina_tui.go
package tui

import (
    "context"
    "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/bubbles/list"
    "github.com/charmbracelet/lipgloss"
    "devqai/cli/pkg/client"
)

type MachinaTUI struct {
    list         list.Model
    client       *client.MachinaClient
    services     []ServiceItem
    width        int
    height       int
}

type ServiceItem struct {
    Name        string
    Status      string  
    BuildType   string
    Description string
}

func (s ServiceItem) FilterValue() string { return s.Name }
func (s ServiceItem) Title() string       { return s.Name }
func (s ServiceItem) Description() string { 
    status := getStatusIndicator(s.Status)
    return fmt.Sprintf("%s %s - %s", status, s.BuildType, s.Description)
}

func getStatusIndicator(status string) string {
    switch status {
    case "healthy":
        return "🟢"
    case "unhealthy": 
        return "🔴"
    default:
        return "🟡"
    }
}

func NewMachinaTUI() *MachinaTUI {
    items := []list.Item{}
    
    l := list.New(items, list.NewDefaultDelegate(), 0, 0)
    l.Title = "MCP Servers"
    l.SetShowStatusBar(false)
    l.SetFilteringEnabled(true)
    
    return &MachinaTUI{
        list:   l,
        client: client.NewMachinaClient(),
    }
}

func (m *MachinaTUI) Init() tea.Cmd {
    return m.fetchServices()
}

func (m *MachinaTUI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.WindowSizeMsg:
        m.width = msg.Width
        m.height = msg.Height
        m.list.SetSize(msg.Width-2, msg.Height-2)
        
    case tea.KeyMsg:
        switch msg.String() {
        case "ctrl+c", "q":
            return m, tea.Quit
        case "r":
            return m, m.fetchServices()
        case "enter":
            selected := m.list.SelectedItem()
            if item, ok := selected.(ServiceItem); ok {
                return m, m.showServiceDetails(item)
            }
        }
    
    case servicesMsg:
        m.services = msg.services
        items := make([]list.Item, len(m.services))
        for i, service := range m.services {
            items[i] = ServiceItem(service)
        }
        return m, m.list.SetItems(items)
    }

    var cmd tea.Cmd
    m.list, cmd = m.list.Update(msg)
    return m, cmd
}

func (m *MachinaTUI) View() string {
    if m.width == 0 {
        return "Loading..."
    }
    
    title := lipgloss.NewStyle().
        Foreground(lipgloss.Color("#00ffff")).
        Bold(true).
        Render("🤖 Machina MCP Registry")
    
    help := lipgloss.NewStyle().
        Foreground(lipgloss.Color("#666666")).
        Render("Press 'r' to refresh, 'enter' to view details, 'q' to quit")
    
    return lipgloss.JoinVertical(lipgloss.Left,
        title,
        "",
        m.list.View(),
        "",
        help,
    )
}
```

### Phase 4: Gum Integration for Scripts
```go
// cli/internal/commands/machina/interactive.go
package machina

import (
    "fmt"
    "os/exec"
    "github.com/spf13/cobra"
    "github.com/charmbracelet/log"
)

func newInteractiveCommand() *cobra.Command {
    return &cobra.Command{
        Use:   "interactive",
        Short: "Interactive MCP server management",
        RunE:  interactiveCommand,
    }
}

func interactiveCommand(cmd *cobra.Command, args []string) error {
    // Use Gum for interactive server selection
    serverCmd := exec.Command("gum", "choose", 
        "Start MCP Server",
        "Stop MCP Server", 
        "Restart MCP Server",
        "View Logs",
        "Health Check",
    )
    
    output, err := serverCmd.Output()
    if err != nil {
        return err
    }
    
    action := string(output)
    
    switch action {
    case "Start MCP Server":
        return startServerInteractive()
    case "Stop MCP Server":
        return stopServerInteractive()
    // ... other actions
    }
    
    return nil
}

func startServerInteractive() error {
    // Get list of available servers
    servers := []string{
        "context7-mcp",
        "github-mcp", 
        "docker-mcp",
        "fastapi-mcp",
        // ... all 13 servers
    }
    
    // Use Gum to select server
    cmd := exec.Command("gum", "choose")
    cmd.Args = append(cmd.Args, servers...)
    
    output, err := cmd.Output()
    if err != nil {
        return err
    }
    
    selectedServer := string(output)
    
    // Confirm action with Gum
    confirmCmd := exec.Command("gum", "confirm", 
        fmt.Sprintf("Start %s?", selectedServer))
    
    if err := confirmCmd.Run(); err != nil {
        return fmt.Errorf("operation cancelled")
    }
    
    // Start the server
    log.Info("Starting MCP server", "server", selectedServer)
    return startMCPServer(selectedServer)
}
```

### Phase 5: Charm Log Integration
```go
// cli/internal/ui/logging.go
package ui

import (
    "os"
    "time"
    "github.com/charmbracelet/log"
    "github.com/charmbracelet/lipgloss"
)

func SetupCLILogging() {
    logger := log.NewWithOptions(os.Stderr, log.Options{
        ReportCaller:    true,
        ReportTimestamp: true,  
        TimeFormat:      time.Kitchen,
        Prefix:          "machina",
    })
    
    // Custom styles matching Cyber theme
    styles := log.DefaultStyles()
    styles.Levels[log.InfoLevel] = lipgloss.NewStyle().
        SetString("INFO").
        Foreground(lipgloss.Color("#00ffff")).
        Bold(true)
    
    styles.Levels[log.ErrorLevel] = lipgloss.NewStyle().
        SetString("ERROR").
        Foreground(lipgloss.Color("#ff0080")).
        Bold(true)
    
    styles.Levels[log.WarnLevel] = lipgloss.NewStyle().
        SetString("WARN").
        Foreground(lipgloss.Color("#ffff00")).
        Bold(true)
    
    logger.SetStyles(styles)
    log.SetDefault(logger)
}
```

## Validation Loop

### Level 1: Go Build & Style
```bash
cd cli/
go mod tidy
go build ./cmd/devqai
gofmt -w .
go vet ./...
```

### Level 2: CLI Command Tests
```bash
# Test machina subcommands
./devqai machina registry status
./devqai machina servers list
./devqai machina health check
```

### Level 3: TUI Testing  
```bash
# Test interactive TUI
./devqai machina tui
# Should display server list with real-time status updates
```

### Level 4: Integration Testing
```bash
# Test with real registry service
go test ./internal/commands/machina/... -v
# Must connect to actual registry API
```

### Level 5: Authentication Testing
```bash
# Test trustlis integration
./devqai auth login
./devqai machina registry status
# Should authenticate via trustlis before API calls
```

**Confidence Score: 8/10** - High confidence due to well-established Charm.sh patterns and clear CLI structure.
```

These PRPs provide comprehensive implementation plans for all three Machina features, following your requirements for FastAPI, FastMCP, Logfire observability, real testing (no mocks), and integration with the DevQ.ai ecosystem. Each PRP includes detailed technical blueprints, validation loops, and confidence scoring for successful one-pass implementation.
