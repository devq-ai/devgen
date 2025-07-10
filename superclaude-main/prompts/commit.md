Write a concise commit message in conventional commits format for these changes:

Files changed: {{file_changes}}

Git diff stat:
{{diff_stat}}

{{additional_context}}

Requirements:
- Use conventional commits format (feat:, fix:, refactor:, etc.)
- Keep the first line under 50 characters
- Add bullet points for details if needed
- Be specific about what was changed
- Output ONLY the raw commit message text
- Do NOT include any prefatory text like 'Here's the commit message:' or 'Based on the changes:'
- Do NOT use markdown formatting, backticks, or code blocks
- Do NOT include attribution lines like 'Generated with Claude Code' or 'Co-Authored-By: Claude'
- Start directly with the commit message (e.g. 'feat: add user authentication') 