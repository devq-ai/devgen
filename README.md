# DevGen - Development & Generation Toolkit

This directory contains resources and tools for AI-assisted development, primarily using Claude Code. It includes configuration files, code generation templates, and documentation to streamline various development tasks.

## Contents

*   `.claude/`: Configuration files for Claude Code, including allowed tools, commands, and project-specific settings.
*   `superclaude-main/`: Contains the SuperClaude CLI tool for AI-powered commit messages, changelogs, documentation, and code reviews. This tool uses Claude AI to automate and enhance development workflows.
*   `claude_md_files/`: A collection of `CLAUDE.md` files for various frameworks (React, Python, Node.js, Next.js). These files provide coding standards and guidelines for Claude Code.
*   `PRPs/`: Contains Product Requirement Prompts (PRPs), which are structured prompts that provide AI coding agents with the information they need to deliver a vertical slice of working software. PRPs include:
    *   `ai_docs/`: AI documentation
    *   `scripts/`: Scripts, such as `prp_runner.py`, used to execute PRPs.
    *   `files/`: Files used by the PRP
    *   `templates/`: Templates for creating PRPs.
*   `CLAUDE.md`: A comprehensive guide for Claude Code, outlining core principles, project structure, coding standards, development commands, and dependency management.
*   `DEVQ_SSH_SUCCESS.md`: Documentation outlining the successful implementation of the DevQ.ai SSH server, including web API and SSH terminal access details.
*   `.gitignore`: Specifies intentionally untracked files that Git should ignore.

## Key Tools & Technologies

*   **Claude Code:** An agentic coding tool that lives in your terminal.
*   **SuperClaude:** A CLI tool for automating commit messages, changelogs, documentation, and code reviews using Claude AI.
*   **FastMCP:** Used for creating MCP servers.
*   **Charm.sh:** Used to build interactive TUI (Text User Interface) CLI applications.
*   **Pydantic:** Used for data validation.

## Usage

This directory provides the building blocks for creating and executing Product Requirement Prompts (PRPs) with Claude Code.

1.  **Configure Claude Code:** Set up Claude Code with the configurations in the `.claude/` directory.
2.  **Explore PRPs:** Review the example PRPs in the `PRPs/` directory to understand how to structure prompts for different tasks.
3.  **Execute PRPs:** Use the `prp_runner.py` script or appropriate Claude commands to execute PRPs and generate code.
4.  **Utilize SuperClaude:** Leverage the SuperClaude CLI to automate common Git tasks, such as generating commit messages and creating changelogs.

## Contributing

Please adhere to the coding standards and guidelines outlined in `CLAUDE.md`. When modifying or creating PRPs, ensure that they are well-documented and thoroughly tested.
