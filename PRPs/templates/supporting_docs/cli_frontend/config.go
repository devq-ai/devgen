package main

import (
    "fmt"
    "os"

    "gopkg.in/yaml.v3"
)

type Config struct {
    Playbook Playbook `yaml:"playbook"`
}

type Playbook struct {
    Name     string   `yaml:"name"`
    Branches []Branch `yaml:"branches"`
}

type Branch struct {
    Name  string `yaml:"name"`
    Steps []Step `yaml:"steps"`
}

type Step struct {
    Agent     string `yaml:"agent"`
    Condition string `yaml:"condition"`
    Action    string `yaml:"action"`
    Status    StepStatus
}

type StepStatus int

const (
    StatusPending StepStatus = iota
    StatusRunning
    StatusCompleted
    StatusFailed
)

func (s StepStatus) String() string {
    switch s {
    case StatusPending:
        return "pending"
    case StatusRunning:
        return "running"
    case StatusCompleted:
        return "completed"
    case StatusFailed:
        return "failed"
    default:
        return "unknown"
    }
}

func loadPlaybook(filename string) (*Playbook, error) {
    data, err := os.ReadFile(filename)
    if err != nil {
        return nil, fmt.Errorf("reading config file: %w", err)
    }

    var config Config
    if err := yaml.Unmarshal(data, &config); err != nil {
        return nil, fmt.Errorf("parsing YAML: %w", err)
    }

    // Initialize all steps as pending
    for i := range config.Playbook.Branches {
        for j := range config.Playbook.Branches[i].Steps {
            config.Playbook.Branches[i].Steps[j].Status = StatusPending
        }
    }

    return &config.Playbook, nil
}