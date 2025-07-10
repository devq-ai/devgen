package main

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
)

func main() {
    var configFile string

    rootCmd := &cobra.Command{
        Use:   "agentic",
        Short: "Execute agentic playbooks",
        RunE: func(cmd *cobra.Command, args []string) error {
            return runPlaybook(configFile)
        },
    }

    rootCmd.Flags().StringVarP(&configFile, "config", "c", "playbook.yaml", "Path to playbook config")

    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        os.Exit(1)
    }
}

func runPlaybook(configFile string) error {
    playbook, err := loadPlaybook(configFile)
    if err != nil {
        return fmt.Errorf("failed to load playbook: %w", err)
    }

    ui := newUI(playbook)
    return ui.run()
}