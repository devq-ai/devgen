# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## 🚨 ABSOLUTE PRIORITY: HONESTY AND TRUTHFUL STATEMENTS ABOVE ALL ELSE

- You are 100% honest in all statements.

### CORE PRINCIPLES (NON-NEGOTIABLE)

1. **COMPLETE HONESTY IN ALL STATEMENTS** - You are 100% honest in all statements. No deception, no fake claims, no false progress reports.

2. **ZERO FAKE/MOCK/HARDCODED DATA** - Functions must return real results or fail appropriately. The presence of such code will immediately fail a subtask (0% completion).

3. **MANDATORY CODE VERIFICATION** - Every subtask the agent must verify the veracity of ALL code and tests will be written not only for the function of the code but that the code is verified by a second source.

4. **TRUTHFUL TEST IMPLEMENTATION** - PyTests cannot be written to pass with mock/fake/hardcoded data. Tests must verify real functionality with deliberate tests of all MCP tools and must return accurate and real results.

## Critical Files
+ .env - contains environment secrets, keys, tokens, passwords and config parameters for critical systems
+ CLAUDE.md - provides guidance to Claude Code (claude.ai/code) when working with code in this repository

## Capitalize on MCP Resources

### REQUIRED MCP SERVERS - MUST BE VERIFIED ON SESSION START

**CRITICAL**: Before beginning ANY work, Claude Code MUST verify connectivity and functionality of ALL MCP servers. These servers provide essential capabilities that are REQUIRED for proper operation.

### Claude Code Tools
- **context7** - Documentation sourcing and semantic search; always use context7 to plan software, never guess the correct syntax, double check your planned code
- **filesystem**: File operations and project management
- **git**: Version control with DevQ.ai team configuration
- **fetch**: API calls and external resource access
- **memory**: Session persistence across Claude Code interactions
- **sequentialthinking**: Step-by-step problem solving

## Project Nature

This is a **PRP (Product Requirement Prompt) Framework** repository, not a traditional software project. The core concept: **"PRP = PRD + curated codebase intelligence + agent/runbook"** - designed to enable AI agents to ship production-ready code on the first pass.

### Core Architecture

#### Command-Driven System

- **28+ pre-configured Claude Code commands** in `.claude/commands/`
- Commands organized by function:
  - `PRPs/` - PRP creation and execution workflows
  - `development/` - Core development utilities (prime-core, onboarding, debug)
  - `code-quality/` - Review and refactoring commands
  - `rapid-development/experimental/` - Parallel PRP creation and hackathon tools
  - `git-operations/` - Conflict resolution and smart git operations

#### Template-Based Methodology

- **PRP Templates** in `PRPs/templates/` follow structured format with validation loops
- **Context-Rich Approach**: Every PRP must include comprehensive documentation, examples, and gotchas
- **Validation-First Design**: Each PRP contains executable validation gates (syntax, tests, integration)

#### AI Documentation Curation

- `PRPs/ai_docs/` contains curated Claude Code documentation for context injection
- `claude_md_files/` provides framework-specific CLAUDE.md examples

### Development Commands

#### Python Package Management

```bash
# This project uses UV package manager
uv venv                    # Create virtual environment
uv sync                    # Install dependencies
uv run [script]            # Run Python scripts
```

#### PRP Execution

```bash
# Interactive mode (recommended for development)
uv run PRPs/scripts/prp_runner.py --prp [prp-name] --interactive

# Headless mode (for CI/CD)
uv run PRPs/scripts/prp_runner.py --prp [prp-name] --output-format json

# Streaming JSON (for real-time monitoring)
uv run PRPs/scripts/prp_runner.py --prp [prp-name] --output-format stream-json
```

#### Key Claude Commands

- `/prp-base-create` - Generate comprehensive PRPs with research
- `/prp-base-execute` - Execute PRPs against codebase
- `/prp-planning-create` - Create planning documents with diagrams
- `/prime-core` - Prime Claude with project context
- `/review-staged-unstaged` - Review git changes using PRP methodology

### Critical Success Patterns

#### The PRP Methodology

1. **Context is King**: Include ALL necessary documentation, examples, and caveats
2. **Validation Loops**: Provide executable tests/lints the AI can run and fix
3. **Information Dense**: Use keywords and patterns from the codebase
4. **Progressive Success**: Start simple, validate, then enhance

#### PRP Structure Requirements

- **Goal**: Specific end state and desires
- **Why**: Business value and user impact
- **What**: User-visible behavior and technical requirements
- **All Needed Context**: Documentation URLs, code examples, gotchas, patterns
- **Implementation Blueprint**: Pseudocode with critical details and task lists
- **Validation Loop**: Executable commands for syntax, tests, integration

#### Validation Gates (Must be Executable)

```bash
# Level 1: Syntax & Style
ruff check --fix && mypy .

# Level 2: Unit Tests
uv run pytest tests/ -v

# Level 3: Integration
curl -X POST http://localhost:8000/endpoint -H "Content-Type: application/json" -d '{...}'
```

### Anti-Patterns to Avoid

- L Don't create minimal context prompts - context is everything
- L Don't skip validation steps - they're critical for one-pass success
- L Don't ignore the structured PRP format - it's battle-tested
- L Don't create new patterns when existing templates work
- L Don't hardcode values that should be config
- L Don't catch all exceptions - be specific

### Working with This Framework

#### When Creating PRPs

1. **Research First**: Analyze codebase patterns and external documentation
2. **Use Templates**: Start with `PRPs/templates/prp_base.md`
3. **Include Context**: Add URLs, code examples, gotchas, and patterns
4. **Validate Early**: Include executable validation commands
5. **Score Confidence**: Rate PRP 1-10 for one-pass implementation success

#### When Executing PRPs

1. **Load PRP**: Read and understand all context and requirements
2. **ULTRATHINK**: Create comprehensive plan, break down into todos
3. **Execute**: Implement following the blueprint
4. **Validate**: Run each validation command, fix failures
5. **Complete**: Ensure all checklist items done

#### Command Usage

- Access via `/` prefix in Claude Code
- Commands are self-documenting with argument placeholders
- Use parallel creation commands for rapid development
- Leverage existing review and refactoring commands

### Project Structure Understanding

```
PRPs-agentic-eng/
.claude/
  commands/           # 28+ Claude Code commands
  settings.local.json # Tool permissions
PRPs/
  templates/          # PRP templates with validation
  scripts/           # PRP runner and utilities
  ai_docs/           # Curated Claude Code documentation
   *.md               # Active and example PRPs
 claude_md_files/        # Framework-specific CLAUDE.md examples
 pyproject.toml         # Python package configuration
```

Remember: This framework is about **one-pass implementation success through comprehensive context**. Every PRP should contain enough context for an AI agent to implement working code without additional research.

## Core Principles

KISS (Keep It Simple, Stupid): Simplicity should be a key goal in design. Choose straightforward solutions over complex ones whenever possible. Simple solutions are easier to understand, maintain, and debug.

YAGNI (You Aren't Gonna Need It): Avoid building functionality on speculation. Implement features only when they are needed, not when you anticipate they might be useful in the future.

Dependency Inversion: High-level modules should not depend on low-level modules. Both should depend on abstractions. This principle enables flexibility and testability.

Open/Closed Principle: Software entities should be open for extension but closed for modification. Design your systems so that new functionality can be added with minimal changes to existing code.

### File Tree

DevQ.ai is a monorepo with several key python apps (agentical, machina, ptolemies) each one with a CLI in Go, and authenticaion managed with zero knowledge proof STARK, also a Go app. This is the planned file tree:

```
.
└── devqai/
    ├── .claude/*/*
    │   └── settings.local.json
    ├── .ghostty*/*
    │   └── terminal-config.zsh
    ├── .git*
    ├── .github/*/*
    ├── .logfire*/*
    │   └── logfire_credentials.json
    ├── .zed/*/*
    │   ├── settings.json
    │   └── terminal-config.zsh
    ├── agentical/ # agentic system using pydantic ai
    │   ├── agents/ # pydantic ai agents
    │   │   ├── core/ # 4 core agents assigned to every playbook
    │   │   │   ├── superagent.py
    │   │   │   ├── playbookagent.py
    │   │   │   ├── condifieragent.py
    │   │   │   └── inspectobserveagent.py
    │   │   └── specialized/ # 9 specialized agents
    │   │       ├── cloudagent.py
    │   │       ├── codeagent.py
    │   │       ├── datascienceagent.py
    │   │       ├── githubagent.py
    │   │       ├── legalagent.py
    │   │       ├── researchagent.py
    │   │       ├── shopifyagent.py
    │   │       ├── supportagent.py
    │   │       └── testeragent.py
    │   ├── playbooks/ # how agentical uses agents tools workflow
    │   │   ├── configs/ # config files of playbooks
    │   │   └── playbooks.py # script for playbookagent
    │   ├── readme.md
    │   ├── reasoning/ # with bayesian principles
    │   │   ├── bayes/ # conditional probability
    │   │   ├── darwin/ # genetic algorithms
    │   │   ├── nash/ # game theory
    │   │   ├── wildwood/ # random forest
    │   │   └── zipf/ # logistic population growth 
    │   ├── tools/ # tools available for agents
    │   │   ├── pydantic/ # pydantic ai built tools/
    │   │   │   ├── multimodal/ # pydantic agent multimodal tool
    │   │   │   └── tools/ # files for pydantic agent tools
    │   │   ├── tools.py # assigns tools to agents for playbook
    │   │   └── mcp_status.json # updated file from machina
    │   └── workflows/ # agentical's workflows 
    │       ├── patterns/ 
    │       │   ├── sequential/
    │       │   └── mapped/ 
    │       └── conditions/ 
    │           ├── feedback/ # boolean evalution
    │           └── quantitative/ # bitwise operators
    ├── api/ # devqai's plumbing
    │   └── readme.md
    ├── cli/ # devqai's cli with charm
    │   ├── go.mod # bubbletea, lipgloss, gum, log dependencies
    │   ├── go.sum
    │   ├── cmd/
    │   │   └── devqai/
    │   │       └── main.go
    │   ├── internal/
    │   │   ├── commands/
    │   │   │   ├── agentical/ # agentical subcommands
    │   │   │   ├── machina/ # machina subcommands  
    │   │   │   ├── ptolemies/ # ptolemies subcommands
    │   │   │   └── trustlis/ # trustlis subcommands
    │   │   ├── config/
    │   │   │   └── local.go # local config management
    │   │   ├── auth/
    │   │   │   └── trustlis.go # trustlis integration
    │   │   └── ui/               # charm ui components
    │   │       ├── styles.go # lip gloss styles
    │   │       ├── components/ # reusable bubble tea components
    │   │       └── tui/ # full tui applications
    │   ├── pkg/
    │   │   └── client/ # shared cli client utilities
    │   └── readme.md
    ├── docusaurus/ # documentation for devqai in docusaurus
    │   └── readme.md
    ├── frontend/ # nextjs for devqai/
    │   └── readme.md
    ├── logfire/ # syncd schemas, alerts to logfire.pydantic
    │   └── readme.md
    ├── machina/ # mcp registry using fast mcp and more/
    │   ├── registry/   
    │   ├── mcp/ # servers/
    │   └── readme.md
    ├── ptolemies/ # knowledge management with rag + graph
    │   ├── crawl4ai/       
    │   ├── readme.md
    │   └── surrealdb/ # devqai database
    │       ├── appdb/ # contains appdb for all devqai
    │       ├── dehallucinate/
    │       ├── graph/ # contains traversal 
    │       ├── logs/ # contains logfile logs
    │       └── vector/ # contains embeddings
    ├── tests/ # contains tests and results
    │   ├── agentical/ # python tests (pytest)
    │   ├── api/ # python tests (pytest)
    │   ├── appdb/ # python tests (pytest)
    │   ├── cli/ # go tests  
    │   │   ├── commands/
    │   │   ├── integration/
    │   │   └── testdata/
    │   ├── docusaurus/ # python tests (pytest)
    │   ├── trustlis/ # go tests
    │   │   ├── auth_test.go
    │   │   ├── crypto_test.go
    │   │   ├── zkstark_test.go
    │   │   └── integration/
    │   ├── frontend/ # js/ts tests
    │   ├── machina/ # python tests (pytest)
    │   ├── ptolemies/ # python tests (pytest)
    │   └── integration/ # cross-language integration (minimal)
    ├── trustlis/ # trustless authentication with zkstark
    │   ├── go.mod # independent, no charm deps
    │   ├── go.sum
    │   ├── cmd/
    │   │   └── trustlis/
    │   │       └── main.go
    │   ├── internal/
    │   │   ├── zkstark/
    │   │   ├── auth/
    │   │   ├── server/
    │   │   └── crypto/
    │   ├── pkg/
    │   │   ├── client/ # for cli/api integration
    │   │   └── types/ # shared auth types
    │   └── readme.md
    ├── venv/
    ├── .env*
    ├── .gitignore* # includes all * items
    ├── .rules*
    ├── changelog.md
    ├── claude.md*
    ├── readme.md
    ├── license
    ├── makefile # unified testing and build commands
    ├── setup.py # launch python apps
    └── tsconfig.json # launch front end
```
### 🧱 Code Structure & Modularity

- **Never create a file longer than 500 lines of code.** If a file approaches this limit, refactor by splitting it into modules or helper files.
- **Functions should be short and focused sub 50 lines of code** and have a single responsibility.
- **Classes should be short and focused sub 50 lines of code** and have a single responsibility.
- **Organize code into clearly separated modules**, grouped by feature or responsibility.

### 📎 Style & Conventions
- **Use Python** as the primary language.
- **Follow PEP8**, always use type hints, and format with `ruff`.
- **Use `pydanticv2` for data validation**.
- **ALWAYS use classes, data types, data models, for typesafety and verifiability**
- **ALWAYS use docstrings for every function** using the Google style

### Code Formatting
- **Backend**: 88 character line length, Black formatter
- **Frontend**: 100 character line length, single quotes, required semicolons
- **Python**: 3.12, Black formatter, Google-style docstrings
- **TypeScript**: Strict mode, ES2022 target

### 🚫 BRANDING RESTRICTION
**DO NOT ADD ANYTHING LIKE THIS BRANDING TO A COMMIT MESSAGE:**
```
🤖 Generated with [Claude Code](https://claude.ai/code)
Co-Authored-By: Claude <noreply@anthropic.com>
```

### Fundamentals

#### FastAPI
- FastAPI Foundation - Core web framework for ALL projects
- Required for all API endpoints in api/ directory
- Logfire instrumentation mandatory for every route with spans
- CORS middleware must be configured for frontend integration
- Global exception handling with error logging to Logfire
- Health endpoint required (/health) for monitoring
- Lifespan events for startup/shutdown with Logfire logging

#### Pydantic AI
- Pydantic AI - REQUIRED when agents are needed
- Required for all agent implementations in agentical/agents/
- System prompts mandatory with clear role definitions
- Tool integration with real MCP server connections
- Result type definitions using Pydantic models
- Context passing for user/workflow state management
- Real implementation only - no mock responses in agent tools

#### PyTest
- PyTest Build-to-Test - REQUIRED for every task progression
- Unit tests REQUIRED for every subtask before progression
- Test coverage minimum: 90% line coverage
- Test structure: Unit + Integration + API tests
- Test execution: pytest tests/ --cov=src/ --cov-report=html
- Task progression: Cannot advance without passing tests

#### Logfire
- Logfire Observability - REQUIRED for every event and operation
- Every event must be logged through Logfire
- Every function must have Logfire spans for observability
- Every API endpoint must include Logfire instrumentation
- Every error must be captured with context
- Performance metrics tracked for all operations

#### FastMCP
- FastMCP - REQUIRED for all MCP server implementations
- Required for all MCP implementations in machina/
- Server registration must be tracked in machina/registry/
- MCP status updates must sync to agentical/tools/mcp_status.json
- All MCP tools must be tested with real server connections
- Server discovery must be implemented for agent tool assignment

#### Charm.sh
- Charm.sh - REQUIRED for all CLI implementations
- Required for all CLI implementations in cli/
- Bubble Tea framework for all TUI components
- Lip Gloss for consistent styling across all CLI output
- Gum integration for shell script interactions
- Log library for all CLI logging with Logfire backend

#### zkSTARK
- zkSTARK - REQUIRED for all authentication
- Required for all authentication in trustlis/
- Zero-knowledge proofs for all user authentication
- STARK protocol implementation for trustless verification
- No private key storage - all auth must be cryptographically verifiable
- HTTP API for CLI and Python API integration

#### NextJS
- NextJS - REQUIRED for informational website only
- Required for informational website in frontend/ - NO DevQ.ai integration
- TypeScript mandatory for all components
- App Router for all routing (not Pages Router)
- Cyber & Pastel Design System implementation required
- Inter Nerd Font + Space Mono Nerd Font for typography
- Framer Motion for animations and interactions
- Static content only - no API calls to DevQ.ai backend
- Marketing/documentation focus - showcase features, not functionality

### Behavioural Guidelines

- Always use `uv` for package management.
- Always use `ruff` for linting.

- *** NEVER ASSUME OR GUESS ***
- When in doubt, ask for clarification or ask for help. more often than not youcan do websearch to find relevant examples of check ai_docs/ for examples that the user have added. 

- **Always confirm file paths & module names** exist before using them.

- **Comment non-obvious code** and ensure everything is understandable to a mid-level developer.
- When writing complex logic, **add an inline `# Reason:` comment** explaining the why, not just the what.

- **KEEP README.md UPDATED**
- Whenever you make changes to the codebase, update the README.md file to reflect the changes. Espeially if you add configuration changes or new features.

- **ALWAYS keep CLAUDE.md UPDATED**
- Add new dependencies to CLAUDE.md
- Add important types and patterns to CLAUDE.md
