package main

import (
    "fmt"
    "time"
)

type Engine struct {
    playbook   *Playbook
    conditions map[string]bool
}

func newEngine(playbook *Playbook) *Engine {
    return &Engine{
        playbook:   playbook,
        conditions: make(map[string]bool),
    }
}

func (e *Engine) executeStep(branchIdx, stepIdx int) error {
    step := &e.playbook.Branches[branchIdx].Steps[stepIdx]
    
    // Check if condition is met
    if !e.isConditionMet(step.Condition) {
        return fmt.Errorf("condition not met: %s", step.Condition)
    }

    step.Status = StatusRunning
    
    // Simulate work
    time.Sleep(time.Millisecond * 500)
    
    // Mark step as completed and set new conditions
    step.Status = StatusCompleted
    e.setCondition(e.getCompletionCondition(step), true)
    
    return nil
}

func (e *Engine) isConditionMet(condition string) bool {
    if condition == "start" {
        return true
    }
    return e.conditions[condition]
}

func (e *Engine) setCondition(condition string, value bool) {
    e.conditions[condition] = value
}

func (e *Engine) getCompletionCondition(step *Step) string {
    switch step.Agent {
    case "data-collector":
        return "data-collected"
    case "data-validator":
        return "data-validated"
    case "email-sender":
        return "email-sent"
    case "tracker":
        return "tracking-active"
    case "reporter":
        return "report-generated"
    default:
        return fmt.Sprintf("%s-completed", step.Agent)
    }
}

func (e *Engine) getNextExecutableSteps() []StepLocation {
    var steps []StepLocation
    
    for branchIdx, branch := range e.playbook.Branches {
        for stepIdx, step := range branch.Steps {
            if step.Status == StatusPending && e.isConditionMet(step.Condition) {
                steps = append(steps, StepLocation{
                    BranchIndex: branchIdx,
                    StepIndex:   stepIdx,
                })
            }
        }
    }
    
    return steps
}

func (e *Engine) isPlaybookComplete() bool {
    for _, branch := range e.playbook.Branches {
        for _, step := range branch.Steps {
            if step.Status != StatusCompleted {
                return false
            }
        }
    }
    return true
}

type StepLocation struct {
    BranchIndex int
    StepIndex   int
}