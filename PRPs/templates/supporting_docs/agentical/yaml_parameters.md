## Workflow Architecture Lexicon v2.0
### Mathematical Formalization with N,K Coefficients
SEQUENTIAL: 1 route with N steps where N ≥ 1
MAPPED:     1→K routes with Nₖ steps per route where K ≥ 1, Nₖ ≥ 1
### WORKFLOW DEFINITION
pattern: {SEQUENTIAL, MAPPED}                    # |patterns| = 2
condition: Condition ∪ {null}                    # Optional global condition
routes: [Route₁, Route₂, ..., Routeₖ]          # |routes| = K for MAPPED
steps: [Step₁, Step₂, ..., Stepₙ]              # |steps| = N for SEQUENTIAL

### ROUTE DEFINITION (MAPPED: K routes where K ≥ 1)
Route_i:
  id: string_i                                   # Unique identifier ∀i ∈ [1,K]
  steps: [Step₁, Step₂, ..., Step_Nᵢ]           # Nᵢ steps in route i
  parallel: {true, false}                       # Boolean parallel execution
  spawns: [route_j₁, route_j₂, ..., route_jₘ]   # Spawnable routes where jₘ ∈ [1,K]
  condition: Condition ∪ {null}                 # Optional route condition

### STEP DEFINITION (N steps total across all routes)
Step_i:
  id: string_i                                  # Unique identifier ∀i ∈ [1,N]
  description: string_i                         # Task description
  agent_tasks: [task₁, task₂, ..., task_Tᵢ]     # Tᵢ tasks for step i
  required: {true, false}                       # Boolean requirement
  prior_step: StepID ∪ [StepID₁, ..., StepIDₚ] ∪ {null}  # P dependencies
  next_step: FlowControl ∪ {null}               # Single next step
  if_success: FlowControl ∪ {null}              # Success branch
  if_fail: FlowControl ∪ {null}                 # Failure branch  
  condition: Condition ∪ {null}                 # Optional step condition
  timeout: ℕ ∪ {null}                          # Seconds ∈ Natural numbers
  retry_limit: ℕ ∪ {null}                      # Attempts ∈ Natural numbers

# CONDITION TYPES (|ConditionTypes| = 3)
ObjectiveCondition:
  type: "OBJECTIVE"
  expression: BooleanExpression                 # f: Variables → {true, false}
  variables: {var₁: value₁, ..., var_Vᵢ: value_Vᵢ}

FeedbackCondition:  
  type: "FEEDBACK"
  evaluator: {"same-agent", "other-agent", "human"}  # |evaluators| = 3
  prompt: string

ConditionExpression:
  type: "EXPRESSION" 
  formula: BitwiseFormula                       # Composition of conditions
  conditions: {A: Condition_A, ..., Z: Condition_Z}  # Named conditions

FlowControl ∈ {
  "step_id",           # Direct step reference
  "END",               # Workflow termination  
  "RETRY",             # Current step retry
  "SKIP",              # Skip to next logical step
  "ROUTE:route_id",    # Route jump (MAPPED only)
  "SPAWN:route_id"     # Parallel route spawn (MAPPED only)
}

### Bitwise Operators (Mathematical Set)
BitwiseOps = {"&", "|", "^", "~", "(", ")"}

### Expression Variables (Constrained Domain)
Variables ∈ {
  Counters: {count, total, percentage, ratio, time_elapsed, attempts, errors},
  Flags: {X_created, X_configured, X_tested, X_approved, X_deployed} 
    where X ∈ ResourceTypes
}

### Mathematical: 1 route, N=4 steps, linear execution
objective: "Data processing pipeline with conditional branching"

workflow:
  pattern: SEQUENTIAL  # N=4 sequential steps
  steps: [extract_data, validate_data, data_cleanup, process_data]

steps:
  - id: "extract_data"                          # Step₁
    description: "Extract data from source"
    agent_tasks: [task₁, task₂, task₃]          # T₁=3 tasks
    required: true
    next_step: "validate_data"                  # Step₁ → Step₂
    if_fail: "error_handling"                   # Step₁ → Stepₑᵣᵣₒᵣ
    condition:
      type: OBJECTIVE
      expression: "records_extracted > 0"       # f₁: ℕ → {true, false}
    timeout: 300                                # t₁ = 300 seconds

  - id: "validate_data"                         # Step₂  
    description: "Validate extracted data quality"
    agent_tasks: [task₁, task₂, task₃]          # T₂=3 tasks
    required: true
    prior_step: "extract_data"                  # Step₁ → Step₂ dependency
    if_success: "process_data"                  # Step₂ → Step₄ (skip cleanup)
    if_fail: "data_cleanup"                     # Step₂ → Step₃ (needs cleanup)
    condition:
      type: OBJECTIVE  
      expression: "error_rate < 0.05"           # f₂: ℝ → {true, false}

  - id: "data_cleanup"                          # Step₃ (optional)
    description: "Clean problematic data"
    agent_tasks: [task₁, task₂, task₃]          # T₃=3 tasks
    required: false                             # Optional step
    prior_step: "validate_data"                 # Step₂ → Step₃ dependency  
    next_step: "validate_data"                  # Step₃ → Step₂ (re-validate)
    retry_limit: 2                              # Max 2 retry attempts

  - id: "process_data"                          # Step₄
    description: "Transform and enrich data" 
    agent_tasks: [task₁, task₂, task₃]          # T₄=3 tasks
    required: true
    prior_step: ["validate_data"]               # Step₂ → Step₄ dependency
    next_step: "END"                            # Workflow termination

### Mathematical: K=3 routes, N₁=2, N₂=2, N₃=3 steps respectively
objective: "Multi-channel content pipeline with parallel execution"

workflow:
  pattern: MAPPED  # K=3 parallel routes
  routes: [content_creation, content_review, distribution]

routes:
  - id: "content_creation"                      # Route₁: N₁=2 steps
    parallel: false                             # Sequential execution within route
    steps: [research_topics, create_content]
    spawns: ["content_review"]                  # Route₁ → Route₂ spawn

    steps:
      - id: "research_topics"                   # Route₁.Step₁
        description: "Research content topics and trends"
        agent_tasks: [task₁, task₂, task₃]      # T₁₁=3 tasks
        required: true
        next_step: "create_content"             # Route₁.Step₁ → Route₁.Step₂
        
      - id: "create_content"                    # Route₁.Step₂ 
        description: "Generate content across formats"
        agent_tasks: [task₁, task₂, task₃]      # T₁₂=3 tasks
        required: true
        prior_step: "research_topics"           # Sequential dependency
        if_success: "SPAWN:content_review"      # Route₁ → Route₂ spawn
        condition:
          type: FEEDBACK
          evaluator: "same-agent"
          prompt: "Is content quality sufficient for publication?"

  - id: "content_review"                        # Route₂: N₂=2 steps
    parallel: true                              # Can run parallel with Route₃
    spawns: ["distribution"]                    # Route₂ → Route₃ spawn
    condition:
      type: OBJECTIVE
      expression: "content_created == true"     # Route execution condition

    steps:
      - id: "quality_review"                    # Route₂.Step₁
        description: "Review content quality and compliance"  
        agent_tasks: [task₁, task₂, task₃]      # T₂₁=3 tasks
        required: true
        if_success: "approve_content"           # Route₂.Step₁ → Route₂.Step₂
        if_fail: "ROUTE:content_creation"       # Route₂ → Route₁ jump

      - id: "approve_content"                   # Route₂.Step₂
        description: "Final approval for publication"
        agent_tasks: [task₁, task₂, task₃]      # T₂₂=3 tasks  
        required: true
        prior_step: "quality_review"            # Sequential dependency
        if_success: "SPAWN:distribution"        # Route₂ → Route₃ spawn
        condition:
          type: FEEDBACK
          evaluator: "human"                    # Human-in-the-loop approval
          prompt: "Approve this content for publication?"

  - id: "distribution"                          # Route₃: N₃=3 steps
    parallel: true                              # Parallel execution
    condition:
      type: OBJECTIVE  
      expression: "content_approved == true"    # Route execution condition

    steps:
      - id: "seo_optimization"                  # Route₃.Step₁
        description: "Optimize content for search engines"
        agent_tasks: [task₁, task₂, task₃]      # T₃₁=3 tasks
        required: true
        next_step: "schedule_posts"             # Route₃.Step₁ → Route₃.Step₂

      - id: "schedule_posts"                    # Route₃.Step₂
        description: "Schedule social media distribution"  
        agent_tasks: [task₁, task₂, task₃]      # T₃₂=3 tasks
        required: true
        prior_step: "seo_optimization"          # Sequential dependency
        next_step: "monitor_engagement"         # Route₃.Step₂ → Route₃.Step₃

      - id: "monitor_engagement"                # Route₃.Step₃
        description: "Monitor and respond to engagement"
        agent_tasks: [task₁, task₂, task₃]      # T₃₃=3 tasks
        required: false                         # Optional monitoring
        prior_step: "schedule_posts"            # Sequential dependency  
        timeout: 86400                          # t₃₃ = 86400 seconds (24h)
        next_step: "END"                        # Workflow termination

workflow:
  pattern: SEQUENTIAL
  condition:
    type: EXPRESSION
    formula: "(A & B) | (C & ~D)"              # Bitwise composition
    conditions:                                 # |conditions| = 4
      A: # Condition₁
        type: OBJECTIVE
        expression: "success_rate > 0.8"        # f_A: ℝ → {true, false}
      B: # Condition₂  
        type: FEEDBACK
        evaluator: "human"
        prompt: "Quality meets standards?"      # f_B: Human → {true, false}
      C: # Condition₃
        type: OBJECTIVE
        expression: "attempts < 3"              # f_C: ℕ → {true, false}  
      D: # Condition₄
        type: OBJECTIVE
        expression: "critical_errors > 0"       # f_D: ℕ → {true, false}

## Mathematical evaluation: 
### Result = (f_A ∧ f_B) ∨ (f_C ∧ ¬f_D)

Total Workflow Complexity: O(K × max(Nₖ) × C)
where:
  K = number of routes (1 for SEQUENTIAL, ≥1 for MAPPED)  
  Nₖ = steps per route k
  C = conditions per step (0 ≤ C ≤ 3)

Space Complexity: O(Σ Nₖ + Σ Tᵢ)
where:
  Σ Nₖ = total steps across all routes
  Σ Tᵢ = total tasks across all steps

Execution Patterns:
  SEQUENTIAL: Serial execution, O(N) time
  MAPPED: Parallel potential, O(max(Nₖ)) time with K workers

### WORKFLOW DEFINITION
pattern: {SEQUENTIAL, MAPPED}                    # |patterns| = 2
condition: Condition ∪ {null}                    # Optional global condition
routes: [Route₁, Route₂, ..., Routeₖ]          # |routes| = K for MAPPED
steps: [Step₁, Step₂, ..., Stepₙ]              # |steps| = N for SEQUENTIAL

### ROUTE DEFINITION (MAPPED: K routes where K ≥ 1)
Route_i:
  id: string_i                                   # Unique identifier ∀i ∈ [1,K]
  steps: [Step₁, Step₂, ..., Step_Nᵢ]           # Nᵢ steps in route i
  parallel: {true, false}                       # Boolean parallel execution
  spawns: [route_j₁, route_j₂, ..., route_jₘ]   # Spawnable routes where jₘ ∈ [1,K]
  condition: Condition ∪ {null}                 # Optional route condition

### STEP DEFINITION (N steps total across all routes)
Step_i:
  id: string_i                                  # Unique identifier ∀i ∈ [1,N]
  description: string_i                         # Task description
  agent_tasks: [task₁, task₂, ..., task_Tᵢ]     # Tᵢ tasks for step i
  required: {true, false}                       # Boolean requirement
  prior_step: StepID ∪ [StepID₁, ..., StepIDₚ] ∪ {null}  # P dependencies
  next_step: FlowControl ∪ {null}               # Single next step
  if_success: FlowControl ∪ {null}              # Success branch
  if_fail: FlowControl ∪ {null}                 # Failure branch  
  condition: Condition ∪ {null}                 # Optional step condition
  timeout: ℕ ∪ {null}                          # Seconds ∈ Natural numbers
  retry_limit: ℕ ∪ {null}                      # Attempts ∈ Natural numbers

# CONDITION TYPES (|ConditionTypes| = 3)
ObjectiveCondition:
  type: "OBJECTIVE"
  expression: BooleanExpression                 # f: Variables → {true, false}
  variables: {var₁: value₁, ..., var_Vᵢ: value_Vᵢ}

FeedbackCondition:  
  type: "FEEDBACK"
  evaluator: {"same-agent", "other-agent", "human"}  # |evaluators| = 3
  prompt: string

ConditionExpression:
  type: "EXPRESSION" 
  formula: LogicalFormula                       # Composition of conditions
  conditions: {A: Condition_A, ..., Z: Condition_Z}  # Named conditions

FlowControl ∈ {
  "step_id",           # Direct step reference
  "END",               # Workflow termination  
  "RETRY",             # Current step retry
  "SKIP",              # Skip to next logical step
  "ROUTE:route_id",    # Route jump (MAPPED only)
  "SPAWN:route_id"     # Parallel route spawn (MAPPED only)
}

### Logical Operators (Mathematical Set)
LogicalOps = {"&", "|", "^", "~", "(", ")"}
where:
  & = AND    (logical conjunction)
  | = OR     (logical disjunction) 
  ^ = XOR    (logical exclusive or)
  ~ = NOT    (logical negation)
  () = Grouping (precedence control)

### Expression Variables (Constrained Domain)
Variables ∈ {
  Counters: {count, total, percentage, ratio, time_elapsed, attempts, errors},
  Flags: {X_created, X_configured, X_tested, X_approved, X_deployed} 
    where X ∈ ResourceTypes
}

### Mathematical: 1 route, N=4 steps, linear execution
objective: "Data processing pipeline with conditional branching"

workflow:
  pattern: SEQUENTIAL  # N=4 sequential steps
  steps: [extract_data, validate_data, data_cleanup, process_data]

steps:
  - id: "extract_data"                          # Step₁
    description: "Extract data from source"
    agent_tasks: [task₁, task₂, task₃]          # T₁=3 tasks
    required: true
    next_step: "validate_data"                  # Step₁ → Step₂
    if_fail: "error_handling"                   # Step₁ → Stepₑᵣᵣₒᵣ
    condition:
      type: OBJECTIVE
      expression: "records_extracted > 0"       # f₁: ℕ → {true, false}
    timeout: 300                                # t₁ = 300 seconds

  - id: "validate_data"                         # Step₂  
    description: "Validate extracted data quality"
    agent_tasks: [task₁, task₂, task₃]          # T₂=3 tasks
    required: true
    prior_step: "extract_data"                  # Step₁ → Step₂ dependency
    if_success: "process_data"                  # Step₂ → Step₄ (skip cleanup)
    if_fail: "data_cleanup"                     # Step₂ → Step₃ (needs cleanup)
    condition:
      type: OBJECTIVE  
      expression: "error_rate < 0.05"           # f₂: ℝ → {true, false}

  - id: "data_cleanup"                          # Step₃ (optional)
    description: "Clean problematic data"
    agent_tasks: [task₁, task₂, task₃]          # T₃=3 tasks
    required: false                             # Optional step
    prior_step: "validate_data"                 # Step₂ → Step₃ dependency  
    next_step: "validate_data"                  # Step₃ → Step₂ (re-validate)
    retry_limit: 2                              # Max 2 retry attempts

  - id: "process_data"                          # Step₄
    description: "Transform and enrich data" 
    agent_tasks: [task₁, task₂, task₃]          # T₄=3 tasks
    required: true
    prior_step: ["validate_data"]               # Step₂ → Step₄ dependency
    next_step: "END"                            # Workflow termination

### Mathematical: K=3 routes, N₁=2, N₂=2, N₃=3 steps respectively
objective: "Multi-channel content pipeline with parallel execution"

workflow:
  pattern: MAPPED  # K=3 parallel routes
  routes: [content_creation, content_review, distribution]

routes:
  - id: "content_creation"                      # Route₁: N₁=2 steps
    parallel: false                             # Sequential execution within route
    steps: [research_topics, create_content]
    spawns: ["content_review"]                  # Route₁ → Route₂ spawn

    steps:
      - id: "research_topics"                   # Route₁.Step₁
        description: "Research content topics and trends"
        agent_tasks: [task₁, task₂, task₃]      # T₁₁=3 tasks
        required: true
        next_step: "create_content"             # Route₁.Step₁ → Route₁.Step₂
        
      - id: "create_content"                    # Route₁.Step₂ 
        description: "Generate content across formats"
        agent_tasks: [task₁, task₂, task₃]      # T₁₂=3 tasks
        required: true
        prior_step: "research_topics"           # Sequential dependency
        if_success: "SPAWN:content_review"      # Route₁ → Route₂ spawn
        condition:
          type: FEEDBACK
          evaluator: "same-agent"
          prompt: "Is content quality sufficient for publication?"

  - id: "content_review"                        # Route₂: N₂=2 steps
    parallel: true                              # Can run parallel with Route₃
    spawns: ["distribution"]                    # Route₂ → Route₃ spawn
    condition:
      type: OBJECTIVE
      expression: "content_created == true"     # Route execution condition

    steps:
      - id: "quality_review"                    # Route₂.Step₁
        description: "Review content quality and compliance"  
        agent_tasks: [task₁, task₂, task₃]      # T₂₁=3 tasks
        required: true
        if_success: "approve_content"           # Route₂.Step₁ → Route₂.Step₂
        if_fail: "ROUTE:content_creation"       # Route₂ → Route₁ jump

      - id: "approve_content"                   # Route₂.Step₂
        description: "Final approval for publication"
        agent_tasks: [task₁, task₂, task₃]      # T₂₂=3 tasks  
        required: true
        prior_step: "quality_review"            # Sequential dependency
        if_success: "SPAWN:distribution"        # Route₂ → Route₃ spawn
        condition:
          type: FEEDBACK
          evaluator: "human"                    # Human-in-the-loop approval
          prompt: "Approve this content for publication?"

  - id: "distribution"                          # Route₃: N₃=3 steps
    parallel: true                              # Parallel execution
    condition:
      type: OBJECTIVE  
      expression: "content_approved == true"    # Route execution condition

    steps:
      - id: "seo_optimization"                  # Route₃.Step₁
        description: "Optimize content for search engines"
        agent_tasks: [task₁, task₂, task₃]      # T₃₁=3 tasks
        required: true
        next_step: "schedule_posts"             # Route₃.Step₁ → Route₃.Step₂

      - id: "schedule_posts"                    # Route₃.Step₂
        description: "Schedule social media distribution"  
        agent_tasks: [task₁, task₂, task₃]      # T₃₂=3 tasks
        required: true
        prior_step: "seo_optimization"          # Sequential dependency
        next_step: "monitor_engagement"         # Route₃.Step₂ → Route₃.Step₃

      - id: "monitor_engagement"                # Route₃.Step₃
        description: "Monitor and respond to engagement"
        agent_tasks: [task₁, task₂, task₃]      # T₃₃=3 tasks
        required: false                         # Optional monitoring
        prior_step: "schedule_posts"            # Sequential dependency  
        timeout: 86400                          # t₃₃ = 86400 seconds (24h)
        next_step: "END"                        # Workflow termination

workflow:
  pattern: SEQUENTIAL
  condition:
    type: EXPRESSION
    formula: "(A & B) | (C & ~D)"              # Logical composition
    conditions:                                 # |conditions| = 4
      A: # Condition₁
        type: OBJECTIVE
        expression: "success_rate > 0.8"        # f_A: ℝ → {true, false}
      B: # Condition₂  
        type: FEEDBACK
        evaluator: "human"
        prompt: "Quality meets standards?"      # f_B: Human → {true, false}
      C: # Condition₃
        type: OBJECTIVE
        expression: "attempts < 3"              # f_C: ℕ → {true, false}  
      D: # Condition₄
        type: OBJECTIVE
        expression: "critical_errors > 0"       # f_D: ℕ → {true, false}

### Mathematical evaluation: 
**Result = (f_A ∧ f_B) ∨ (f_C ∧ ¬f_D)**

Total Workflow Complexity: O(K × max(Nₖ) × C)
where:
  K = number of routes (1 for SEQUENTIAL, ≥1 for MAPPED)  
  Nₖ = steps per route k
  C = conditions per step (0 ≤ C ≤ 3)

Space Complexity: O(Σ Nₖ + Σ Tᵢ)
where:
  Σ Nₖ = total steps across all routes
  Σ Tᵢ = total tasks across all steps

Execution Patterns:
  SEQUENTIAL: Serial execution, O(N) time
  MAPPED: Parallel potential, O(max(Nₖ)) time with K workers

Logical Evaluation:
  LogicalOps operate on boolean conditions
  Truth table evaluation for compound expressions
  Short-circuit evaluation for performance optimization


| parameter:   | datatype: | allowed_values: |
----------------------------------------------
| pattern      | string    |                 |
| num_of_steps | integer   |                 |
| step_number  | integer   |                 |
| step_name    | string    |                 |
| step_desc    | string    |                 |
| input		   | 		   |                 |
| process	   |           |                 |
| condition    |           |                 |
| output	   |           |                 |
| start_order  |           |                 |
|              |           |                 |
|              |           |                 |
|              |           |                 |
|              |           |                 |
|              |           |                 |
|              |           |                 |
|              |           |                 |
|              |           |                 |
|              |           |                 |
|              |           |                 |
----------------------------------------------

---

### logical_operator  
```
and: 1 & 1 = 1
 or: 1 | 0 = 1
xor: 1 ^ 0 = 1
not:   ~ 0 = 1
```

playbook
name
id
version
date_last_run
objective