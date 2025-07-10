# PRP-1: FastMCP Machina Registry Development

## Goal
Create a production-ready MCP server registry using FastMCP that allows service discovery, health monitoring, and dynamic tool assignment for agentical workflows.

## Why
Agentical agents need reliable access to MCP tools through a centralized registry. The current manual tool assignment is brittle and doesn't scale. A registry enables automatic tool discovery, health checking, and seamless integration with pydantic-ai agents.

## What
A FastMCP-based registry service that:
- Registers and tracks MCP server instances
- Provides health monitoring and automatic failover
- Exposes a unified API for tool discovery
- Syncs tool availability to agentical's mcp_status.json

## Implementation Blueprint

### Phase 1: Core Registry Service
```python
# machina/registry/main.py
from fastmcp import FastMCP
import logfire
from typing import Dict, List
import asyncio

class MCPRegistry:
    def __init__(self):
        self.mcp = FastMCP("mcp-registry")
        self.servers: Dict[str, dict] = {}
        self._setup_tools()
    
    @logfire.instrument()
    def _setup_tools(self):
        @self.mcp.tool()
        async def register_server(name: str, endpoint: str, tools: List[str]) -> Dict[str, str]:
            """Register a new MCP server"""
            # Implementation here
            pass
```

### Validation Commands
```bash
# Level 1: Syntax
ruff check machina/ && mypy machina/

# Level 2: Tests  
pytest tests/test_registry.py -v

# Level 3: Integration
curl -X POST http://localhost:8000/register -d '{"name":"test","endpoint":"localhost:8001"}'
```

## Confidence Score: 8/10
High confidence due to clear requirements and existing FastMCP patterns.