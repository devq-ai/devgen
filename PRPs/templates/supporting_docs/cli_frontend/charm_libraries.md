## Core Framework/Libraries

+ Bubble Tea. GitHub - charmbracelet/gum: A tool for glamorous shell scripts 🎀 - TUI framework based on Elm architecture: github.com/charmbracelet/bubbletea
+ Lip Gloss. GitHub - charmbracelet/bubbletea: A powerful little TUI framework 🏗 - CSS-like styling for terminal layouts: github.com/charmbracelet/lipgloss
+ Bubbles. charmbracelet repositories · GitHub - Component library for Bubble Tea (inputs, lists, progress bars, etc.): github.com/charmbracelet/bubbles
+ Glamour. Building an Awesome Terminal User Interface Using Go, Bubble Tea, and Lip Gloss | Grootan Technologies - Stylesheet-driven markdown renderer
+ Wish. GitHub - charmbracelet/lipgloss: Style definitions for nice terminal layouts 👄 - SSH server framework with middleware
+ Harmonica. Building an Awesome Terminal User Interface Using Go, Bubble Tea, and Lip Gloss | Grootan Technologies - Physics-based animation toolkit
+ Log. Minimal, colorful logging library: github.com/charmbracelet/log
+ Huh. Terminal forms and prompts

### Available Components:

+ Text Input Enhancing Shell Scripts with Charmbracelet Gum: A Practical Guide | by Iamjignyasa | Medium: Single-line input with unicode support, pasting, scrolling, password masking, autocompletion, and validation
+ Text Area How Do I Use Charmbracelet Gum to Improve My Scripts: Multi-line input with vertical scrolling and customization options
+ List Non-interactive value to gum commands · Issue #788 · charmbracelet/gum: Feature-rich list browser with filtering, pagination, help, status messages, and spinner
+ Table How Do I Use Charmbracelet Gum to Improve My Scripts: Tabular data display with vertical scrolling and customization
+ Progress Bar GitHub - charmbracelet/log: A minimal, colorful Go logging library 🪵: Customizable progress meter with optional spring animations via Harmonica
+ Spinner How Do I Use Charmbracelet Gum to Improve My Scripts: Loading indicators with multiple built-in animations
+ Viewport: Scrollable content areas
+ Paginator: Page navigation controls
+ Help: Built-in help system
+ File Picker: File browser component
+ Stopwatch/Timer: Time tracking components

```
// Text input with validation
input := textinput.New()
input.Placeholder = "Enter email..."
input.CharLimit = 50
input.Focus()

// List with filtering
items := []list.Item{...}
delegate := list.NewDefaultDelegate()
list := list.New(items, delegate, 80, 40)
list.Title = "Choose an option"
```

### Log - Logging
+ Minimal, colorful Go logging library with leveled structured logging and small API.
+ Key Features:
	+ Leveled Logging: Debug, Info, Warn, Error, Fatal levels
	+ Structured Logging: Key-value pairs for context
	+ Multiple Formatters: Text (styled with Lip Gloss), JSON, and Logfmt
	+ Customizable Styles: Override colors and formatting per level/key
	+ Standard Log Adapter: Drop-in replacement for stdlib log
	+ Slog Handler: Compatible with Go's structured logging
	+ Sub-loggers: Create loggers with specific fields and prefixes
+ References:
	+ Input log package - github.com/charmbracelet/log - Go Packages: gum input --placeholder "Enter name" --password
	+ Write log package - github.com/charmbracelet/log - Go Packages: gum write --placeholder "Enter details" --width 80
	+ Choose log package - github.com/charmbracelet/log - Go Packages: gum choose "Option A" "Option B" --limit 2
	+ Filter log package - github.com/charmbracelet/log - Go Packages: cat list.txt | gum filter --limit 3 --no-limit
	+ Confirm log package - github.com/charmbracelet/log - Go Packages: gum confirm "Delete file?" && rm file.txt
	+ Spin log package - github.com/charmbracelet/log - Go Packages: gum spin --title "Loading..." -- sleep 5
	+ Style log package - github.com/charmbracelet/log - Go Packages: gum style --foreground 212 --border double "Styled text"
	+ Join log package - github.com/charmbracelet/log - Go Packages: Combine text horizontally or vertically
	+ File log package - github.com/charmbracelet/log - Go Packages: File picker interface
	+ Pager log package - github.com/charmbracelet/log - Go Packages: Scroll through content with line numbers
	+ Table log package - github.com/charmbracelet/log - Go Packages: Select from tabular data
	+ Log log package - github.com/charmbracelet/log - Go Packages: Structured logging with different levels and styling log package - github.com/charmbracelet/log - Go Packages

```
#!/bin/bash
TYPE=$(gum choose "fix" "feat" "docs" "style")
SCOPE=$(gum input --placeholder "scope")
SUMMARY=$(gum input --value "$TYPE($SCOPE): " --placeholder "Summary")
DESCRIPTION=$(gum write --placeholder "Details")

gum confirm "Commit changes?" && \
  git commit -m "$SUMMARY" -m "$DESCRIPTION"
```

```
// Basic logging
log.Info("Server started", "port", 8080, "env", "production")
log.Error("Database connection failed", "err", err, "retries", 3)

// Custom styling
styles := log.DefaultStyles()
styles.Levels[log.ErrorLevel] = lipgloss.NewStyle().
    SetString("CRITICAL").
    Background(lipgloss.Color("196"))
logger := log.New(os.Stderr)
logger.SetStyles(styles)

// Sub-logger
dbLogger := logger.With("component", "database", "version", "v2.1")
dbLogger.Info("Connection established")

// Multiple formatters
logger.SetFormatter(log.JSONFormatter)  // For production
logger.SetFormatter(log.TextFormatter)  // For development
```

